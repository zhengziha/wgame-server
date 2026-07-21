package model

import "time"

// LuckDrawTimesReward 抽奖次数奖励表。
// times 字段在 SQL 中为 varchar(128)（存储次数范围），Java 中声明为 Integer，
// 此处按 Java 类型规则映射为 int32。
type LuckDrawTimesReward struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// Times 抽奖次数
	Times int32 `gorm:"column:times" json:"times"`

	// Title 奖励称号
	Title string `gorm:"size:128;column:title" json:"title"`

	// AddTime 创建时间
	AddTime time.Time `gorm:"autoCreateTime;column:add_time" json:"addTime"`
}

// TableName 显式指定表名
func (LuckDrawTimesReward) TableName() string {
	return "luck_draw_times_reward"
}
