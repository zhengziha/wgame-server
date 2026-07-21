package model

// GrowthConfig 成长配置表。
// level 为 SQL 保留字，column tag 直接写列名，GORM 会自动加反引号。
type GrowthConfig struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// Level 等级（列名 level 为 SQL 保留字）
	Level int32 `gorm:"column:level" json:"level"`

	// Exp 升级所需经验
	Exp int32 `gorm:"column:exp" json:"exp"`

	// BasicsAward 基础奖励
	BasicsAward string `gorm:"size:255;column:basics_award" json:"basicsAward"`

	// AdvancedAward 进阶奖励
	AdvancedAward string `gorm:"size:255;column:advanced_award" json:"advancedAward"`

	// SupersAward 超级奖励
	SupersAward string `gorm:"size:255;column:supers_award" json:"supersAward"`
}

// TableName 显式指定表名
func (GrowthConfig) TableName() string {
	return "growth_config"
}
