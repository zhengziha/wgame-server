package model

import "time"

// PartyWarBid 帮派战报名竞价表。
// Java 实体字段使用 snake_case 命名（party_name/add_time），Go 字段转为驼峰。
type PartyWarBid struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// No 第几届
	No int32 `gorm:"column:no" json:"no"`

	// PartyName 帮派名字
	PartyName string `gorm:"size:255;column:party_name" json:"partyName"`

	// Cash 报名费用
	Cash int32 `gorm:"column:cash" json:"cash"`

	// AddTime 报名时间
	AddTime time.Time `gorm:"autoCreateTime;column:add_time" json:"addTime"`

	// State 状态 -1报名失败,0报名中,1报名成功
	State int32 `gorm:"column:state" json:"state"`
}

// TableName 显式指定表名
func (PartyWarBid) TableName() string {
	return "party_war_bid"
}
