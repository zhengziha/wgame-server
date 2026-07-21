package model

import "time"

// ChargeGetRecord 累计充值领取记录表。
type ChargeGetRecord struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// Money 领取金额类型
	Money int32 `gorm:"column:money" json:"money"`

	// Name 玩家姓名
	Name string `gorm:"size:255;column:name" json:"name"`

	// Account 玩家 uuid
	Account string `gorm:"size:255;column:account" json:"account"`

	// CreateTime 领取时间
	CreateTime time.Time `gorm:"autoCreateTime;column:create_time" json:"createTime"`
}

// TableName 显式指定表名
func (ChargeGetRecord) TableName() string {
	return "charge_get_record"
}
