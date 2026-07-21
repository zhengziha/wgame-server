package model

import "time"

// WeddingList 婚礼礼单表。
type WeddingList struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// Name 名字
	Name string `gorm:"size:255;column:name" json:"name"`

	// Price 价格
	Price int32 `gorm:"column:price" json:"price"`

	// Type 类型（花车/音乐/祝词）
	Type string `gorm:"size:255;column:type" json:"type"`

	// Icon 外观
	Icon int32 `gorm:"column:icon" json:"icon"`

	// AddTime 时间
	AddTime time.Time `gorm:"autoCreateTime;column:add_time" json:"addTime"`

	// PlayTime 播放时间
	PlayTime int32 `gorm:"column:play_time" json:"playTime"`
}

// TableName 显式指定表名
func (WeddingList) TableName() string {
	return "wedding_list"
}
