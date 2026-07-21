package model

import "time"

// Reports 充值报表表。
type Reports struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// Zhanghao 账号
	Zhanghao string `gorm:"size:111;not null;column:zhanghao" json:"zhanghao"`

	// Yuanbaoshu 元宝数
	Yuanbaoshu int32 `gorm:"not null;column:yuanbaoshu" json:"yuanbaoshu"`

	// Shifouchongzhi 是否充值
	Shifouchongzhi string `gorm:"size:4;not null;column:shifouchongzhi" json:"shifouchongzhi"`

	// AddTime 创建时间
	AddTime time.Time `gorm:"autoCreateTime;column:add_time" json:"addTime"`

	// UpdateTime 更新时间
	UpdateTime time.Time `gorm:"autoUpdateTime;column:update_time" json:"updateTime"`

	// Deleted 逻辑删除标记（tinyint(1)）
	Deleted bool `gorm:"column:deleted" json:"deleted"`
}

// TableName 显式指定表名
func (Reports) TableName() string {
	return "reports"
}
