package broadcaster

import (
	"sync"

	"wgame-server/server/context"
)

// Broadcaster 维护 sessionId -> MyCmdContext 的在线索引，
// 提供 O(1) 单点发送与批量广播能力。
//
// 参考 hero_story.go_server/biz_server/network/broadcaster/broadcaster.go。
//
// 所有方法都通过内部 sync.RWMutex 串行化，可在多 goroutine 下安全调用。
// 业务 handler 一般运行在主线程（main_thread.Process）内，
// 调用 Broadcast 时只是把消息体投递给每条连接自己的发送队列。
//
// 索引：
//   - bySession: sessionId  → ctx（连接维度，始终存在）
//   - byUser   : userId     → ctx（登录后由 BindUserId 触发建立）
//
// 设计说明：
//   - 单用户可能有多条连接（双端登录、断线重连瞬时叠加），byUser 只保留最新一条
//   - 业务若需要"向用户全部连接推送"，可自行维护 userId → []sessionId
var (
	mu sync.RWMutex

	bySession = make(map[int32]context.MyCmdContext, 4096)
	byUser    = make(map[int64]context.MyCmdContext, 1024)
)

// AddCmdCtx 注册一条新连接。若 sessionId 已存在则覆盖。
// 注意：AddCmdCtx 不会建立 userId 索引；若该 ctx 已绑定 userId，
// 调用方应在 BindUserId 后调用 BindUserToCtx 同步建立索引。
func AddCmdCtx(ctx context.MyCmdContext) {
	mu.Lock()
	bySession[ctx.GetSessionId()] = ctx
	// 若已经绑定 userId，顺手建立 byUser 索引（覆盖旧值）
	if uid := ctx.GetUserId(); uid != 0 {
		byUser[uid] = ctx
	}
	mu.Unlock()
}

// BindUserToCtx 在 ctx.BindUserId 调用之后同步建立 userId 索引。
// 应在业务 handler 登录成功路径中调用：
//
//	ctx.BindUserId(uid)
//	broadcaster.BindUserToCtx(uid, ctx)
//
// 重复调用会覆盖旧 ctx（旧连接的索引会被指向新连接）。
func BindUserToCtx(userId int64, ctx context.MyCmdContext) {
	if userId == 0 || ctx == nil {
		return
	}
	mu.Lock()
	byUser[userId] = ctx
	mu.Unlock()
}

// RemoveCmdCtxBySessionId 移除并返回被移除的 ctx，不存在返回 nil。
// 同时会清理该 ctx 对应的 userId 索引（若指向同一 ctx）。
func RemoveCmdCtxBySessionId(sessionId int32) context.MyCmdContext {
	mu.Lock()
	defer mu.Unlock()
	ctx, ok := bySession[sessionId]
	if !ok {
		return nil
	}
	delete(bySession, sessionId)
	// 仅在 byUser 仍指向同一条 ctx 时才删除，
	// 避免"同用户新连接已覆盖旧索引"时误删新连接的索引。
	if uid := ctx.GetUserId(); uid != 0 {
		if cur, ok := byUser[uid]; ok && cur.GetSessionId() == sessionId {
			delete(byUser, uid)
		}
	}
	return ctx
}

// GetCmdCtx 按 sessionId 查找在线 ctx
func GetCmdCtx(sessionId int32) (context.MyCmdContext, bool) {
	mu.RLock()
	defer mu.RUnlock()
	ctx, ok := bySession[sessionId]
	return ctx, ok
}

// GetCmdCtxByUserId 按 userId 查找最新绑定的在线 ctx。
// 未登录或未注册返回 false。
func GetCmdCtxByUserId(userId int64) (context.MyCmdContext, bool) {
	if userId == 0 {
		return nil, false
	}
	mu.RLock()
	defer mu.RUnlock()
	ctx, ok := byUser[userId]
	return ctx, ok
}

// Count 返回当前在线连接数（按 sessionId）
func Count() int {
	mu.RLock()
	defer mu.RUnlock()
	return len(bySession)
}

// CountUsers 返回当前已绑定 userId 的在线连接数。
func CountUsers() int {
	mu.RLock()
	defer mu.RUnlock()
	return len(byUser)
}

// Broadcast 向所有在线连接发送同一条消息对象。
// 调用方应提供一个构造好的 msgObj（实现 msg.OutMessage 或 ctx.Write 可识别的形态）。
// 调用此函数不会阻塞任一连接的写队列。
func Broadcast(msgObj interface{}) {
	mu.RLock()
	defer mu.RUnlock()
	for _, ctx := range bySession {
		ctx.Write(msgObj)
	}
}

// BroadcastExcept 向除 excludeSessionId 外的所有在线连接发送消息
func BroadcastExcept(msgObj interface{}, excludeSessionId int32) {
	mu.RLock()
	defer mu.RUnlock()
	for sid, ctx := range bySession {
		if sid == excludeSessionId {
			continue
		}
		ctx.Write(msgObj)
	}
}

// SendToUser 向指定 userId 的最新连接发送消息。
// 未绑定或已离线返回 false。
func SendToUser(msgObj interface{}, userId int64) bool {
	ctx, ok := GetCmdCtxByUserId(userId)
	if !ok {
		return false
	}
	ctx.Write(msgObj)
	return true
}

// AllSnapshot 返回当前所有 ctx 的快照（用于自定义遍历逻辑）。
// 调用方可以在持有切片期间放心遍历，不会受 broadcaster 内部并发修改影响。
func AllSnapshot() []context.MyCmdContext {
	mu.RLock()
	defer mu.RUnlock()
	out := make([]context.MyCmdContext, 0, len(bySession))
	for _, ctx := range bySession {
		out = append(out, ctx)
	}
	return out
}
