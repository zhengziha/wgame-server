package model

import "time"

// ChargeConfig 累计充值配置表。
// 注意：menu 在 Java 实体中声明但未出现在
// wd-game-18.sql 建表语句中，按普通列保留以兼容 Java 逻辑。
type ChargeConfig struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// Money 充值金额（唯一）
	Money int32 `gorm:"uniqueIndex;column:money" json:"money"`

	// Menu 菜单（Java 实体声明，SQL 未建列）
	Menu string `gorm:"size:255;column:menu" json:"menu"`

	// Config 充值配置
	Config string `gorm:"type:text;column:config" json:"config"`

	// CreateTime 创建时间
	CreateTime time.Time `gorm:"autoCreateTime;column:create_time" json:"createTime"`
}

// TableName 显式指定表名
func (ChargeConfig) TableName() string {
	return "charge_config"
}
