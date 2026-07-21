package model

import "time"

// ExchangeConfig 兑换配置表。
// menu 列在 SQL 中建有唯一索引（UNIQUE INDEX money(menu)）。
type ExchangeConfig struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// Npc 指定 NPC
	Npc string `gorm:"size:255;not null;column:npc" json:"npc"`

	// Menu 菜单内容（唯一索引）
	Menu string `gorm:"size:255;not null;uniqueIndex;column:menu" json:"menu"`

	// CostStr 消耗目标
	CostStr string `gorm:"type:text;not null;column:cost_str" json:"costStr"`

	// TargetStr 兑换目标
	TargetStr string `gorm:"type:text;not null;column:target_str" json:"targetStr"`

	// LimitNum 限制次数
	LimitNum int32 `gorm:"column:limit_num" json:"limitNum"`

	// CreateTime 创建时间
	CreateTime time.Time `gorm:"autoCreateTime;column:create_time" json:"createTime"`

	// Deleted 逻辑删除标记（tinyint）
	Deleted bool `gorm:"column:deleted" json:"deleted"`
}

// TableName 显式指定表名
func (ExchangeConfig) TableName() string {
	return "exchange_config"
}
