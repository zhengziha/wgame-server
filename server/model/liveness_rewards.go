package model

import "time"

// LivenessRewards 活跃度奖励领取记录表。
type LivenessRewards struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// Gid 用户 id
	Gid string `gorm:"size:255;column:gid" json:"gid"`

	// Name 角色名
	Name string `gorm:"size:255;column:name" json:"name"`

	// Activity 活跃度
	Activity int32 `gorm:"column:activity" json:"activity"`

	// AddTime 领取时间
	AddTime time.Time `gorm:"autoCreateTime;column:add_time" json:"addTime"`
}

// TableName 显式指定表名
func (LivenessRewards) TableName() string {
	return "liveness_rewards"
}
