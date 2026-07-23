package system

import (
	"time"
	"wgame-server/server/msg"
)

// MsgReplyServerTime 对应 Java MSG_REPLY_SERVER_TIME (cmd=41009)
// 服务器时间消息
// Java写入顺序：server_time(int), open_flag(int), time_zone(byte), ip(String)
type MsgReplyServerTime struct {
	ServerTime int32  `codec:"int"`    // 服务器时间
	OpenFlag   int32  `codec:"int"`    // 第二套加点方案开启标记
	TimeZone   int8   `codec:"byte"`   // 时区
	IP         string `codec:"string"` // IP地址
}

func (m *MsgReplyServerTime) Cmd() uint16 {
	return 41009
}

func NewMsgReplyServerTime() *MsgReplyServerTime {
	return &MsgReplyServerTime{
		ServerTime: int32(time.Now().Unix()),
		OpenFlag:   1,
		TimeZone:   8,
		IP:         "",
	}
}

var _ msg.OutMessage = (*MsgReplyServerTime)(nil)
