package model

import "time"

// CharaPet 宠物信息表。
// petStatus 在 Java 实体中声明，wd-game-18.sql 建表语句未出现该列，按普通列保留以兼容 Java 逻辑。
type CharaPet struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// Uuid 用户 uuid
	Uuid string `gorm:"size:255;column:uuid" json:"uuid"`

	// Cid 用户 id
	Cid int32 `gorm:"column:cid" json:"cid"`

	// Pet 宠物 JSON
	Pet string `gorm:"type:mediumtext;column:pet" json:"pet"`

	// PetName 宠物名称
	PetName string `gorm:"size:50;column:pet_name" json:"petName"`

	// PetStatus 宠物状态
	PetStatus int32 `gorm:"column:pet_status" json:"petStatus"`

	// OwnerName 所属角色名
	OwnerName string `gorm:"size:50;column:owner_name" json:"ownerName"`

	// AddTime 创建时间
	AddTime time.Time `gorm:"autoCreateTime;column:add_time" json:"addTime"`

	// UpdateTime 更新时间
	UpdateTime time.Time `gorm:"autoUpdateTime;column:update_time" json:"updateTime"`
}

// TableName 显式指定表名
func (CharaPet) TableName() string {
	return "chara_pet"
}
