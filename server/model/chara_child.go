package model

import "time"

// CharaChild 宠物娃娃信息表。
// 对应建表语句位于 wd-game-test.sql。
type CharaChild struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// Uuid 用户 uuid
	Uuid string `gorm:"size:255;column:uuid" json:"uuid"`

	// Cid 用户 id
	Cid int32 `gorm:"column:cid" json:"cid"`

	// JsonData 娃娃 JSON
	JsonData string `gorm:"type:mediumtext;column:json_data" json:"jsonData"`

	// ChildName 娃娃名称
	ChildName string `gorm:"size:50;column:child_name" json:"childName"`

	// OwnerName 所属角色名
	OwnerName string `gorm:"size:50;column:owner_name" json:"ownerName"`

	// AddTime 创建时间
	AddTime time.Time `gorm:"autoCreateTime;column:add_time" json:"addTime"`

	// UpdateTime 更新时间
	UpdateTime time.Time `gorm:"autoUpdateTime;column:update_time" json:"updateTime"`
}

// TableName 显式指定表名
func (CharaChild) TableName() string {
	return "chara_child"
}
