// Package handler 提供两种 cmd 处理器注册方式：
//
//  1. 原始方式 Register：handler 自己从 GameReader 手动读取字段
//  2. 类型化方式 RegisterTyped：框架根据 Req 工厂自动反序列化，
//     业务函数直接拿到具体 Req 指针，无需关心协议读取
//
// 类型化方式适合"POJO 消息体"——struct 字段加 `codec:"..."` tag 即可，
// 不必再为每个消息写 ReadBody 模板代码。
package handler

import (
	"wgame-server/server/codec"
	"wgame-server/server/context"
	"wgame-server/server/msg"
)

// TypedCmdHandlerFunc 类型化处理器函数签名。
//
//   - ctx   : 当前连接上下文
//   - frame : 已解出的完整帧
//   - req   : 已根据 Req 工厂构造并用反射填充好的具体 Req 对象
//
// 返回 error 用于上报日志；不影响写回消息。
type TypedCmdHandlerFunc[T msg.InMessage] func(ctx context.MyCmdContext, frame *codec.Frame, req T) error

// RegisterTyped 注册一个"类型化"处理器：
//
//   - cmd          : 处理的 cmd
//   - name         : 处理器名称（日志用）
//   - newReq       : Req 构造工厂，返回类型即 T（由编译器保证类型一致）
//   - handlerFunc  : 业务处理函数，直接拿到具体 Req 类型
//
// 框架在派发时会：
//
//  1. 调用 newReq() 构造空 Req
//  2. 通过反射按字段顺序从 frame.Body 反序列化
//  3. 把具体 Req 传给 handlerFunc
//
// 类型安全：newReq 的签名是 func() T，T 由调用方通过具体 Req 类型推导，
// 因此 newReq 返回的对象与 handlerFunc 期望的类型在编译期就被绑定，
// 不存在运行时类型断言失败的风险。
func RegisterTyped[T msg.InMessage](cmd uint16, name string, newReq func() T, handlerFunc TypedCmdHandlerFunc[T]) {
	if newReq == nil {
		panic("handler: RegisterTyped newReq is nil for cmd " + name)
	}
	if handlerFunc == nil {
		panic("handler: RegisterTyped handler is nil for cmd " + name)
	}

	// 同时注册到 msg.ReqRegistry，便于其他模块（调试/客户端模拟）通过 cmd 构造请求
	msg.RegisterReq(cmd, func() msg.InMessage { return newReq() })

	wrapped := func(ctx context.MyCmdContext, frame *codec.Frame, reader *codec.GameReader) error {
		req := newReq()
		if err := msg.ReadFramePayload(req, reader); err != nil {
			return err
		}
		return handlerFunc(ctx, frame, req)
	}
	Register(cmd, name, wrapped)
}

// GetReq 根据 cmd 构造一个 Req 对象（未注册返回 nil,false）。
// 是 msg.NewReq 的转发别名，便于 handler 包使用者避免额外 import msg。
func GetReq(cmd uint16) (msg.InMessage, bool) {
	return msg.NewReq(cmd)
}
