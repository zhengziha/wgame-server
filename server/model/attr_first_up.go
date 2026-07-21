package model

import "time"

// AttrFirstUp 属性首升活动表。
type AttrFirstUp struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// Attr 属性值
	Attr int32 `gorm:"column:attr" json:"attr"`

	// AttrName 属性名称
	AttrName string `gorm:"size:128;column:attr_name" json:"attrName"`

	// Title 奖励称号
	Title string `gorm:"size:128;column:title" json:"title"`

	// Uuid 获得的角色 gid
	Uuid string `gorm:"size:128;column:uuid" json:"uuid"`

	// Name 获得的角色名称
	Name string `gorm:"size:128;column:name" json:"name"`

	// AddTime 创建时间
	AddTime time.Time `gorm:"autoCreateTime;column:add_time" json:"addTime"`

	// UpdateTime 更新时间
	UpdateTime time.Time `gorm:"autoUpdateTime;column:update_time" json:"updateTime"`
}

// TableName 显式指定表名
func (AttrFirstUp) TableName() string {
	return "attr_first_up"
}
