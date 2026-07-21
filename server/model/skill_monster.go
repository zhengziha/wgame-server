package model

import "time"

// SkillMonster 技能怪物表。
type SkillMonster struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// Name 怪物名字
	Name string `gorm:"size:255;column:name" json:"name"`

	// Skills 怪物技能
	Skills string `gorm:"size:255;column:skills" json:"skills"`

	// Type 类型
	Type int32 `gorm:"column:type" json:"type"`

	// AddTime 创建时间
	AddTime time.Time `gorm:"autoCreateTime;column:add_time" json:"addTime"`

	// UpdateTime 更新时间
	UpdateTime time.Time `gorm:"autoUpdateTime;column:update_time" json:"updateTime"`

	// Deleted 逻辑删除标记（tinyint(1)）
	Deleted bool `gorm:"column:deleted" json:"deleted"`
}

// TableName 显式指定表名
func (SkillMonster) TableName() string {
	return "skill_monster"
}
