package model

import "time"

// BlackList 黑名单表。
type BlackList struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// Data 黑名单数据
	Data string `gorm:"size:255;not null;column:data" json:"data"`

	// Type 类型（1:ip 2:机器码）
	Type int32 `gorm:"column:type" json:"type"`

	// AddTime 添加时间
	AddTime time.Time `gorm:"autoCreateTime;column:add_time" json:"addTime"`
}

// TableName 显式指定表名
func (BlackList) TableName() string {
	return "black_list"
}
