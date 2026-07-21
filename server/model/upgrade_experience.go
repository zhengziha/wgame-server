package model

import "time"

// UpgradeExperience 升级经验配置表。
// 注意：Java 实体未声明 @Table 也未标注 @Id，表名依据
// wd-game-test.sql 中的 upgrade_experience 建表语句确定。
// SQL 中 level 列为 AUTO_INCREMENT PRIMARY KEY。
type UpgradeExperience struct {
	// Level 等级（SQL 中为 AUTO_INCREMENT PRIMARY KEY）
	Level int32 `gorm:"primaryKey;autoIncrement;column:level" json:"level"`

	// Experience 升级经验
	Experience int32 `gorm:"column:experience" json:"experience"`

	// AddTime 创建时间
	AddTime time.Time `gorm:"autoCreateTime;column:add_time" json:"addTime"`

	// UpdateTime 更新时间
	UpdateTime time.Time `gorm:"autoUpdateTime;column:update_time" json:"updateTime"`

	// Deleted 逻辑删除标记（tinyint(1)）
	Deleted bool `gorm:"column:deleted" json:"deleted"`
}

// TableName 显式指定表名
func (UpgradeExperience) TableName() string {
	return "upgrade_experience"
}
