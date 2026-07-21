package socket

import (
	"sync"
	"sync/atomic"

	"wgame-server/comm/log"
	"wgame-server/server/codec"
	"wgame-server/server/context"
	"wgame-server/server/msg"
	"wgame-server/server/network/firewall"
	"wgame-server/server/session"
)

// SocketCmdContext 基于 raw TCP 的 MyCmdContext 实现。
// 每条连接对应一个 SocketCmdContext + GameSession。
//
// 写路径：handler 调用 Write -> 编码帧 -> 写 chan sendQ -> sendLoop 持续 flush。
// 读路径：readLoop 读帧 -> firewall.Check -> 主线程 dispatch handler。
type SocketCmdContext struct {
	session *session.GameSession

	conn *tcpConn // 由 server.go 内部持有，避免直接暴露 net.Conn

	// 出站发送队列：handler 投递、sendLoop 消费
	sendQ chan []byte

	// 关闭状态
	closed  atomic.Bool
	closeMu sync.Mutex

	// 防火墙
	firewall *firewall.Firewall

	// session id 自增（在 server 层注入）
	sessionId int32
}

// NewSocketCmdContext 构造一个新的连接上下文（尚未 start）
func NewSocketCmdContext(c *tcpConn, sessionId int32, fw *firewall.Firewall) *SocketCmdContext {
	sess := session.NewGameSession(sessionId, c.remoteIP, c.id)
	ctx := &SocketCmdContext{
		session:   sess,
		conn:      c,
		sendQ:     make(chan []byte, 256),
		firewall:  fw,
		sessionId: sessionId,
	}
	sess.BindCmdCtx(ctx)
	return ctx
}

// 实现 context.MyCmdContext

func (c *SocketCmdContext) BindUserId(val int32) {
	c.session.SetID(val)
}

func (c *SocketCmdContext) GetUserId() int32 {
	return c.session.ID()
}

func (c *SocketCmdContext) BindGid(gid string) {
	c.session.BindGid(gid)
}

func (c *SocketCmdContext) GetGid() string {
	return c.session.Gid()
}

func (c *SocketCmdContext) GetSessionId() int32 {
	return c.sessionId
}

func (c *SocketCmdContext) GetClientIpAddr() string {
	return c.session.ClientIP
}

// Write 把消息对象编码成帧并投递到发送队列。
// msgObj 必须实现 msg.OutMessage，否则直接丢弃并记录。
//
// 错误处理策略：
//   - msgObj 不是 OutMessage：记 Error（业务代码错误，必须修）
//   - WriteFrame 编码失败（通常是反射 tag 错误）：记 Error（业务代码错误，必须修）
//   - 连接已关闭：静默返回（正常生命周期事件，不是错误）
func (c *SocketCmdContext) Write(msgObj interface{}) {
	m, ok := msgObj.(msg.OutMessage)
	if !ok {
		log.Error("[socket] Write rejected: msgObj %T does not implement msg.OutMessage; cmd sid=%d",
			msgObj, c.sessionId)
		return
	}
	data, err := msg.WriteFrame(m, -1, 0)
	if err != nil {
		// 编码失败几乎全是反射 tag 错误或 nil 指针，属于程序 bug
		log.Error("[socket] Write encode failed: cmd=0x%04X sid=%d err=%v",
			m.Cmd(), c.sessionId, err)
		return
	}
	c.enqueueFrame(data)
}

// enqueueFrame 把已编码的字节切片投递到发送队列。
// 满则关闭连接，避免慢消费 client 拖垮服务端。
func (c *SocketCmdContext) enqueueFrame(data []byte) {
	if c.closed.Load() {
		return
	}
	// closed-check 之后到 send 之间，sendQ 可能已被关闭，需要 recover
	defer func() { _ = recover() }()
	select {
	case c.sendQ <- data:
	default:
		// 队列满 -> 视为慢 client，主动断开
		c.Disconnect()
	}
}

// SendError 默认实现：业务上可以构造一个错误消息再 Write。
// 基础框架不约束错误协议格式，这里只做占位。
func (c *SocketCmdContext) SendError(errorCode int, errorInfo string) {
	// 由业务模块覆盖：例如实现一个 GameErrorMsg{Code, Info} 并 Write。
}

// Disconnect 关闭底层连接并关闭发送队列（幂等）。
// 调用后 sendLoop 会通过 `sendQ 关闭` 退出。
func (c *SocketCmdContext) Disconnect() {
	c.closeMu.Lock()
	defer c.closeMu.Unlock()
	if c.closed.Load() {
		return
	}
	c.closed.Store(true)
	c.session.MarkOffline(timeNowMilli())
	c.conn.close()
	close(c.sendQ)
}

// Session 暴露内部 GameSession，业务侧可自行扩展
func (c *SocketCmdContext) Session() *session.GameSession { return c.session }

// Firewall 暴露给 server 的 readLoop 调用
func (c *SocketCmdContext) Firewall() *firewall.Firewall { return c.firewall }

// SendQ 暴露发送队列，由 server.go 的 sendLoop 消费
func (c *SocketCmdContext) SendQ() <-chan []byte { return c.sendQ }

// Compile-time check
var _ context.MyCmdContext = (*SocketCmdContext)(nil)

// timeNowMilli 提供可替换的时钟（避免在低层依赖 time 包的副作用）
func timeNowMilli() int64 {
	return codec.NowMilli()
}
