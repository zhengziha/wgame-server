package session

import (
	"sync"
	"sync/atomic"
)

// GameSession 表示一条逻辑会话。
// 参考 Java wd-server-fl core/GameSession.java。
//
// 字段刻意保持精简：业务层可以自由扩展（如 Chara、Team、Match 字段）。
type GameSession struct {
	// 逻辑字段（atomic / mutex 保护）
	id        int64 // 玩家 id；未登录为 0
	sessionId int32 // 连接 id；与 CmdContext.SessionId 对应

	// 连接元信息
	ClientIP  string
	ChannelID string // 等价 Java Netty channel long text id

	// 账号信息（登录后填充）
	AccountID   int64
	AccountName string

	// chara 业务占位字段，业务侧可自行替换为具体类型
	Chara interface{}

	// 离线相关
	isConnection atomic.Bool
	lastInactive int64 // unix milli

	// 会话级串行锁（与 Java GameSession.lock 等价）。
	// 由 handler 派发层在调用业务 handler 前后 CAS 持有/释放。
	lock atomic.Bool

	// 反向引用到 CmdContext（仅持有接口，避免循环依赖）
	cmdCtx interface{} // 实际类型为 context.MyCmdContext

	mu sync.RWMutex
}

// NewGameSession 构造一条新会话
func NewGameSession(sessionId int32, clientIP, channelID string) *GameSession {
	s := &GameSession{
		sessionId: sessionId,
		ClientIP:  clientIP,
		ChannelID: channelID,
	}
	s.isConnection.Store(true)
	return s
}

// ID 返回玩家 id
func (s *GameSession) ID() int64 { return s.id }

// SetID 设置玩家 id（登录成功后调用）
func (s *GameSession) SetID(id int64) {
	s.mu.Lock()
	s.id = id
	s.mu.Unlock()
}

// SessionID 返回连接 id
func (s *GameSession) SessionID() int32 { return s.sessionId }

// ActorKey 返回用于 worker 分片的 key。
// 与 Java 一致：未登录用 accountId，登录后用 chara.id（这里用 ID()）。
func (s *GameSession) ActorKey() int {
	if s.id != 0 {
		return int(s.id)
	}
	return int(s.AccountID)
}

// IsOnline 返回会话是否在线
func (s *GameSession) IsOnline() bool { return s.isConnection.Load() }

// BindCmdContext 把 CmdContext 反向绑定到 session（用 interface 规避循环依赖）
func (s *GameSession) BindCmdCtx(ctx interface{}) {
	s.mu.Lock()
	s.cmdCtx = ctx
	s.mu.Unlock()
}

// CmdCtx 取出绑定的 CmdContext（业务侧通常会做类型断言）
func (s *GameSession) CmdCtx() interface{} {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.cmdCtx
}

// TryLock 尝试获取会话级串行锁（CAS）。
// 返回 true 表示获取成功，调用方负责调用 Unlock。
func (s *GameSession) TryLock() bool {
	return s.lock.CompareAndSwap(false, true)
}

// Unlock 释放会话级串行锁
func (s *GameSession) Unlock() {
	s.lock.Store(false)
}

// MarkOffline 标记会话离线
func (s *GameSession) MarkOffline(nowMilli int64) {
	s.isConnection.Store(false)
	atomic.StoreInt64(&s.lastInactive, nowMilli)
}

// LastInactiveMilli 返回最后离线时间
func (s *GameSession) LastInactiveMilli() int64 {
	return atomic.LoadInt64(&s.lastInactive)
}
