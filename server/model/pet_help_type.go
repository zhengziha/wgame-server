package model

import "time"

// PetHelpType 宠物帮助类型表。
type PetHelpType struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// Type 类型
	Type int32 `gorm:"column:type" json:"type"`

	// Name 名称
	Name string `gorm:"size:32;not null;column:name" json:"name"`

	// AddTime 创建时间
	AddTime time.Time `gorm:"autoCreateTime;column:add_time" json:"addTime"`

	// UpdateTime 更新时间
	UpdateTime time.Time `gorm:"autoUpdateTime;column:update_time" json:"updateTime"`

	// Deleted 逻辑删除标记（tinyint(1)）
	Deleted bool `gorm:"column:deleted" json:"deleted"`

	// Quality 品质
	Quality int32 `gorm:"column:quality" json:"quality"`

	// Money 金额
	Money int32 `gorm:"column:money" json:"money"`

	// Polar 相性
	Polar int32 `gorm:"column:polar" json:"polar"`
}

// TableName 显式指定表名
func (PetHelpType) TableName() string {
	return "pet_help_type"
}
