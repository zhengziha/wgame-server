package session

import (
	"sync"
	"sync/atomic"
	"wgame-server/server/game"
)

// GameSession 表示一条逻辑会话。
// 参考 Java wd-server-fl core/GameSession.java。
//
// 字段刻意保持精简：业务层可以自由扩展（如 Chara、Team、Match 字段）。
type GameSession struct {
	// 逻辑字段（atomic / mutex 保护）
	// id：玩家 id（对应数据库 characters.id，int32）；未登录为 0。
	// 读写都走 atomic，消除 data race。
	id atomic.Int32

	// gid：全局唯一 id（对应 characters.gid，32 字符无横杠 UUID）。
	// 客户端协议只认 gid，不认 id。登录成功后填充，由 mu 保护。
	gid string

	// sessionId 在构造后不变，无需同步。
	sessionId int32 // 连接 id；与 CmdContext.SessionId 对应

	// 连接元信息
	ClientIP  string
	ChannelID string // 等价 Java Netty channel long text id

	// 账号信息（登录后填充；读写由 mu 保护）
	// AccountID 对应 accounts.id，与 wd-server-fl 一致使用 int32。
	AccountID   int32
	AccountName string

	// chara 玩家运行时数据（登录后加载）
	// 参考 Java wd-server-fl core/domain/Chara.java
	Chara *game.Chara

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

// ID 返回玩家 id（原子读，并发安全）
func (s *GameSession) ID() int32 { return s.id.Load() }

// SetID 设置玩家 id（原子写，并发安全；登录成功后调用）
func (s *GameSession) SetID(id int32) {
	s.id.Store(id)
}

// Gid 返回全局唯一 id（登录后填充，并发安全）。
// 客户端协议使用 gid 作为玩家标识，不使用数据库主键 id。
func (s *GameSession) Gid() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.gid
}

// BindGid 绑定全局唯一 id（登录成功后调用，并发安全）。
// gid 来自 characters.gid，在创建角色时由 uuidutil.NewGid() 生成。
func (s *GameSession) BindGid(gid string) {
	s.mu.Lock()
	s.gid = gid
	s.mu.Unlock()
}

// SessionID 返回连接 id
func (s *GameSession) SessionID() int32 { return s.sessionId }

// ActorKey 返回用于 worker 分片的 key。
// 与 Java 一致：未登录用 accountId，登录后用 chara.id（这里用 ID()）。
func (s *GameSession) ActorKey() int {
	id := s.ID()
	if id != 0 {
		return int(id)
	}
	s.mu.RLock()
	aid := s.AccountID
	s.mu.RUnlock()
	return int(aid)
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

// GetChara 返回玩家运行时数据
// 实现 game.Session 接口
func (s *GameSession) GetChara() *game.Chara {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.Chara
}

// SetChara 设置玩家运行时数据（登录后调用）
func (s *GameSession) SetChara(chara *game.Chara) {
	s.mu.Lock()
	s.Chara = chara
	s.mu.Unlock()
}
