package system

import (
	"wgame-server/server/codec"
	"wgame-server/server/msg"
)

// MsgReplyEcho 对应 Java MSG_REPLY_ECHO (cmd=4275)
// 心跳响应
type MsgReplyEcho struct {
	A int32 // 服务器时间（含随机延迟）
}

func (m *MsgReplyEcho) Cmd() uint16 {
	return 4275
}

func (m *MsgReplyEcho) WriteBody(w *codec.GameWriter) {
	w.WriteInt(m.A)
}

var _ msg.OutMessage = (*MsgReplyEcho)(nil)
