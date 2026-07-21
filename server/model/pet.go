package model

import "time"

// Pet 宠物表。
// Index 列名为 SQL 保留字，GORM 会自动加反引号。
// ExchangeBrand、LimitTime 在 Java 实体中声明但未出现在 wd-game-18.sql 建表语句中，按普通列保留。
type Pet struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// Index 对应图鉴的位置（列名 index 为 SQL 保留字）
	Index int32 `gorm:"column:index" json:"index"`

	// PetType 宠物类型（0-1野怪,2宝宝,3变异,4神兽,6鬼卒 等）
	PetType int32 `gorm:"column:pet_type" json:"petType"`

	// Level 阶数
	Level int32 `gorm:"column:level" json:"level"`

	// LevelReq 携带等级
	LevelReq int32 `gorm:"column:level_req" json:"levelReq"`

	// Life 血气成长
	Life int32 `gorm:"column:life" json:"life"`

	// Mana 法力成长
	Mana int32 `gorm:"column:mana" json:"mana"`

	// Speed 速度成长
	Speed int32 `gorm:"column:speed" json:"speed"`

	// PhyAttack 物攻成长
	PhyAttack int32 `gorm:"column:phy_attack" json:"phyAttack"`

	// MagAttack 法功成长
	MagAttack int32 `gorm:"column:mag_attack" json:"magAttack"`

	// Polar 金木水火土
	Polar string `gorm:"size:255;column:polar" json:"polar"`

	// Skiils 拥有天生技能（注：沿用 Java 字段拼写）
	Skiils string `gorm:"size:255;column:skiils" json:"skiils"`

	// Zoon 所在地图
	Zoon string `gorm:"size:255;column:zoon" json:"zoon"`

	// Icon 外观
	Icon int32 `gorm:"column:icon" json:"icon"`

	// AddTime 创建时间
	AddTime time.Time `gorm:"autoCreateTime;column:add_time" json:"addTime"`

	// UpdateTime 更新时间
	UpdateTime time.Time `gorm:"autoUpdateTime;column:update_time" json:"updateTime"`

	// Deleted 逻辑删除标记（tinyint(1)）
	Deleted bool `gorm:"column:deleted" json:"deleted"`

	// Name 名字
	Name string `gorm:"size:255;column:name" json:"name"`

	// SkillRange 技能范围
	SkillRange int32 `gorm:"column:skill_range" json:"skillRange"`

	// RecoveryScore 回收评分
	RecoveryScore int32 `gorm:"column:recovery_score" json:"recoveryScore"`

	// ExchangeBrand 兑换宠物需要的牌子数
	ExchangeBrand int32 `gorm:"column:exchange_brand" json:"exchangeBrand"`

	// LimitTime 限制交易天数
	LimitTime int32 `gorm:"column:limit_time" json:"limitTime"`
}

// TableName 显式指定表名
func (Pet) TableName() string {
	return "pet"
}
