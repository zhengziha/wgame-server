package model

import "time"

// ChargeRank 充值排行表。
// 注意：该表未出现在 wd-game-18.sql，表名与列结构依据
// wd-game-test.sql 中的 charge_rank 建表语句确定。
type ChargeRank struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// RankNo 排名序号
	RankNo int32 `gorm:"column:rank_no" json:"rankNo"`

	// Money 充值金额
	Money string `gorm:"size:255;column:money" json:"money"`

	// AccountName 玩家账号
	AccountName string `gorm:"size:255;column:account_name" json:"accountName"`

	// State 领取状态
	State int32 `gorm:"column:state" json:"state"`

	// Gid 角色 id
	Gid string `gorm:"size:128;column:gid" json:"gid"`

	// CharaName 角色名
	CharaName string `gorm:"size:255;column:chara_name" json:"charaName"`

	// Rewards 奖励
	Rewards string `gorm:"type:text;column:rewards" json:"rewards"`

	// ReceiveTime 领取时间
	ReceiveTime time.Time `gorm:"column:receive_time" json:"receiveTime"`

	// CreateTime 创建时间
	CreateTime time.Time `gorm:"autoCreateTime;column:create_time" json:"createTime"`

	// UpdateTime 更新时间
	UpdateTime time.Time `gorm:"autoUpdateTime;column:update_time" json:"updateTime"`
}

// TableName 显式指定表名
func (ChargeRank) TableName() string {
	return "charge_rank"
}
