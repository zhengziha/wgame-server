// Package handlers 提供一组 demo 处理器，用于验证整条链路（codec / dispatch / firewall / broadcaster）。
//
// 每个处理器在 init() 中通过 handler.RegisterTyped 自注册，
// 因此只要在某处（例如 main.go）以匿名导入方式 import 此包即可生效。
package handlers

import (
	"wgame-server/comm/log"
	"wgame-server/server/codec"
	"wgame-server/server/context"
	demomsg "wgame-server/server/demo/msg"
	"wgame-server/server/network/handler"
)

func init() {
	handler.RegisterTyped(
		demomsg.CmdEchoReq,
		"EchoReq",
		func() *demomsg.EchoReq { return &demomsg.EchoReq{} },
		EchoHandler,
	)
}

// EchoHandler 直接拿到已反序列化好的 *EchoReq，
// 不再需要手写 reader.ReadString()。
func EchoHandler(ctx context.MyCmdContext, frame *codec.Frame, req *demomsg.EchoReq) error {
	log.Info("[demo] echo sid=%d uid=%d text=%q",
		ctx.GetSessionId(), ctx.GetUserId(), req.Text)

	ctx.Write(&demomsg.EchoMsg{Text: req.Text})
	return nil
}
