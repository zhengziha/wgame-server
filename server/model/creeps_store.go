package model

import "time"

// CreepsStore 妖怪商店表。
type CreepsStore struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// Name 妖怪名称
	Name string `gorm:"size:255;column:name" json:"name"`

	// Price 价格
	Price int32 `gorm:"column:price" json:"price"`

	// AddTime 创建时间
	AddTime time.Time `gorm:"autoCreateTime;column:add_time" json:"addTime"`

	// UpdateTime 更新时间
	UpdateTime time.Time `gorm:"autoUpdateTime;column:update_time" json:"updateTime"`

	// Deleted 逻辑删除标记（tinyint(1)）
	Deleted bool `gorm:"column:deleted" json:"deleted"`
}

// TableName 显式指定表名
func (CreepsStore) TableName() string {
	return "creeps_store"
}
