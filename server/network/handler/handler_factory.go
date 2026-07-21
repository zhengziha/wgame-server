package handler

import (
	"sync"

	"wgame-server/server/codec"
	"wgame-server/server/context"
)

// CmdHandlerFunc 业务处理器函数签名。
//
//   - ctx   : 当前连接上下文（登录后可读 userId）
//   - frame : 已解出的完整帧（含 cmd 与 payload）
//   - reader: 包装好 frame.Body 的 GameReader，handler 直接按字段顺序读取
//
// 返回 error 用于上报防火墙/日志；不影响写回消息。
type CmdHandlerFunc func(ctx context.MyCmdContext, frame *codec.Frame, reader *codec.GameReader) error

// CmdHandler 包装处理器及其元信息
type CmdHandler struct {
	Cmd     uint16
	Name    string
	Handler CmdHandlerFunc
}

var (
	mu       sync.RWMutex
	registry = make(map[uint16]*CmdHandler, 256)
)

// Register 注册一个 cmd 处理器。
// 推荐在各 handler 包的 init() 中调用，实现"插件式"自注册。
// 重复注册（同 cmd）会覆盖旧处理器，并保留最后注册的 Name。
func Register(cmd uint16, name string, h CmdHandlerFunc) {
	if h == nil {
		panic("handler: Register handler is nil for cmd " + name)
	}
	mu.Lock()
	defer mu.Unlock()
	registry[cmd] = &CmdHandler{
		Cmd:     cmd,
		Name:    name,
		Handler: h,
	}
}

// Get 查找 cmd 对应的处理器
func Get(cmd uint16) (*CmdHandler, bool) {
	mu.RLock()
	defer mu.RUnlock()
	h, ok := registry[cmd]
	return h, ok
}

// Count 返回已注册的处理器数量
func Count() int {
	mu.RLock()
	defer mu.RUnlock()
	return len(registry)
}

// AllSnapshot 返回所有已注册处理器的快照（主要用于调试/启动日志）
func AllSnapshot() []*CmdHandler {
	mu.RLock()
	defer mu.RUnlock()
	out := make([]*CmdHandler, 0, len(registry))
	for _, h := range registry {
		out = append(out, h)
	}
	return out
}

// Dispatch 调用 cmd 对应的处理器；未注册时返回 false（由上层决定是否记录警告/踢线）。
func Dispatch(ctx context.MyCmdContext, frame *codec.Frame) (bool, error) {
	h, ok := Get(frame.Cmd)
	if !ok {
		return false, nil
	}
	reader := codec.NewGameReader(frame.Body)
	return true, h.Handler(ctx, frame, reader)
}
