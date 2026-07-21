package model

import "time"

// RenwuMonster 任务怪物表。
type RenwuMonster struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// MapName 地图名称
	MapName string `gorm:"size:64;column:map_name" json:"mapName"`

	// X 横坐标
	X int32 `gorm:"column:x" json:"x"`

	// Y 纵坐标
	Y int32 `gorm:"column:y" json:"y"`

	// Name 怪物名字
	Name string `gorm:"size:255;column:name" json:"name"`

	// Icon 怪物外观
	Icon int32 `gorm:"column:icon" json:"icon"`

	// Skills 怪物技能
	Skills string `gorm:"size:255;column:skills" json:"skills"`

	// Type 任务类型（除暴1/低级刷道2 等）
	Type int32 `gorm:"column:type" json:"type"`

	// AddTime 创建时间
	AddTime time.Time `gorm:"autoCreateTime;column:add_time" json:"addTime"`

	// UpdateTime 更新时间
	UpdateTime time.Time `gorm:"autoUpdateTime;column:update_time" json:"updateTime"`

	// Deleted 逻辑删除标记（tinyint(1)）
	Deleted bool `gorm:"column:deleted" json:"deleted"`
}

// TableName 显式指定表名
func (RenwuMonster) TableName() string {
	return "renwu_monster"
}
