package model

import "time"

// LuckDrawItem 抽奖物品表。
// level 为 SQL 保留字，column tag 直接写列名，GORM 会自动加反引号。
type LuckDrawItem struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// Item 奖品
	Item string `gorm:"size:255;column:item" json:"item"`

	// Level 等级（0:特等奖,1:一等奖,2:二等奖...4:四等奖）
	Level int32 `gorm:"column:level" json:"level"`

	// AddTime 创建时间
	AddTime time.Time `gorm:"autoCreateTime;column:add_time" json:"addTime"`
}

// TableName 显式指定表名
func (LuckDrawItem) TableName() string {
	return "luck_draw_item"
}
