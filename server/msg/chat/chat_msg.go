package chat_msg

import (
	"wgame-server/server/msg"
)

// ChatCard 聊天卡片
type ChatCard struct {
	CardID string `codec:"string"` // 卡片ID
}

// MsgMessageEx 对应 Java MSG_MESSAGE_EX (cmd=16383)
// 聊天消息
// 简化版本，只包含核心字段
type MsgMessageEx struct {
	Channel    int32      `codec:"int"`     // 频道类型
	ID         int32      `codec:"int"`     // 发送者ID
	Name       string     `codec:"string"`  // 发送者名称
	Msg        string     `codec:"string2"` // 消息内容
	Time       int32      `codec:"int"`     // 时间戳
	Privilege  int32      `codec:"int"`     // 权限等级
	LineName   string     `codec:"string"`  // 线路名称
	ShowExtra  int16      `codec:"short"`   // 消息类型 1=语音 0=文字
	Compress   int16      `codec:"short"`   // 压缩标识
	OrgLength  int16      `codec:"short"`   // 原始长度
	CardCount  int16      `codec:"short"`   // 卡片数量
	Cards      []ChatCard `codec:"list:short"` // 卡片列表
	VoiceTime  int32      `codec:"int"`     // 语音时长
	Token      string     `codec:"string2"` // 语音token
	Checksum   int32      `codec:"int"`     // 校验和
	FilteredMsg string    `codec:"string2"` // 过滤消息（83版本）
	MapSize    int16      `codec:"short"`   // Map大小
	ItemCookieCount int32 `codec:"int"`     // 物品Cookie数量
	TipIndex   int16      `codec:"short"`   // Tip索引
}

func (m *MsgMessageEx) Cmd() uint16 {
	return 16383
}

var _ msg.OutMessage = (*MsgMessageEx)(nil)
