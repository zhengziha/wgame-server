package system

import (
	"wgame-server/server/msg"
)

// MsgCharAlreadyLogin 对应 Java MSG_CHAR_ALREADY_LOGIN (cmd=45121)
// 通知客户端角色已登录
// Java写入顺序：roleName(String)
type MsgCharAlreadyLogin struct {
	Name string `codec:"string"` // 角色名
}

func (m *MsgCharAlreadyLogin) Cmd() uint16 {
	return 45121
}

// MsgCsServerType 对应 Java MSG_CS_SERVER_TYPE (cmd=45277)
// 更新服务器类型
// Java写入顺序：server_type(int)
type MsgCsServerType struct {
	ServerType int32 `codec:"int"` // 服务器类型
}

func (m *MsgCsServerType) Cmd() uint16 {
	return 45277
}

// MsgSetSetting 对应 Java MSG_SET_SETTING (cmd=61589)
// 系统设置，包含62个键值对
// Java写入顺序：key0(String), settingkey0(int), key1(String), settingkey1(int), ..., key61(String), settingkey61(int)
type MsgSetSettingItem struct {
	Key   string `codec:"string"`
	Value int32  `codec:"int"`
}

type MsgSetSetting struct {
	Items []MsgSetSettingItem `codec:"list:short"` // 设置项列表
}

func (m *MsgSetSetting) Cmd() uint16 {
	return 61589
}

// NewMsgSetSetting 创建默认系统设置
func NewMsgSetSetting() *MsgSetSetting {
	items := make([]MsgSetSettingItem, 62)
	for i := 0; i < 62; i++ {
		items[i] = MsgSetSettingItem{
			Key:   "",
			Value: 0,
		}
	}
	return &MsgSetSetting{
		Items: items,
	}
}

// MsgUpdate 对应 Java MSG_UPDATE (cmd=65527)
// 角色属性更新消息。Java写入顺序：id(int), 然后 BuildField 数量 + 各属性。
// 这里把所有 BuildField 属性封装到 MsgUpdateProps 子结构体，用 codec:"buildfields" 标记，
// 序列化时自动写 short(字段数) + 展开 bf 字段。
type MsgUpdate struct {
	Id    int32          `codec:"int"`         // 角色ID
	Props MsgUpdateProps `codec:"buildfields"` // BuildField 属性容器
}

// MsgUpdateProps 封装 MSG_UPDATE 的所有 BuildField 属性。
// 内部只允许 bf 字段，字段数编译期固定。
type MsgUpdateProps struct {
	// 基础属性
	Name        string `codec:"bf:Name"`           // key=1 角色名
	Level       int32  `codec:"bf:Level"`          // key=31 等级
	Exp         int32  `codec:"bf:Exp"`            // key=5 经验
	ExpToNextLv int32  `codec:"bf:ExpToNextLevel"` // key=55 升级所需经验

	// 气血和法力
	Life      int32 `codec:"bf:Life"`      // key=7 生命
	MaxLife   int32 `codec:"bf:MaxLife"`   // key=22 最大生命
	Mana      int32 `codec:"bf:Mana"`      // key=8 法力
	MaxMana   int32 `codec:"bf:MaxMana"`   // key=23 最大法力
	ExtraLife int32 `codec:"bf:ExtraLife"` // key=74 额外生命
	ExtraMana int32 `codec:"bf:ExtraMana"` // key=75 额外法力

	// 五行属性
	Metal int32 `codec:"bf:Metal"` // key=9 金
	Wood  int32 `codec:"bf:Wood"`  // key=10 木
	Water int32 `codec:"bf:Water"` // key=11 水
	Fire  int32 `codec:"bf:Fire"`  // key=12 火
	Earth int32 `codec:"bf:Earth"` // key=13 土
	Polar int32 `codec:"bf:Polar"` // key=44 五行

	// 抗性
	ResistMetal int32 `codec:"bf:ResistMetal"` // key=14 抗金
	ResistWood  int32 `codec:"bf:ResistWood"`  // key=15 抗木
	ResistWater int32 `codec:"bf:ResistWater"` // key=16 抗水
	ResistFire  int32 `codec:"bf:ResistFire"`  // key=17 抗火
	ResistEarth int32 `codec:"bf:ResistEarth"` // key=18 抗土

	// 基础属性
	Str int32 `codec:"bf:Str"` // key=2 力量
	Con int32 `codec:"bf:Con"` // key=3 体质
	Dex int32 `codec:"bf:Dex"` // key=4 敏捷
	Wiz int32 `codec:"bf:Wiz"` // key=6 智力

	// 战斗属性
	PhyPower   int32 `codec:"bf:PhyPower"`   // key=19 物理攻击
	MagPower   int32 `codec:"bf:MagPower"`   // key=20 法术攻击
	Def        int32 `codec:"bf:Def"`        // key=21 防御
	Speed      int32 `codec:"bf:Speed"`      // key=24 速度
	Stamina    int32 `codec:"bf:Stamina"`    // key=25 耐力
	MaxStamina int32 `codec:"bf:MaxStamina"` // key=26 最大耐力

	// 金钱
	Cash     int32 `codec:"bf:Cash"`     // key=28 金币
	Balance  int32 `codec:"bf:Balance"`  // key=29 元宝
	GoldCoin int64 `codec:"bf:GoldCoin"` // key=38 元宝(long)

	// 角色信息
	Icon       int32 `codec:"bf:Icon"`       // key=40 图标
	Portrait   int32 `codec:"bf:Portrait"`   // key=86 头像
	Nice       int32 `codec:"bf:Nice"`       // key=47 好心值
	Reputation int32 `codec:"bf:Reputation"` // key=48 声望
	Tao        int32 `codec:"bf:Tao"`        // key=46 道行
	TaoEx      int32 `codec:"bf:TaoEx"`      // key=53 道行(ex)

	// 师徒系统
	Master string `codec:"bf:Master"` // key=30 师傅名
	Couple int32  `codec:"bf:Couple"` // key=54 配偶ID

	// 帮派信息
	PartyId      int32  `codec:"bf:PartyId"`      // key=55 帮派ID
	PartyName    string `codec:"bf:PartyName"`    // key=56 帮派名
	PartyContrib int32  `codec:"bf:PartyContrib"` // key=57 帮派贡献

	// 删除相关
	ToBeDeleted      int32 `codec:"bf:ToBeDeleted"`      // key=262 是否待删除
	LeftTimeToDelete int32 `codec:"bf:LeftTimeToDelete"` // key=263 剩余删除时间

	// 在线状态
	Online int32 `codec:"bf:Online"` // key=36 在线状态

	// 升级系统
	UpgradeState       int32 `codec:"bf:UpgradeState"`          // key=340 升级状态
	UpgradeLevel       int32 `codec:"bf:UpgradeLevel"`          // key=342 升级等级
	UpgradeExp         int32 `codec:"bf:UpgradeExp"`            // key=343 升级经验
	UpgradeExpToNextLv int32 `codec:"bf:UpgradeExpToNextLevel"` // key=344 升级到下一级经验
}

func (m *MsgUpdate) Cmd() uint16 {
	return 65527
}

// NewMsgUpdate 创建角色属性更新消息
func NewMsgUpdate(id int32) *MsgUpdate {
	return &MsgUpdate{
		Id: id,
	}
}

// NewMsgUpdateFromChara 从 Chara 创建角色属性更新消息
func NewMsgUpdateFromChara(chara interface{}) *MsgUpdate {
	msg := &MsgUpdate{}
	switch c := chara.(type) {
	case interface {
		GetID() int32
		GetName() string
		GetLevel() int32
	}:
		msg.Id = c.GetID()
		msg.Props.Name = c.GetName()
		msg.Props.Level = c.GetLevel()
	}
	return msg
}

// MsgGeneralNotify 对应 Java MSG_GENERAL_NOTIFY (cmd=9129)
// 通用通知消息
// Java写入顺序：notify(int), para(String)
type MsgGeneralNotify struct {
	Notify int32  `codec:"int"`    // 通知ID
	Para   string `codec:"string"` // 参数
}

func (m *MsgGeneralNotify) Cmd() uint16 {
	return 9129
}

// MsgSetPushSettings 对应 Java MSG_SET_PUSH_SETTINGS (cmd=53399)
// 通知推送开关信息
// Java写入顺序：value(String)
type MsgSetPushSettings struct {
	Value string `codec:"string"` // 推送设置值
}

func (m *MsgSetPushSettings) Cmd() uint16 {
	return 53399
}

// MsgFuzzyIdentity 对应 Java MSG_FUZZY_IDENTITY (cmd=53417)
// 模糊身份验证
// Java写入顺序：isBindName(byte), isBindPhone(byte), bindName(String), bindId(String), bindPhone(String)
type MsgFuzzyIdentity struct {
	IsBindName  int8   `codec:"byte"`   // 是否绑定名字
	IsBindPhone int8   `codec:"byte"`   // 是否绑定手机
	BindName    string `codec:"string"` // 绑定名字
	BindId      string `codec:"string"` // 绑定ID
	BindPhone   string `codec:"string"` // 绑定手机
}

func (m *MsgFuzzyIdentity) Cmd() uint16 {
	return 53417
}

// MsgExecuteLuaCode 对应 Java MSG_EXECUTE_LUA_CODE (cmd=53605)
// 通知客户端执行一段Lua代码
// Java写入顺序：cookie(int), code(String2), flag(byte)
type MsgExecuteLuaCode struct {
	Cookie int32  `codec:"int"`     // Cookie值
	Code   string `codec:"string2"` // Lua代码
	Flag   int8   `codec:"byte"`    // 标志
}

func (m *MsgExecuteLuaCode) Cmd() uint16 {
	return 53605
}

// 确保实现 msg.OutMessage 接口
var _ msg.OutMessage = (*MsgCharAlreadyLogin)(nil)
var _ msg.OutMessage = (*MsgCsServerType)(nil)
var _ msg.OutMessage = (*MsgSetSetting)(nil)
var _ msg.OutMessage = (*MsgUpdate)(nil)
var _ msg.OutMessage = (*MsgGeneralNotify)(nil)
var _ msg.OutMessage = (*MsgSetPushSettings)(nil)
var _ msg.OutMessage = (*MsgFuzzyIdentity)(nil)
var _ msg.OutMessage = (*MsgExecuteLuaCode)(nil)
