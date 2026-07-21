package model

import "time"

// CharaNickname 角色昵称池表。
// 注意主键 id 在 SQL 中是 bigint(20)，对应 Go int64。
type CharaNickname struct {
	// ID 主键（bigint）
	ID int64 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// Name 昵称
	Name string `gorm:"size:15;column:name" json:"name"`

	// Sex 性别
	Sex string `gorm:"size:4;column:sex" json:"sex"`

	// Label 标签
	Label string `gorm:"size:20;column:label" json:"label"`

	// Delete 是否删除（0:否）
	Delete int32 `gorm:"column:delete" json:"delete"`

	// CreateTime 创建时间
	CreateTime time.Time `gorm:"autoCreateTime;column:create_time" json:"createTime"`
}

// TableName 显式指定表名
func (CharaNickname) TableName() string {
	return "chara_nickname"
}
