package context

// MyCmdContext 是与传输无关的“连接上下文”抽象。
// 参考 hero_story.go_server/biz_server/base/my_cmd_context.go。
//
// 业务 handler 只依赖该接口，无需关心底层是 TCP / WebSocket / KCP。
// 当前实现位于 server/socket.SocketCmdContext（基于 raw TCP）。
type MyCmdContext interface {
	// BindUserId 在登录成功后绑定业务 user id
	// 类型为 int32，对应数据库 characters.id（与 wd-server-fl 对齐）
	BindUserId(val int32)

	// GetUserId 返回当前绑定的 user id；未登录返回 0
	GetUserId() int32

	// BindGid 在登录成功后绑定全局唯一 id
	// gid 来自 characters.gid（32 字符无横杠 UUID）
	BindGid(gid string)

	// GetGid 返回当前绑定的 gid；未登录返回空串
	// 客户端协议应使用 gid 作为玩家标识
	GetGid() string

	// GetSessionId 返回连接的唯一标识（int32 自增）
	GetSessionId() int32

	// GetClientIpAddr 返回对端 IP 地址（不含端口）
	GetClientIpAddr() string

	// Write 投递一条出站消息到发送队列（非阻塞）
	Write(msgObj interface{})

	// SendError 向客户端发送错误码消息（具体协议由业务层决定）
	SendError(errorCode int, errorInfo string)

	// Disconnect 关闭底层连接
	Disconnect()
}
