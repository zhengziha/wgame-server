package model

import "time"

// FightObjectInfo 战斗对象信息表。
// Fabao 字段在 Java 实体中声明但未出现在 wd-game-18.sql 建表语句中，按普通列保留。
type FightObjectInfo struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// Type 类型
	Type string `gorm:"size:255;column:type" json:"type"`

	// Name 名字
	Name string `gorm:"size:255;column:name" json:"name"`

	// Level 等级
	Level int32 `gorm:"column:level" json:"level"`

	// ShowName 在游戏显示的名字
	ShowName string `gorm:"size:255;column:show_name" json:"showName"`

	// Life 生命值
	Life int32 `gorm:"column:life" json:"life"`

	// Mana 法力值
	Mana int32 `gorm:"column:mana" json:"mana"`

	// PhyAttack 物攻值
	PhyAttack int32 `gorm:"column:phy_attack" json:"phyAttack"`

	// MagAttack 法功值
	MagAttack int32 `gorm:"column:mag_attack" json:"magAttack"`

	// Polar 金木水火土
	Polar string `gorm:"size:255;column:polar" json:"polar"`

	// Speed 速度值
	Speed int32 `gorm:"column:speed" json:"speed"`

	// Def 防御值
	Def int32 `gorm:"column:def" json:"def"`

	// Icon 外观
	Icon int32 `gorm:"column:icon" json:"icon"`

	// Daohang 道行
	Daohang int32 `gorm:"column:daohang" json:"daohang"`

	// PetMartial 宠物武学
	PetMartial int32 `gorm:"column:pet_martial" json:"petMartial"`

	// Skill 技能（逗号分隔，多系用 # 分隔）
	Skill string `gorm:"type:text;column:skill" json:"skill"`

	// PetTianshu 宠物天书技能（逗号分隔）
	PetTianshu string `gorm:"size:255;column:pet_tianshu" json:"petTianshu"`

	// Fabao 法宝
	Fabao string `gorm:"size:255;column:fabao" json:"fabao"`

	// AllResistPolar 抗所有相性
	AllResistPolar int32 `gorm:"column:all_resist_polar" json:"allResistPolar"`

	// ResistMetal 抗金
	ResistMetal int32 `gorm:"column:resist_metal" json:"resistMetal"`

	// ResistWood 抗木
	ResistWood int32 `gorm:"column:resist_wood" json:"resistWood"`

	// ResistWater 抗水
	ResistWater int32 `gorm:"column:resist_water" json:"resistWater"`

	// ResistFire 抗火
	ResistFire int32 `gorm:"column:resist_fire" json:"resistFire"`

	// ResistEarth 抗土
	ResistEarth int32 `gorm:"column:resist_earth" json:"resistEarth"`

	// DoubleHitRate 连击率
	DoubleHitRate int32 `gorm:"column:double_hit_rate" json:"doubleHitRate"`

	// DoubleHit 连击数
	DoubleHit int32 `gorm:"column:double_hit" json:"doubleHit"`

	// MstuntRate 法术必杀
	MstuntRate int32 `gorm:"column:mstunt_rate" json:"mstuntRate"`

	// ResistForgotten 抗毒
	ResistForgotten int32 `gorm:"column:resist_forgotten" json:"resistForgotten"`

	// ResistPoison 抵抗冻结
	ResistPoison int32 `gorm:"column:resist_poison" json:"resistPoison"`

	// ResistFrozen 抵抗睡眠
	ResistFrozen int32 `gorm:"column:resist_frozen" json:"resistFrozen"`

	// ResistSleep 抵抗被遗忘
	ResistSleep int32 `gorm:"column:resist_sleep" json:"resistSleep"`

	// ResistConfusion 抵制混乱
	ResistConfusion int32 `gorm:"column:resist_confusion" json:"resistConfusion"`

	// AllResistExcept 所有抗异常
	AllResistExcept int32 `gorm:"column:all_resist_except" json:"allResistExcept"`

	// AddTime 创建时间
	AddTime time.Time `gorm:"autoCreateTime;column:add_time" json:"addTime"`

	// UpdateTime 更新时间
	UpdateTime time.Time `gorm:"autoUpdateTime;column:update_time" json:"updateTime"`

	// Deleted 逻辑删除标记（tinyint(1)）
	Deleted bool `gorm:"column:deleted" json:"deleted"`
}

// TableName 显式指定表名
func (FightObjectInfo) TableName() string {
	return "fight_object_info"
}
