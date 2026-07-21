package model

import "time"

// FirstKill 首杀活动表。
type FirstKill struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// BossType boss 类型
	BossType string `gorm:"size:128;column:boss_type" json:"bossType"`

	// BossName boss 名称
	BossName string `gorm:"size:128;column:boss_name" json:"bossName"`

	// Title 奖励称号
	Title string `gorm:"size:128;column:title" json:"title"`

	// Uuid 获得的角色 gid
	Uuid string `gorm:"size:128;column:uuid" json:"uuid"`

	// Name 获得的角色名称
	Name string `gorm:"size:128;column:name" json:"name"`

	// PlayersInfo 参与玩家信息 JSON
	PlayersInfo string `gorm:"type:text;column:players_info" json:"playersInfo"`

	// AddTime 创建时间
	AddTime time.Time `gorm:"autoCreateTime;column:add_time" json:"addTime"`

	// GetDate 获得日期
	GetDate time.Time `gorm:"column:get_date" json:"getDate"`

	// CostTime 战斗花费时间
	CostTime int32 `gorm:"column:cost_time" json:"costTime"`
}

// TableName 显式指定表名
func (FirstKill) TableName() string {
	return "first_kill"
}
