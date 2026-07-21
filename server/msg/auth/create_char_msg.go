package auth

import (
	"wgame-server/server/codec"
	"wgame-server/server/msg"
)

// MsgCreateNewChar 对应 Java MSG_CREATE_NEW_CHAR (cmd=8285)
// 创建角色成功后返回
type MsgCreateNewChar struct {
	Name string // 角色名
	Gid  string // 全局唯一ID
}

func (m *MsgCreateNewChar) Cmd() uint16 {
	return 8285
}

func (m *MsgCreateNewChar) WriteBody(w *codec.GameWriter) {
	w.WriteString(m.Name)
	w.WriteString(m.Gid)
}

var _ msg.OutMessage = (*MsgCreateNewChar)(nil)
