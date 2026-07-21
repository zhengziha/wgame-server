package model

import "time"

// GiftPackConfig 礼包配置表。
// status 为 SQL 保留字，column tag 直接写列名，GORM 会自动加反引号。
type GiftPackConfig struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// Status 活动状态
	Status string `gorm:"size:6;column:status" json:"status"`

	// Code 礼包码
	Code string `gorm:"size:64;column:code" json:"code"`

	// AwardGold 奖励金元宝
	AwardGold int32 `gorm:"column:award_gold" json:"awardGold"`

	// AwardSilver 奖励银元宝
	AwardSilver int32 `gorm:"column:award_silver" json:"awardSilver"`

	// AwardPoint 奖励积分
	AwardPoint int32 `gorm:"column:award_point" json:"awardPoint"`

	// AwardPet 奖励宠物
	AwardPet string `gorm:"size:255;column:award_pet" json:"awardPet"`

	// AwardGoods 奖励道具
	AwardGoods string `gorm:"size:255;column:award_goods" json:"awardGoods"`

	// GoodsLimitTimes 道具奖励限制次数
	GoodsLimitTimes string `gorm:"size:255;column:goods_limit_times" json:"goodsLimitTimes"`

	// AwardTitle 奖励称号
	AwardTitle string `gorm:"size:255;column:award_title" json:"awardTitle"`

	// ShuaDaoRatio 刷道效率（%）
	ShuaDaoRatio int32 `gorm:"column:shua_dao_ratio" json:"shuaDaoRatio"`

	// AwardCharge 奖励额外充值
	AwardCharge int32 `gorm:"column:award_charge" json:"awardCharge"`

	// AwardChargeStatus 奖励充值是否计入累充（1:开启）
	AwardChargeStatus int32 `gorm:"column:award_charge_status" json:"awardChargeStatus"`

	// AwardChargeRankStatus 奖励充值是否计入排行（1:开启）
	AwardChargeRankStatus int32 `gorm:"default:1;column:award_charge_rank_status" json:"awardChargeRankStatus"`

	// AddTime 创建时间
	AddTime time.Time `gorm:"autoCreateTime;column:add_time" json:"addTime"`
}

// TableName 显式指定表名
func (GiftPackConfig) TableName() string {
	return "gift_pack_config"
}
