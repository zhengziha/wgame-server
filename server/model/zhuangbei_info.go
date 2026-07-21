package model

import "time"

// ZhuangbeiInfo 装备信息表。
type ZhuangbeiInfo struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// Attrib 等级/属性标记
	Attrib int32 `gorm:"column:attrib" json:"attrib"`

	// Amount 佩戴位置
	Amount int32 `gorm:"column:amount" json:"amount"`

	// Type 外观
	Type int32 `gorm:"column:type" json:"type"`

	// Str 装备名字
	Str string `gorm:"size:255;column:str" json:"str"`

	// Quality 颜色/品质
	Quality string `gorm:"size:255;column:quality" json:"quality"`

	// Master 佩戴限制
	Master int32 `gorm:"column:master" json:"master"`

	// Metal 金木水火土
	Metal int32 `gorm:"column:metal" json:"metal"`

	// Mana 伤害
	Mana int32 `gorm:"column:mana" json:"mana"`

	// Accurate 命中/伤害
	Accurate int32 `gorm:"column:accurate" json:"accurate"`

	// Def 气血
	Def int32 `gorm:"column:def" json:"def"`

	// Dex 法力
	Dex int32 `gorm:"column:dex" json:"dex"`

	// Wiz 防御
	Wiz int32 `gorm:"column:wiz" json:"wiz"`

	// Parry 速度
	Parry int32 `gorm:"column:parry" json:"parry"`

	// AddTime 创建时间
	AddTime time.Time `gorm:"autoCreateTime;column:add_time" json:"addTime"`

	// UpdateTime 更新时间
	UpdateTime time.Time `gorm:"autoUpdateTime;column:update_time" json:"updateTime"`

	// Deleted 逻辑删除标记（tinyint(1)）
	Deleted bool `gorm:"column:deleted" json:"deleted"`
}

// TableName 显式指定表名
func (ZhuangbeiInfo) TableName() string {
	return "zhuangbei_info"
}
