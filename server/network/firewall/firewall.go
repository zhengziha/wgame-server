package firewall

import (
	"sync"
	"time"

	"wgame-server/comm/log"
	"wgame-server/server/codec"
	"wgame-server/server/context"
)

// Firewall 单连接维度的安全检测，参考 hero_story.go_server/biz_server/network/firewall。
//
// 主要能力：
//   - 单连接最小帧间隔检测（过快视为 WPE/脚本）
//   - 单连接时间窗口内的帧数限流
//   - 非法魔数 / 超长帧由 codec 层直接拒收
//
// 所有判断都基于"该连接最近 N 个包的到达序列"，不涉及业务。
// 因此可以放在 IO goroutine 中调用，无需占用主线程。
type Firewall struct {
	// 最小包间隔（小于该值视为异常）
	MinInterval time.Duration
	// 滑动窗口长度
	WindowSize time.Duration
	// 窗口内允许的最大包数
	MaxPacketsInWindow int

	mu             sync.Mutex
	lastTs         time.Time
	recentPacketTs []time.Time
	warnedAt       time.Time
	kickCount      int
}

// NewFirewall 构造一个默认参数的防火墙
func NewFirewall() *Firewall {
	return &Firewall{
		MinInterval:        20 * time.Millisecond, // 单连接每秒最多约 50 包
		WindowSize:         5 * time.Second,
		MaxPacketsInWindow: 300,
		recentPacketTs:     make([]time.Time, 0, 64),
	}
}

// Check 在每个入站帧到达时调用。
// 返回 true 表示通过；false 表示应被防火墙拦截（上层应关闭连接）。
// 同时会记录告警日志，避免日志风暴。
func (f *Firewall) Check(ctx context.MyCmdContext, frame *codec.Frame) bool {
	if ctx == nil {
		return true
	}
	now := time.Now()

	f.mu.Lock()
	defer f.mu.Unlock()

	// 1) 最小间隔检测
	if !f.lastTs.IsZero() {
		if now.Sub(f.lastTs) < f.MinInterval {
			f.kickCount++
			if now.Sub(f.warnedAt) > 5*time.Second {
				log.Error("[firewall] packet too fast, sid=%d ip=%s cmd=%d kick=%d",
					ctx.GetSessionId(), ctx.GetClientIpAddr(), frame.Cmd, f.kickCount)
				f.warnedAt = now
			}
			return false
		}
	}
	f.lastTs = now

	// 2) 滑动窗口限流
	//    使用切片头部 reslice 替代"每次都新建底层数组"，
	//    避免高频包场景下 O(N) 拷贝。仅当尾部累计空闲 >= windowSize 时
	//    才整体 compact 一次，平摊代价 O(1)。
	f.recentPacketTs = append(f.recentPacketTs, now)
	cutoff := now.Add(-f.WindowSize)
	idx := 0
	for ; idx < len(f.recentPacketTs); idx++ {
		if f.recentPacketTs[idx].After(cutoff) {
			break
		}
	}
	if idx > 0 {
		f.recentPacketTs = f.recentPacketTs[idx:]
	}
	// 当头部已退化的累积量超过窗口容量的一半时，做一次 compact 回收内存
	if cap(f.recentPacketTs)-len(f.recentPacketTs) > f.MaxPacketsInWindow/2 {
		f.recentPacketTs = append([]time.Time(nil), f.recentPacketTs...)
	}
	if len(f.recentPacketTs) > f.MaxPacketsInWindow {
		f.kickCount++
		if now.Sub(f.warnedAt) > 5*time.Second {
			log.Error("[firewall] packet rate too high, sid=%d ip=%s count=%d/%ds",
				ctx.GetSessionId(), ctx.GetClientIpAddr(),
				len(f.recentPacketTs), int(f.WindowSize.Seconds()))
			f.warnedAt = now
		}
		return false
	}
	return true
}
