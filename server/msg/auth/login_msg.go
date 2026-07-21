package auth

import (
	"wgame-server/server/codec"
	"wgame-server/server/msg"
)

// MsgAgentResult 对应 Java MSG_L_AGENT_RESULT (cmd=9041)
// 账号认证成功后返回给客户端
type MsgAgentResult struct {
	Result       int32  // 结果 1=成功
	ID           int32  // 账号ID
	Privilege    int32  // 权限等级
	IP           string // 服务器IP
	Port         int32  // 服务器端口
	ServerName   string // 服务器名称
	ServerStatus int32  // 服务器状态
	Msg          string // 消息
}

func (m *MsgAgentResult) Cmd() uint16 {
	return 9041
}

func (m *MsgAgentResult) WriteBody(w *codec.GameWriter) {
	w.WriteInt(m.Result)
	w.WriteInt(m.ID)
	w.WriteInt(m.Privilege)
	w.WriteString(m.IP)
	w.WriteInt(m.Port)
	w.WriteString(m.ServerName)
	w.WriteInt(m.ServerStatus)
	w.WriteString(m.Msg)
}

// MsgAuth 对应 Java MSG_L_AUTH (cmd=9042)
// 账号认证失败时返回给客户端
type MsgAuth struct {
	Msg string // 错误消息
}

func (m *MsgAuth) Cmd() uint16 {
	return 9042
}

func (m *MsgAuth) WriteBody(w *codec.GameWriter) {
	w.WriteString(m.Msg)
}

// MsgExistedCharList 对应 Java MSG_EXISTED_CHAR_LIST (cmd=61538)
// 返回账号下的角色列表
type MsgExistedCharList struct {
	AccountOnline int32           // 账号是否在线
	VoList        []*VoExistedChar // 角色列表
}

func (m *MsgExistedCharList) Cmd() uint16 {
	return 61538
}

func (m *MsgExistedCharList) WriteBody(w *codec.GameWriter) {
	w.WriteInt(m.AccountOnline)
	w.WriteUShort(uint16(len(m.VoList)))
	for _, vo := range m.VoList {
		vo.WriteBody(w)
	}
}

// VoExistedChar 对应 Java Vo_61537_0
type VoExistedChar struct {
	CharID          int32  // 角色ID
	Name            string // 角色名
	Level           int32  // 等级
	Polar           int32  // 门派
	Sex             int32  // 性别
	OnlineState     int32  // 在线状态 0=离线, 1=在线
	FashionIcon     int32  // 时装图标
	UpgradeLevel    int32  // 飞升等级
	PetIcon         int32  // 宠物图标
	MountIcon       int32  // 坐骑图标
	SpecialIcon     int32  // 特殊图标
	GenchongIcon    int32  // 跟宠图标
	UpgradeType     int32  // 飞升类型
	Nice            int32  // 好心值
	WeeklyLoginDays int32  // 本周登录天数
	IsFeisheng      int32  // 是否飞升
	Tao             int32  // 道行(天)
	Gid             string // 全局唯一ID
	MapID           int32  // 所在地图ID
	MapName         string // 所在地图名称
	Line            int32  // 线路号
	X               int32  // X坐标
	Y               int32  // Y坐标
	PartyName       string // 帮派名称
	Family          string // 家族名称
	Title           string // 称号
}

func (v *VoExistedChar) WriteBody(w *codec.GameWriter) {
	w.WriteInt(v.CharID)
	w.WriteString(v.Name)
	w.WriteInt(v.Level)
	w.WriteInt(v.Polar)
	w.WriteInt(v.Sex)
	w.WriteInt(v.OnlineState)
	w.WriteInt(v.FashionIcon)
	w.WriteInt(v.UpgradeLevel)
	w.WriteInt(v.PetIcon)
	w.WriteInt(v.MountIcon)
	w.WriteInt(v.SpecialIcon)
	w.WriteInt(v.GenchongIcon)
	w.WriteInt(v.UpgradeType)
	w.WriteInt(v.Nice)
	w.WriteInt(v.WeeklyLoginDays)
	w.WriteInt(v.IsFeisheng)
	w.WriteInt(v.Tao)
	w.WriteString(v.Gid)
	w.WriteInt(v.MapID)
	w.WriteString(v.MapName)
	w.WriteInt(v.Line)
	w.WriteInt(v.X)
	w.WriteInt(v.Y)
	w.WriteString(v.PartyName)
	w.WriteString(v.Family)
	w.WriteString(v.Title)
}

// MsgKickOff 对应 Java MSG_KICK_OFF (cmd=13142)
// 踢下线消息
type MsgKickOff struct {
	Msg string // 踢下线原因
}

func (m *MsgKickOff) Cmd() uint16 {
	return 13142
}

func (m *MsgKickOff) WriteBody(w *codec.GameWriter) {
	w.WriteString(m.Msg)
}

// MsgShowReconnectPara 对应 Java MSG_SHOW_RECONNECT_PARA (cmd=21260)
// 重连参数
type MsgShowReconnectPara struct {
	IP       string // 服务器IP
	Port     int32  // 端口
	AuthKey  int32  // 认证key
	Seed     int32  // 种子
}

func (m *MsgShowReconnectPara) Cmd() uint16 {
	return 21260
}

func (m *MsgShowReconnectPara) WriteBody(w *codec.GameWriter) {
	w.WriteString(m.IP)
	w.WriteInt(m.Port)
	w.WriteInt(m.AuthKey)
	w.WriteInt(m.Seed)
}

// 确保实现 msg.OutMessage 接口
var _ msg.OutMessage = (*MsgAgentResult)(nil)
var _ msg.OutMessage = (*MsgAuth)(nil)
var _ msg.OutMessage = (*MsgExistedCharList)(nil)
var _ msg.OutMessage = (*MsgKickOff)(nil)
var _ msg.OutMessage = (*MsgShowReconnectPara)(nil)
