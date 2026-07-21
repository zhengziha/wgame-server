package model

import "time"

// HouseFurniture 房屋家具表。
type HouseFurniture struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// Name 家具名称
	Name string `gorm:"size:15;column:name" json:"name"`

	// MapIndex 地图索引
	MapIndex int32 `gorm:"not null;column:map_index" json:"mapIndex"`

	// MapName 地图名称
	MapName string `gorm:"size:15;column:map_name" json:"mapName"`

	// FloorIndex 地板样式
	FloorIndex int32 `gorm:"column:floor_index" json:"floorIndex"`

	// WallIndex 墙体样式
	WallIndex int32 `gorm:"column:wall_index" json:"wallIndex"`

	// Pos 位置
	Pos int32 `gorm:"column:pos" json:"pos"`

	// Fid 样式 id
	Fid int32 `gorm:"column:fid" json:"fid"`

	// Bsx 起始坐标 x
	Bsx int32 `gorm:"column:bsx" json:"bsx"`

	// Bsy 起始坐标 y
	Bsy int32 `gorm:"column:bsy" json:"bsy"`

	// X 坐标 x
	X int32 `gorm:"column:x" json:"x"`

	// Y 坐标 y
	Y int32 `gorm:"column:y" json:"y"`

	// Flip 翻转
	Flip int32 `gorm:"column:flip" json:"flip"`

	// Durability 耐久度
	Durability int32 `gorm:"column:durability" json:"durability"`

	// HouseId 居所 id
	HouseId string `gorm:"size:50;column:house_id" json:"houseId"`

	// Gid 玩家 gid
	Gid string `gorm:"size:50;column:gid" json:"gid"`

	// AddTime 添加时间
	AddTime time.Time `gorm:"autoCreateTime;column:add_time" json:"addTime"`
}

// TableName 显式指定表名
func (HouseFurniture) TableName() string {
	return "house_furniture"
}
