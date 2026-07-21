package model

import "time"

// OnceCharge 单次充值配置表。
// status 为 SQL 保留字，column tag 直接写列名，GORM 会自动加反引号。
// awardLotteryDraw/goodsLimitTimes/shuaDaoRatio/awardCharge/awardChargeRandom/
// awardChargeStatus/onceGiveStatus/limitAwardCharge/limitOnce 在 Java 实体中声明，
// wd-game-18.sql 建表语句未出现这些列，按普通列保留以兼容 Java 逻辑。
type OnceCharge struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// Status 活动状态（列名 status 为 SQL 保留字）
	Status string `gorm:"size:6;column:status" json:"status"`

	// ChargeMoney 充值金额
	ChargeMoney int32 `gorm:"column:charge_money" json:"chargeMoney"`

	// AwardGold 奖励金元宝
	AwardGold int32 `gorm:"column:award_gold" json:"awardGold"`

	// AwardSilver 奖励银元宝
	AwardSilver int32 `gorm:"column:award_silver" json:"awardSilver"`

	// AwardPoint 奖励积分
	AwardPoint int32 `gorm:"column:award_point" json:"awardPoint"`

	// AwardLotteryDraw 融丹次数
	AwardLotteryDraw int32 `gorm:"column:award_lottery_draw" json:"awardLotteryDraw"`

	// AwardPet 奖励宠物
	AwardPet string `gorm:"size:255;column:award_pet" json:"awardPet"`

	// AwardGoods 奖励道具
	AwardGoods string `gorm:"size:255;column:award_goods" json:"awardGoods"`

	// GoodsLimitTimes 道具奖励限制次数
	GoodsLimitTimes string `gorm:"column:goods_limit_times" json:"goodsLimitTimes"`

	// AwardTitle 奖励称号
	AwardTitle string `gorm:"size:255;column:award_title" json:"awardTitle"`

	// ShuaDaoRatio 刷道效率
	ShuaDaoRatio int32 `gorm:"column:shua_dao_ratio" json:"shuaDaoRatio"`

	// AwardCharge 奖励额外充值
	AwardCharge string `gorm:"column:award_charge" json:"awardCharge"`

	// AwardChargeRandom 奖励额外充值随机
	AwardChargeRandom int32 `gorm:"column:award_charge_random" json:"awardChargeRandom"`

	// AwardChargeStatus 奖励充值是否计入累充 1开启
	AwardChargeStatus int32 `gorm:"column:award_charge_status" json:"awardChargeStatus"`

	// OnceGiveStatus 奖励充值是否可以领取单笔 1开启
	OnceGiveStatus int32 `gorm:"column:once_give_status" json:"onceGiveStatus"`

	// LimitAwardCharge 奖励每天次数限制
	LimitAwardCharge int32 `gorm:"column:limit_award_charge" json:"limitAwardCharge"`

	// LimitOnce 只能领取一次的单笔
	LimitOnce int32 `gorm:"column:limit_once" json:"limitOnce"`

	// AddTime 创建时间
	AddTime time.Time `gorm:"autoCreateTime;column:add_time" json:"addTime"`
}

// TableName 显式指定表名
func (OnceCharge) TableName() string {
	return "one_charge"
}
