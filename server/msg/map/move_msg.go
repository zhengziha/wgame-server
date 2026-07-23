package map_msg

import (
	"wgame-server/server/msg"
)

// MsgTeamMoved 对应 Java MSG_TEAM_MOVED (cmd=16430)
// 队伍移动消息
// Java写入顺序：id(short), y(short), id(int), dir(byte) - 实际为 id, x, y, map_id
//
// 注意：Go 结构体字段的 codec 标签决定了序列化时的类型
// codec:"int"   -> 序列化写入 4 字节 int，类似 Java 的 writeInt()
// codec:"short" -> 序列化写入 2 字节 short，类似 Java 的 writeShort()
// codec:"string" -> 序列化写入字符串（带长度前缀），类似 Java 的 writeString()
type MsgTeamMoved struct {
	ID    int32 `codec:"int"`   // 角色ID
	X     int16 `codec:"short"` // X坐标
	Y     int16 `codec:"short"` // Y坐标
	MapID int32 `codec:"int"`   // 地图ID
}

func (m *MsgTeamMoved) Cmd() uint16 {
	return 16429
}

// MsgTeleportFailed 对应 Java MSG_TELEPORT_FAILED (cmd=61476)
// 传送失败消息
// Java写入顺序：msg(String)
type MsgTeleportFailed struct {
	Msg string `codec:"string"` // 失败原因
}

func (m *MsgTeleportFailed) Cmd() uint16 {
	return 53523
}

// 编译期接口检查：确保消息类型实现了 OutMessage 接口
// 相当于 Java 中的 implements 检查
var _ msg.OutMessage = (*MsgTeamMoved)(nil)
var _ msg.OutMessage = (*MsgTeleportFailed)(nil)
