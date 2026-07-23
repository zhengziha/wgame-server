package map_msg

import (
	"wgame-server/server/msg"
)

// MsgAppear 对应 Java MSG_APPEAR (cmd=65529)
// 玩家出现（进入视野）
// 简化版本，只包含核心字段
type MsgAppear struct {
	ID         int32  `codec:"int"`    // 角色ID
	X          int16  `codec:"short"`  // X坐标
	Y          int16  `codec:"short"`  // Y坐标
	Dir        int16  `codec:"short"`  // 方向
	Icon       int32  `codec:"int"`    // 图标
	WeaponIcon int32  `codec:"int"`    // 武器图标
	Type       int16  `codec:"short"`  // 类型
	SubType    int32  `codec:"int"`    // 子类型
	OwnerID    int32  `codec:"int"`    // 拥有者ID
	LeaderID   int32  `codec:"int"`    // 队长ID
	Name       string `codec:"string"` // 角色名
	Level      int16  `codec:"short"`  // 等级
	Title      string `codec:"string"` // 称号
	Family     string `codec:"string"` // 家族
	Party      string `codec:"string"` // 帮派
	Status     int32  `codec:"int"`    // 状态
	SpecialIcon int32 `codec:"int"`   // 特殊图标
	OrgIcon    int32  `codec:"int"`    // 组织图标
	SuitIcon   int32  `codec:"int"`    // 时装图标
	SuitLight  int32  `codec:"int"`    // 时装光效
	GuardIcon  int32  `codec:"int"`    // 守护图标
	PetIcon    int32  `codec:"int"`    // 宠物图标
	ShadowIcon int32  `codec:"int"`    // 影子图标
	ShelterIcon int32 `codec:"int"`   // 庇护图标
	MountIcon  int32  `codec:"int"`    // 坐骑图标
	AliName    string `codec:"string"` // 别名
	Gid        string `codec:"string"` // 全局ID
	Camp       string `codec:"string"` // 阵营
	VipType    int8   `codec:"byte"`   // VIP类型
	IsHide     int8   `codec:"byte"`   // 是否隐身
	MoveSpeed  int8   `codec:"byte"`   // 移动速度百分比
	Score      int32  `codec:"int"`    // 分数
	Opacity    int8   `codec:"byte"`   // 不透明度
	Masquerade int32  `codec:"int"`    // 伪装
	UpgradeState int8 `codec:"byte"`  // 飞升状态
	UpgradeType int8  `codec:"byte"`   // 飞升类型
	Obstacle   int8   `codec:"byte"`   // 障碍物
	EffectCount int16 `codec:"short"` // 特效图标数量
	Effects    []int32 `codec:"list:short"` // 特效图标列表
	ShareMountIcon int32 `codec:"int"`  // 共享坐骑图标
	ShareMountLeaderID int32 `codec:"int"` // 共享坐骑队长ID
	ShareMountShadow int32 `codec:"int"` // 共享坐骑影子
	GatherCount int16 `codec:"short"` // 聚集数量
	GatherMountIcons []int32 `codec:"list:short"` // 聚集坐骑图标
	GatherNameNum int16 `codec:"short"` // 聚集名字数量
	GatherNames []string `codec:"list:short"` // 聚集名字列表
	Portrait   int32  `codec:"int"`    // 头像
	CustomIcon string `codec:"string"` // 自定义图标
	TeamIcon   int16  `codec:"short"`  // 队伍图标
	ExtraScale int16  `codec:"short"`  // 额外缩放
	GatherSuitIcons int16 `codec:"short"` // 聚集时装图标（固定0）
	BanRule    string `codec:"string"` // 封禁规则
	TitleBanRule string `codec:"string"` // 称号封禁规则
	XOffset    int8   `codec:"byte"`   // X偏移
	YOffset    int8   `codec:"byte"`   // Y偏移
	MoveType   int8   `codec:"byte"`   // 移动类型
	FlyType    int8   `codec:"byte"`   // 飞行类型
	MoveIDCount int16 `codec:"short"`  // 移动ID数量
	MoveIDs    []int32 `codec:"list:short"` // 移动ID列表
}

func (m *MsgAppear) Cmd() uint16 {
	return 65529
}

// MsgDisappear 对应 Java MSG_DISAPPEAR (cmd=12285)
// 玩家消失（离开视野）
// Java写入顺序：id(int), type(byte)
type MsgDisappear struct {
	CharID int32 `codec:"int"`  // 角色ID
	Type   int8  `codec:"byte"` // 类型 1=玩家, 其他=NPC等
}

func (m *MsgDisappear) Cmd() uint16 {
	return 12285
}

// MsgMapInfo 对应 Java MSG_ENTER_ROOM (cmd=65505)
// 地图信息
// Java写入顺序：map_name(String), zeros(short), zerob(byte), map_id(short), x1(byte), x2(byte), y1(byte), y2(byte), map_show_name(String), is_safe_zone(byte), is_task_walk(byte), wall_index(byte), enter_effect_index(short)
type MsgMapInfo struct {
	MapName         string `codec:"string"` // 地图名称
	Zeros           int16  `codec:"short"`  // 保留（0）
	ZeroB           int8   `codec:"byte"`   // 保留（0）
	MapID           int16  `codec:"short"`  // 地图ID
	X1              int8   `codec:"byte"`   // 坐标X1
	X2              int8   `codec:"byte"`   // 坐标X2
	Y1              int8   `codec:"byte"`   // 坐标Y1
	Y2              int8   `codec:"byte"`   // 坐标Y2
	MapShowName     string `codec:"string"` // 地图显示名称
	IsSafeZone      int8   `codec:"byte"`   // 是否安全区
	IsTaskWalk      int8   `codec:"byte"`   // 是否任务行走
	WallIndex       int8   `codec:"byte"`   // 墙壁索引
	EnterEffectIndex int16 `codec:"short"` // 进入特效索引
}

func (m *MsgMapInfo) Cmd() uint16 {
	return 65505
}

// MsgClearAllChar 对应 Java MSG_CLEAR_ALL_CHAR (cmd=45157)
// 清除所有角色（进入地图前发送）
type MsgClearAllChar struct {
	ID    int32 `codec:"int"` // 角色ID
	MapID int32 `codec:"int"` // 地图ID
}

func (m *MsgClearAllChar) Cmd() uint16 {
	return 45157
}

// MsgEnterRoom79 对应 Java MSG_ENTER_ROOM_79 (cmd=65505)
// 进入房间（包含地图信息和坐标）
// Java写入顺序：map_name(String), map_show_name(String), map_id(int), x(short), y(short), dir(byte), map_index(int), compact_map_index(short), floor_index(byte), wall_index(byte), safe_zone(byte), is_task_walk(byte), enter_effect_index(byte)
type MsgEnterRoom79 struct {
	MapName          string `codec:"string"` // 地图名称
	MapShowName      string `codec:"string"` // 地图显示名称
	MapID            int32  `codec:"int"`    // 地图ID
	X                int16  `codec:"short"`  // X坐标
	Y                int16  `codec:"short"`  // Y坐标
	Dir              int8   `codec:"byte"`   // 方向
	MapIndex         int32  `codec:"int"`    // 地图索引
	CompactMapIndex  int16  `codec:"short"`  // 压缩地图索引
	FloorIndex       int8   `codec:"byte"`   // 楼层索引
	WallIndex        int8   `codec:"byte"`   // 墙壁索引
	SafeZone         int8   `codec:"byte"`   // 安全区
	IsTaskWalk       int8   `codec:"byte"`   // 是否任务行走
	EnterEffectIndex int8   `codec:"byte"`   // 进入特效索引
}

func (m *MsgEnterRoom79) Cmd() uint16 {
	return 65505
}

// MsgMoved 对应 Java MSG_MOVED (cmd=16431)
// 玩家移动消息
// Java写入顺序：x(short), y(short), id(int), dir(byte)
type MsgMoved struct {
	X   int16 `codec:"short"` // X坐标
	Y   int16 `codec:"short"` // Y坐标
	ID  int32 `codec:"int"`   // 角色ID
	Dir int8  `codec:"byte"`  // 方向
}

func (m *MsgMoved) Cmd() uint16 {
	return 16431
}

var _ msg.OutMessage = (*MsgAppear)(nil)
var _ msg.OutMessage = (*MsgDisappear)(nil)
var _ msg.OutMessage = (*MsgMapInfo)(nil)
var _ msg.OutMessage = (*MsgClearAllChar)(nil)
var _ msg.OutMessage = (*MsgEnterRoom79)(nil)
var _ msg.OutMessage = (*MsgMoved)(nil)
