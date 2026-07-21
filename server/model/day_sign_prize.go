package model

import "time"

// DaySignPrize 签到奖励表。
// day 字段在数据库中为 UNIQUE 索引。
type DaySignPrize struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// Name 奖励名称
	Name string `gorm:"size:38;not null;column:name" json:"name"`

	// Day 签到天数（唯一）
	Day int32 `gorm:"not null;column:day" json:"day"`

	// AwardStr 奖励格式串
	AwardStr string `gorm:"size:255;column:award_str" json:"awardStr"`

	// Type 类型
	Type string `gorm:"size:11;column:type" json:"type"`

	// Num 数值
	Num int32 `gorm:"default:1;column:num" json:"num"`

	// AddTime 创建时间
	AddTime time.Time `gorm:"autoCreateTime;column:add_time" json:"addTime"`

	// UpdateTime 更新时间
	UpdateTime time.Time `gorm:"autoUpdateTime;column:update_time" json:"updateTime"`

	// Deleted 逻辑删除标记（tinyint(1)）
	Deleted bool `gorm:"column:deleted" json:"deleted"`
}

// TableName 显式指定表名
func (DaySignPrize) TableName() string {
	return "day_sign_prize"
}
