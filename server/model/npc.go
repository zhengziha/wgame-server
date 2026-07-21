package model

import "time"

// Npc NPC 信息表。
// 注意：Type/SubType/MapName 为非持久化字段，数据库中不存在对应列。
type Npc struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// Icon 图标编号
	Icon int32 `gorm:"column:icon" json:"icon"`

	// Type NPC 类型（非持久化字段，数据库无此列）
	Type int32 `gorm:"-" json:"type"`

	// SubType NPC 子类型（非持久化字段，数据库无此列）
	SubType int32 `gorm:"-" json:"subType"`

	// X 横坐标
	X int32 `gorm:"column:x" json:"x"`

	// Y 纵坐标
	Y int32 `gorm:"column:y" json:"y"`

	// Name NPC 名称
	Name string `gorm:"size:255;column:name" json:"name"`

	// ShowName 显示名称（数据库列为驼峰 showName）
	ShowName string `gorm:"size:255;column:showName" json:"showName"`

	// MapID 所属地图编号
	MapID int32 `gorm:"column:map_id" json:"mapId"`

	// MapName 地图名称（非持久化字段，数据库无此列）
	MapName string `gorm:"-" json:"mapName"`

	// Ext 扩展字段（方向等信息）
	Ext string `gorm:"size:255;column:ext" json:"ext"`

	// AddTime 创建时间
	AddTime time.Time `gorm:"autoCreateTime;column:add_time" json:"addTime"`

	// UpdateTime 更新时间
	UpdateTime time.Time `gorm:"autoUpdateTime;column:update_time" json:"updateTime"`

	// Deleted 逻辑删除标记（tinyint(1)）
	Deleted bool `gorm:"column:deleted" json:"deleted"`
}

// TableName 显式指定表名
func (Npc) TableName() string {
	return "npc"
}
