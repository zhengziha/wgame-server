package model

import "time"

// Characters 角色主表（玩家核心数据）。
// level 为 SQL 保留字，column tag 直接写列名，GORM 会自动加反引号。
// chargeTotal/lastMonTao/lastDailyReset/lastWeeklyReset/otherGoods 在 Java 实体中声明，
// wd-game-18.sql 建表语句未出现这些列，按普通列保留以兼容 Java 逻辑。
type Characters struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// Name 角色名
	Name string `gorm:"size:255;column:name" json:"name"`

	// Level 真身等级（列名 level 为 SQL 保留字）
	Level int32 `gorm:"column:level" json:"level"`

	// Polar 门派
	Polar int32 `gorm:"column:polar" json:"polar"`

	// Sex 性别（1女 2男）
	Sex int32 `gorm:"column:sex" json:"sex"`

	// Portrait 头像
	Portrait int32 `gorm:"column:portrait" json:"portrait"`

	// Gid 用户唯一 id
	Gid string `gorm:"size:255;column:gid" json:"gid"`

	// ChargeScore 充值积分
	ChargeScore int32 `gorm:"column:charge_score" json:"chargeScore"`

	// GoldCoin 金元宝
	GoldCoin int32 `gorm:"column:gold_coin" json:"goldCoin"`

	// ChargeTotal 累计充值
	ChargeTotal int32 `gorm:"column:charge_total" json:"chargeTotal"`

	// MapId 地图 id
	MapId int32 `gorm:"column:map_id" json:"mapId"`

	// MapName 地图名字
	MapName string `gorm:"size:15;column:map_name" json:"mapName"`

	// X 坐标 x
	X int32 `gorm:"column:x" json:"x"`

	// Y 坐标 y
	Y int32 `gorm:"column:y" json:"y"`

	// MonthTao 月道行（原元婴等级）
	MonthTao int32 `gorm:"column:month_tao" json:"monthTao"`

	// LastMonTao 上月道行（独立列，周维度全表原子归档 monthTao）
	LastMonTao int32 `gorm:"column:last_mon_tao" json:"lastMonTao"`

	// LastDailyReset 日重置水位 yyyyMMdd，懒重置用，独立列
	LastDailyReset int32 `gorm:"column:last_daily_reset" json:"lastDailyReset"`

	// LastWeeklyReset 周重置水位 所在周一 yyyyMMdd，懒重置用，独立列
	LastWeeklyReset int32 `gorm:"column:last_weekly_reset" json:"lastWeeklyReset"`

	// AccountId 账号 id
	AccountId int32 `gorm:"column:account_id" json:"accountId"`

	// AddTime 注册时间
	AddTime time.Time `gorm:"autoCreateTime;column:add_time" json:"addTime"`

	// UpdateTime 更新时间
	UpdateTime time.Time `gorm:"autoUpdateTime;column:update_time" json:"updateTime"`

	// Deleted 逻辑删除标记（tinyint(1)）
	Deleted bool `gorm:"column:deleted" json:"deleted"`

	// LastLoginTime 最后登录时间
	LastLoginTime int32 `gorm:"column:last_login_time" json:"lastLoginTime"`

	// Online 是否在线（0:离线 1:在线）
	Online int32 `gorm:"column:online" json:"online"`

	// LastLoginIp 最后一次登录 ip
	LastLoginIp string `gorm:"size:255;column:last_login_ip" json:"lastLoginIp"`

	// Block 封号标识（1开启 0关闭）
	Block int32 `gorm:"column:block" json:"block"`

	// Xiaozi 问道小子（0:默认 1:是）
	Xiaozi int32 `gorm:"column:xiaozi" json:"xiaozi"`

	// Data 角色数据 JSON
	Data string `gorm:"type:mediumtext;column:data" json:"data"`

	// BaseAttr 擂台基础属性 JSON
	BaseAttr string `gorm:"type:text;column:base_attr" json:"baseAttr"`

	// ZbAttributeJson 装备属性 JSON
	ZbAttributeJson string `gorm:"type:text;column:zb_attribute_json" json:"zbAttributeJson"`

	// Cangku 仓库
	Cangku string `gorm:"type:mediumtext;column:cangku" json:"cangku"`

	// OtherGoods 其它物品 JSON
	OtherGoods string `gorm:"type:mediumtext;column:other_goods" json:"otherGoods"`

	// Backpack 背包
	Backpack string `gorm:"type:mediumtext;column:backpack" json:"backpack"`

	// PetStore 宠物仓库
	PetStore string `gorm:"type:mediumtext;column:pet_store" json:"petStore"`

	// Texiao 特效
	Texiao string `gorm:"type:longtext;column:texiao" json:"texiao"`

	// Genchong 根虫
	Genchong string `gorm:"type:longtext;column:genchong" json:"genchong"`

	// Listshouhu 守护列表
	Listshouhu string `gorm:"type:mediumtext;column:listshouhu" json:"listshouhu"`

	// Shizhuang 时装
	Shizhuang string `gorm:"type:mediumtext;column:shizhuang" json:"shizhuang"`

	// CardStore 卡套
	CardStore string `gorm:"type:mediumtext;column:card_store" json:"cardStore"`

	// CustomShizhuang 自定义时装
	CustomShizhuang string `gorm:"type:mediumtext;column:custom_shizhuang" json:"customShizhuang"`
}

// TableName 显式指定表名（characters 为复数）
func (Characters) TableName() string {
	return "characters"
}
