package model

import "time"

// LuckyDrawRecord 融丹（抽奖）记录表。
type LuckyDrawRecord struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// Gid 角色 uuid
	Gid string `gorm:"size:32;column:gid" json:"gid"`

	// CharaName 角色名
	CharaName string `gorm:"size:32;column:chara_name" json:"charaName"`

	// Count 融丹次数
	Count int32 `gorm:"column:count" json:"count"`

	// CreateTime 创建时间
	CreateTime time.Time `gorm:"autoCreateTime;column:create_time" json:"createTime"`

	// UpdateTime 更新时间
	UpdateTime time.Time `gorm:"autoUpdateTime;column:update_time" json:"updateTime"`
}

// TableName 显式指定表名
func (LuckyDrawRecord) TableName() string {
	return "lucky_draw_record"
}
