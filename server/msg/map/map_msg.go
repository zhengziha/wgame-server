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

// MsgClearAllChar 对应 Java MSG_CLEAR_ALL_CHAR (cmd=45157)
// 清除所有角色（进入地图前发送）
type MsgClearAllChar struct {
	ID    int32 // 角色ID
	MapID int32 // 地图ID
}

func (m *MsgClearAllChar) Cmd() uint16 {
	return 45157
}

func (m *MsgClearAllChar) WriteBody(w *codec.GameWriter) {
	w.WriteInt(m.ID)
	w.WriteInt(m.MapID)
}

// MsgEnterRoom79 对应 Java MSG_ENTER_ROOM_79 (cmd=65505)
// 进入房间（包含地图信息和坐标）
type MsgEnterRoom79 struct {
	MapName          string // 地图名称
	MapShowName      string // 地图显示名称
	MapID            int32  // 地图ID
	X                int32  // X坐标
	Y                int32  // Y坐标
	Dir              int32  // 方向
	MapIndex         int32  // 地图索引
	CompactMapIndex  int32  // 压缩地图索引
	FloorIndex       int32  // 楼层索引
	WallIndex        int32  // 墙壁索引
	SafeZone         int32  // 安全区
	IsTaskWalk       int32  // 是否任务行走
	EnterEffectIndex int32  // 进入特效索引
}

func (m *MsgEnterRoom79) Cmd() uint16 {
	return 65505
}

func (m *MsgEnterRoom79) WriteBody(w *codec.GameWriter) {
	w.WriteString(m.MapName)
	w.WriteString(m.MapShowName)
	w.WriteInt(m.MapID)
	w.WriteShort(int16(m.X))
	w.WriteShort(int16(m.Y))
	w.WriteUByte(int(m.Dir))
	w.WriteInt(m.MapIndex)
	w.WriteShort(int16(m.CompactMapIndex))
	w.WriteUByte(int(m.FloorIndex))
	w.WriteUByte(int(m.WallIndex))
	w.WriteUByte(int(m.SafeZone))
	w.WriteUByte(int(m.IsTaskWalk))
	w.WriteUByte(int(m.EnterEffectIndex))
}

var _ msg.OutMessage = (*MsgClearAllChar)(nil)
var _ msg.OutMessage = (*MsgEnterRoom79)(nil)
