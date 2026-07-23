package auth

import (
	"wgame-server/server/msg"
)

// MsgCheckUserData 对应 Java MSG_L_CHECK_USER_DATA (cmd=11012)
// 检查用户数据消息
type MsgCheckUserData struct {
	Result int32  `codec:"int"`    // 结果，固定为1
	Cookie string `codec:"string"` // Cookie值，固定为"47Q60635Q22"
}

func (m *MsgCheckUserData) Cmd() uint16 {
	return 11012
}

var _ msg.OutMessage = (*MsgCheckUserData)(nil)
