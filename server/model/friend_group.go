package model

import "time"

// FriendGroup 好友分组表。
type FriendGroup struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// Name 分组名称
	Name string `gorm:"size:255;column:name" json:"name"`

	// GroupId 分组 id
	GroupId string `gorm:"size:255;column:group_id" json:"groupId"`

	// Gid 用户 id
	Gid string `gorm:"size:255;column:gid" json:"gid"`

	// UpdateTime 更新时间
	UpdateTime time.Time `gorm:"autoUpdateTime;column:update_time" json:"updateTime"`

	// AddTime 添加时间
	AddTime time.Time `gorm:"autoCreateTime;column:add_time" json:"addTime"`
}

// TableName 显式指定表名
func (FriendGroup) TableName() string {
	return "friend_group"
}
