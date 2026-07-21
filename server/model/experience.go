package model

import "time"

// Experience 升级经验配置表。
type Experience struct {
	// Attrib 等级（SQL 中为 PRIMARY KEY）
	Attrib int32 `gorm:"primaryKey;autoIncrement;column:attrib" json:"attrib"`

	// MaxLevel 升级经验
	MaxLevel int32 `gorm:"column:max_level" json:"maxLevel"`

	// AddTime 创建时间
	AddTime time.Time `gorm:"autoCreateTime;column:add_time" json:"addTime"`

	// UpdateTime 更新时间
	UpdateTime time.Time `gorm:"autoUpdateTime;column:update_time" json:"updateTime"`

	// Deleted 逻辑删除标记（tinyint(1)）
	Deleted bool `gorm:"column:deleted" json:"deleted"`
}

// TableName 显式指定表名
func (Experience) TableName() string {
	return "experience"
}
