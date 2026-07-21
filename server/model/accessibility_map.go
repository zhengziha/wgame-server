package model

import "time"

// AccessibilityMap 无障碍地图表。
// 注意：Delete 字段对应 SQL 列名 delete（SQL 保留字），
// column tag 直接写列名，GORM 会自动加反引号。
type AccessibilityMap struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// MapName 地图名字
	MapName string `gorm:"size:255;column:map_name" json:"mapName"`

	// X 坐标 x
	X int32 `gorm:"column:x" json:"x"`

	// Y 坐标 y
	Y int32 `gorm:"column:y" json:"y"`

	// CreateTime 创建时间
	CreateTime time.Time `gorm:"autoCreateTime;column:create_time" json:"createTime"`

	// Delete 是否删除（列名 delete 为 SQL 保留字）
	Delete int32 `gorm:"column:delete" json:"delete"`
}

// TableName 显式指定表名
func (AccessibilityMap) TableName() string {
	return "accessibility_map"
}
