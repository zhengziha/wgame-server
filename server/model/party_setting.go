package model

// PartySetting 帮派活动设置表。
// 对应建表语句位于 wd-game-test.sql。
// partyId + activityType 构成唯一索引 uq_party_id_type。
type PartySetting struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// PartyId 帮派 id
	PartyId string `gorm:"size:64;not null;column:party_id" json:"partyId"`

	// ActivityType 活动类型
	ActivityType string `gorm:"size:32;not null;column:activity_type" json:"activityType"`

	// OpenTime 开放时间
	OpenTime string `gorm:"size:64;not null;column:open_time" json:"openTime"`

	// AutoType 自动类型
	AutoType int32 `gorm:"column:auto_type" json:"autoType"`
}

// TableName 显式指定表名
func (PartySetting) TableName() string {
	return "party_setting"
}
