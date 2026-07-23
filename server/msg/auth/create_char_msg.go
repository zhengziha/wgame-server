package auth

import (
	"wgame-server/server/msg"
)

// MsgCreateNewChar 对应 Java MSG_CREATE_NEW_CHAR (cmd=8285)
// 创建角色成功后返回
// Java写入顺序：gid(String), name(String)
type MsgCreateNewChar struct {
	Gid  string `codec:"string"` // 全局唯一ID
	Name string `codec:"string"` // 角色名
}

func (m *MsgCreateNewChar) Cmd() uint16 {
	return 8285
}

var _ msg.OutMessage = (*MsgCreateNewChar)(nil)
