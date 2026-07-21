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
var (
	mu  sync.RWMutex
	all = make(map[int32]context.MyCmdContext, 4096)
)

// AddCmdCtx 注册一条新连接。若 sessionId 已存在则覆盖。
func AddCmdCtx(ctx context.MyCmdContext) {
	mu.Lock()
	all[ctx.GetSessionId()] = ctx
	mu.Unlock()
}

// RemoveCmdCtxBySessionId 移除并返回被移除的 ctx，不存在返回 nil。
func RemoveCmdCtxBySessionId(sessionId int32) context.MyCmdContext {
	mu.Lock()
	defer mu.Unlock()
	ctx, ok := all[sessionId]
	if !ok {
		return nil
	}
	delete(all, sessionId)
	return ctx
}

// GetCmdCtx 按 sessionId 查找在线 ctx
func GetCmdCtx(sessionId int32) (context.MyCmdContext, bool) {
	mu.RLock()
	defer mu.RUnlock()
	ctx, ok := all[sessionId]
	return ctx, ok
}

// Count 返回当前在线连接数
func Count() int {
	mu.RLock()
	defer mu.RUnlock()
	return len(all)
}

// Broadcast 向所有在线连接发送同一条消息对象。
// 调用方应提供一个构造好的 msgObj（实现 msg.OutMessage 或 ctx.Write 可识别的形态）。
// 调用此函数不会阻塞任一连接的写队列。
func Broadcast(msgObj interface{}) {
	mu.RLock()
	defer mu.RUnlock()
	for _, ctx := range all {
		ctx.Write(msgObj)
	}
}

// BroadcastExcept 向除 excludeSessionId 外的所有在线连接发送消息
func BroadcastExcept(msgObj interface{}, excludeSessionId int32) {
	mu.RLock()
	defer mu.RUnlock()
	for sid, ctx := range all {
		if sid == excludeSessionId {
			continue
		}
		ctx.Write(msgObj)
	}
}

// AllSnapshot 返回当前所有 ctx 的快照（用于自定义遍历逻辑）。
// 调用方可以在持有切片期间放心遍历，不会受 broadcaster 内部并发修改影响。
func AllSnapshot() []context.MyCmdContext {
	mu.RLock()
	defer mu.RUnlock()
	out := make([]context.MyCmdContext, 0, len(all))
	for _, ctx := range all {
		out = append(out, ctx)
	}
	return out
}
