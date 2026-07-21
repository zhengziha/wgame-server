package model

import "time"

// MysticalRecharge 神秘充值任务配置表。
type MysticalRecharge struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// TriggerMoney 触发任务金额
	TriggerMoney int32 `gorm:"column:trigger_money" json:"triggerMoney"`

	// ExtAwardMoney 任务完成奖励金额（金额/倍数）
	ExtAwardMoney int32 `gorm:"column:ext_award_money" json:"extAwardMoney"`

	// AwardMoneys 随机金额列表
	AwardMoneys string `gorm:"size:255;column:award_moneys" json:"awardMoneys"`

	// AwardType 任务奖励类型（0:按金额,1:按随机2,3倍数,2:按金额随机）
	AwardType int32 `gorm:"default:0;column:award_type" json:"awardType"`

	// AddTime 创建时间
	AddTime time.Time `gorm:"autoCreateTime;column:add_time" json:"addTime"`
}

// TableName 显式指定表名
func (MysticalRecharge) TableName() string {
	return "mystical_recharge"
}
