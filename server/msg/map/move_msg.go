package map_msg

import (
	"wgame-server/server/msg"
)

// MsgTeamMoved 对应 Java MSG_TEAM_MOVED (cmd=16430)
// 队伍移动消息
// Java写入顺序：id(short), y(short), id(int), dir(byte) - 实际为 id, x, y, map_id
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

var _ msg.OutMessage = (*MsgTeamMoved)(nil)
var _ msg.OutMessage = (*MsgTeleportFailed)(nil)
