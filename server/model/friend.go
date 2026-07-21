package model

import "time"

// Friend 好友表。
type Friend struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// FriendScore 好友度
	FriendScore int32 `gorm:"column:friend_score" json:"friendScore"`

	// FriendName 好友名字
	FriendName string `gorm:"size:255;column:friend_name" json:"friendName"`

	// FriendGid 好友 gid
	FriendGid string `gorm:"size:255;column:friend_gid" json:"friendGid"`

	// Gid 角色表 id
	Gid string `gorm:"size:255;column:gid" json:"gid"`

	// GroupId 分组 id
	GroupId string `gorm:"size:255;column:group_id" json:"groupId"`

	// GroupName 分组名称
	GroupName string `gorm:"size:255;column:group_name" json:"groupName"`

	// UpdateTime 更新时间
	UpdateTime time.Time `gorm:"autoUpdateTime;column:update_time" json:"updateTime"`

	// AddTime 添加时间
	AddTime time.Time `gorm:"autoCreateTime;column:add_time" json:"addTime"`

	// ExtJson 拓展字段 JSON
	ExtJson string `gorm:"type:longtext;column:ext_json" json:"extJson"`
}

// TableName 显式指定表名
func (Friend) TableName() string {
	return "friend"
}
