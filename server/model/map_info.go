package model

import "time"

// MapInfo 地图信息表（SQL 表名为保留字 map）。
type MapInfo struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// Name 地图名称
	Name string `gorm:"size:255;not null;column:name" json:"name"`

	// MapID 地图编号
	MapID int32 `gorm:"not null;column:map_id" json:"mapId"`

	// X 横坐标
	X int32 `gorm:"not null;column:x" json:"x"`

	// Y 纵坐标
	Y int32 `gorm:"not null;column:y" json:"y"`

	// Icon 图标
	Icon string `gorm:"size:255;column:icon" json:"icon"`

	// MonsterLevel 怪物等级
	MonsterLevel int32 `gorm:"column:monster_level" json:"monsterLevel"`

	// AddTime 创建时间
	AddTime time.Time `gorm:"autoCreateTime;column:add_time" json:"addTime"`

	// UpdateTime 更新时间
	UpdateTime time.Time `gorm:"autoUpdateTime;column:update_time" json:"updateTime"`

	// Deleted 逻辑删除标记（@LogicDelete，tinyint(1)）
	Deleted bool `gorm:"column:deleted" json:"deleted"`
}

// TableName 显式指定表名（map 为 SQL 保留字）
func (MapInfo) TableName() string {
	return "map"
}
