// Package handlers 提供一组 demo 处理器，用于验证整条链路（codec / dispatch / firewall / broadcaster）。
//
// 每个处理器在 init() 中通过 handler.Register 自注册，
// 因此只要在某处（例如 main.go）以匿名导入方式 import 此包即可生效。
package handlers

import (
	"wgame-server/comm/log"
	"wgame-server/server/codec"
	"wgame-server/server/context"
	demomsg "wgame-server/server/demo/msg"
	"wgame-server/server/network/handler"
)

// CmdEchoReq 客户端发来的回显请求，约定 cmd=0x0101
const CmdEchoReq uint16 = 0x0101

func init() {
	handler.Register(CmdEchoReq, "EchoReq", EchoHandler)
}

// EchoHandler 读一个字符串，原样回发一条 EchoMsg。
// 可用于联调测试：客户端发送任意字符串，服务端应回同样字符串。
func EchoHandler(ctx context.MyCmdContext, frame *codec.Frame, reader *codec.GameReader) error {
	text, _ := reader.ReadString()
	log.Info("[demo] echo sid=%d uid=%d text=%q",
		ctx.GetSessionId(), ctx.GetUserId(), text)

	ctx.Write(&demomsg.EchoMsg{Text: text})
	return nil
}
