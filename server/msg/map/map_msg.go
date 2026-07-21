package map_msg

import (
	"wgame-server/server/codec"
	"wgame-server/server/msg"
)

// MsgAppear 对应 Java MSG_APPEAR (cmd=16392)
// 玩家出现（进入视野）
type MsgAppear struct {
	CharID      int32  // 角色ID
	Name        string // 角色名
	Gid         string // 全局ID
	Level       int32  // 等级
	Polar       int32  // 门派
	Sex         int32  // 性别
	X           int32  // X坐标
	Y           int32  // Y坐标
	Dir         int32  // 方向
	Waiguan     int32  // 外观
	Nice        int32  // 好心值
	FashionIcon int32  // 时装图标
}

func (m *MsgAppear) Cmd() uint16 {
	return 16392
}

func (m *MsgAppear) WriteBody(w *codec.GameWriter) {
	w.WriteInt(m.CharID)
	w.WriteString(m.Name)
	w.WriteString(m.Gid)
	w.WriteInt(m.Level)
	w.WriteInt(m.Polar)
	w.WriteInt(m.Sex)
	w.WriteShort(int16(m.X))
	w.WriteShort(int16(m.Y))
	w.WriteShort(int16(m.Dir))
	w.WriteInt(m.Waiguan)
	w.WriteInt(m.Nice)
	w.WriteInt(m.FashionIcon)
}

// MsgDisappear 对应 Java MSG_DISAPPEAR (cmd=16394)
// 玩家消失（离开视野）
type MsgDisappear struct {
	CharID int32 // 角色ID
}

func (m *MsgDisappear) Cmd() uint16 {
	return 16394
}

func (m *MsgDisappear) WriteBody(w *codec.GameWriter) {
	w.WriteInt(m.CharID)
}

// MsgMapInfo 对应 Java MSG_MAP_INFO (cmd=16416)
// 地图信息
type MsgMapInfo struct {
	MapID   int32  // 地图ID
	MapName string // 地图名称
}

func (m *MsgMapInfo) Cmd() uint16 {
	return 16416
}

func (m *MsgMapInfo) WriteBody(w *codec.GameWriter) {
	w.WriteInt(m.MapID)
	w.WriteString(m.MapName)
}

var _ msg.OutMessage = (*MsgAppear)(nil)
var _ msg.OutMessage = (*MsgDisappear)(nil)
var _ msg.OutMessage = (*MsgMapInfo)(nil)
