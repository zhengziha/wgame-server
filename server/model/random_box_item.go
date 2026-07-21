package model

import "time"

// RandomBoxItem 随机盒子物品表。
// level 为 SQL 保留字，column tag 直接写列名，GORM 会自动加反引号。
// probability 字段在 Java 中是 Integer，但 SQL 中 random_box_item 表未出现该列；
// 此处按 Java 实体保留为普通列以兼容业务逻辑。
type RandomBoxItem struct {
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
func (RandomBoxItem) TableName() string {
	return "random_box_item"
}
