package model

import "time"

// LuckyDrawRank 抽奖排名表。
// state 为业务状态字段，按普通列保留。
type LuckyDrawRank struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// RankNo 排名
	RankNo int32 `gorm:"column:rank_no" json:"rankNo"`

	// Count 抽奖次数
	Count int32 `gorm:"column:count" json:"count"`

	// State 领取状态
	State int32 `gorm:"column:state" json:"state"`

	// Gid 角色 uuid
	Gid string `gorm:"size:36;column:gid" json:"gid"`

	// CharaName 角色名
	CharaName string `gorm:"size:36;column:chara_name" json:"charaName"`

	// Rewards 奖励内容 JSON
	Rewards string `gorm:"type:text;column:rewards" json:"rewards"`

	// ReceiveTime 领取时间
	ReceiveTime time.Time `gorm:"column:receive_time" json:"receiveTime"`

	// CreateTime 创建时间
	CreateTime time.Time `gorm:"autoCreateTime;column:create_time" json:"createTime"`

	// UpdateTime 更新时间
	UpdateTime time.Time `gorm:"autoUpdateTime;column:update_time" json:"updateTime"`
}

// TableName 显式指定表名
func (LuckyDrawRank) TableName() string {
	return "lucky_draw_rank"
}
