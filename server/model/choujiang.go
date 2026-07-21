package model

import "time"

// Choujiang 抽奖表。
// Desc/Level 列名为 SQL 保留字。
type Choujiang struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// No 编号
	No int32 `gorm:"column:no" json:"no"`

	// Name 名称
	Name string `gorm:"size:255;column:name" json:"name"`

	// Desc 描述（列名 desc 为 SQL 保留字）
	Desc string `gorm:"size:255;column:desc" json:"desc"`

	// Level 等级（列名 level 为 SQL 保留字）
	Level int32 `gorm:"column:level" json:"level"`

	// AddTime 创建时间
	AddTime time.Time `gorm:"autoCreateTime;column:add_time" json:"addTime"`

	// UpdateTime 更新时间
	UpdateTime time.Time `gorm:"autoUpdateTime;column:update_time" json:"updateTime"`

	// Deleted 逻辑删除标记（tinyint(1)）
	Deleted bool `gorm:"column:deleted" json:"deleted"`
}

// TableName 显式指定表名
func (Choujiang) TableName() string {
	return "choujiang"
}
