package model

import "time"

// NpcPoint 地图出入点（传送点）表。
type NpcPoint struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// Mapname 进入的目标地图名称
	Mapname string `gorm:"size:128;column:mapname" json:"mapname"`

	// Doorname 角色当前所在的地图名称
	Doorname string `gorm:"size:128;column:doorname" json:"doorname"`

	// X 出入点显示横坐标
	X int32 `gorm:"column:x" json:"x"`

	// Y 出入点显示纵坐标
	Y int32 `gorm:"column:y" json:"y"`

	// Z 方向
	Z int32 `gorm:"column:z" json:"z"`

	// Inx 作为进入点时的横坐标
	Inx int32 `gorm:"column:inx" json:"inx"`

	// Iny 作为进入点时的纵坐标
	Iny int32 `gorm:"column:iny" json:"iny"`

	// AddTime 创建时间
	AddTime time.Time `gorm:"autoCreateTime;column:add_time" json:"addTime"`

	// UpdateTime 更新时间
	UpdateTime time.Time `gorm:"autoUpdateTime;column:update_time" json:"updateTime"`

	// Deleted 逻辑删除标记（tinyint(1)）
	Deleted bool `gorm:"column:deleted" json:"deleted"`
}

// TableName 显式指定表名
func (NpcPoint) TableName() string {
	return "npc_point"
}
