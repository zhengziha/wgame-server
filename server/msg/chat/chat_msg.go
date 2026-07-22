package chat_msg

import (
	"wgame-server/server/codec"
	"wgame-server/server/msg"
)

// MsgMessageEx 对应 Java MSG_MESSAGE_EX (cmd=16383)
// 聊天消息
type MsgMessageEx struct {
	Channel    int32  // 频道类型
	ID         int32  // 发送者ID
	Name       string // 发送者名称
	Msg        string // 消息内容
	Time       int32  // 时间戳
	Privilege  int32  // 权限等级
	LineName   string // 线路名称
	ShowExtra  int16  // 消息类型 1=语音 0=文字
	Compress   int16  // 压缩标识
	OrgLength  int16  // 原始长度
	CardCount  int16  // 卡片数量
	CardId     string // 卡片ID
	VoiceTime  int32  // 语音时长
	Token      string // 语音token
	Checksum   int32  // 校验和
}

func (m *MsgMessageEx) Cmd() uint16 {
	return 16383
}

func (m *MsgMessageEx) WriteBody(w *codec.GameWriter) {
	w.WriteInt(m.Channel)
	w.WriteInt(m.ID)
	w.WriteString(m.Name)
	w.WriteString2(m.Msg)
	w.WriteInt(m.Time)
	w.WriteInt(m.Privilege)
	w.WriteString(m.LineName)
	w.WriteShort(m.ShowExtra)
	w.WriteShort(m.Compress)
	w.WriteShort(m.OrgLength)
	w.WriteShort(m.CardCount)
	for i := int16(0); i < m.CardCount; i++ {
		w.WriteString(m.CardId)
	}
	w.WriteInt(m.VoiceTime)
	w.WriteString2(m.Token)
	w.WriteInt(m.Checksum)
	w.WriteString2("") // filtered_msg (83版本)
	w.WriteShort(0)    // map.size
	w.WriteInt(0)      // nItemCookieCount
	w.WriteShort(0)    // tip_index
}

var _ msg.OutMessage = (*MsgMessageEx)(nil)