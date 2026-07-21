package model

import "time"

// PartyMember 帮派成员表。
// 注意：wd-game-18.sql 中存在 chara_id 列，但 Java 实体未声明该字段，按 Java 实体字段为准。
type PartyMember struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// CharaGid 角色 uuid
	CharaGid string `gorm:"size:100;column:chara_gid" json:"charaGid"`

	// PartyId 帮派 id
	PartyId string `gorm:"size:128;column:party_id" json:"partyId"`

	// Name 名字
	Name string `gorm:"size:50;column:name" json:"name"`

	// Polar 相性
	Polar int32 `gorm:"column:polar" json:"polar"`

	// Job 职位
	Job string `gorm:"size:128;column:job" json:"job"`

	// LastWeekActive 上周活力
	LastWeekActive int32 `gorm:"column:last_week_active" json:"lastWeekActive"`

	// CurrWeekActive 本周活力
	CurrWeekActive int32 `gorm:"column:curr_week_active" json:"currWeekActive"`

	// Active 活力
	Active int32 `gorm:"column:active" json:"active"`

	// LogoutTime 离线时间
	LogoutTime time.Time `gorm:"column:logout_time" json:"logoutTime"`

	// CreateTime 入帮时间
	CreateTime time.Time `gorm:"autoCreateTime;column:create_time" json:"createTime"`

	// Info 拓展字段 JSON
	Info string `gorm:"type:text;column:info" json:"info"`
}

// TableName 显式指定表名
func (PartyMember) TableName() string {
	return "party_member"
}
