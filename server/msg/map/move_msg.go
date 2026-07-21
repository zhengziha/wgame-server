package map_msg

import (
	"wgame-server/server/codec"
	"wgame-server/server/msg"
)

// MsgMoved 对应 Java MSG_MOVED (cmd=16432)
// 玩家移动消息
type MsgMoved struct {
	ID  int32 // 角色ID
	X   int32 // X坐标
	Y   int32 // Y坐标
	Dir int32 // 方向
}

func (m *MsgMoved) Cmd() uint16 {
	return 16432
}

func (m *MsgMoved) WriteBody(w *codec.GameWriter) {
	w.WriteInt(m.ID)
	w.WriteShort(int16(m.X))
	w.WriteShort(int16(m.Y))
	w.WriteShort(int16(m.Dir))
}

// MsgTeamMoved 对应 Java MSG_TEAM_MOVED (cmd=16430)
// 队伍移动消息
type MsgTeamMoved struct {
	ID    int32 // 角色ID
	X     int32 // X坐标
	Y     int32 // Y坐标
	MapID int32 // 地图ID
}

func (m *MsgTeamMoved) Cmd() uint16 {
	return 16430
}

func (m *MsgTeamMoved) WriteBody(w *codec.GameWriter) {
	w.WriteInt(m.ID)
	w.WriteShort(int16(m.X))
	w.WriteShort(int16(m.Y))
	w.WriteInt(m.MapID)
}

var _ msg.OutMessage = (*MsgMoved)(nil)
var _ msg.OutMessage = (*MsgTeamMoved)(nil)
