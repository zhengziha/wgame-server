package model

import "time"

// LuckyBagRandomBoxItem 福袋随机盒子物品表。
// level 为 SQL 保留字，column tag 直接写列名，GORM 会自动加反引号。
// probability 在 SQL 中为 double(11,4)，Java 中声明为 Integer，
// 此处按 Java 类型规则映射为 int32。
// 注：此表仅在 wd-game-preview.sql 中出现。
type LuckyBagRandomBoxItem struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// Item 奖品
	Item string `gorm:"size:255;column:item" json:"item"`

	// Level 等级（0:特等奖,1:一等奖,2:二等奖...4:四等奖）
	Level int32 `gorm:"column:level" json:"level"`

	// Probability 概率
	Probability int32 `gorm:"column:probability" json:"probability"`

	// AddTime 创建时间
	AddTime time.Time `gorm:"autoCreateTime;column:add_time" json:"addTime"`
}

// TableName 显式指定表名
func (LuckyBagRandomBoxItem) TableName() string {
	return "lucky_bag_random_box_item"
}
