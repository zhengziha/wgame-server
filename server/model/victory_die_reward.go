package model

import "time"

// VictoryDieReward 胜利死亡奖励配置表。
// 注意：expRate、taoRate 在 Java 实体中为 Boolean 类型声明，
// 但未出现在 wd-game-18.sql 建表语句中，按普通列保留以兼容 Java 逻辑。
type VictoryDieReward struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// Name 名称
	Name string `gorm:"size:255;column:name" json:"name"`

	// Score 充值积分
	Score int32 `gorm:"column:score" json:"score"`

	// ExpRate 经验是否按百分比（Java 实体声明，SQL 未建列）
	ExpRate bool `gorm:"column:exp_rate" json:"expRate"`

	// Exp 经验
	Exp int32 `gorm:"column:exp" json:"exp"`

	// Tao 道行
	Tao int32 `gorm:"column:tao" json:"tao"`

	// TaoRate 道行是否衰减（Java 实体声明，SQL 未建列）
	TaoRate bool `gorm:"column:tao_rate" json:"taoRate"`

	// GoldCoin 金元宝
	GoldCoin int32 `gorm:"column:gold_coin" json:"goldCoin"`

	// SilverCoin 银元宝
	SilverCoin int32 `gorm:"column:silver_coin" json:"silverCoin"`

	// ShenHun 神魂
	ShenHun int32 `gorm:"column:shen_hun" json:"shenHun"`

	// Chenghao 称号
	Chenghao string `gorm:"size:255;column:chenghao" json:"chenghao"`

	// ChenghaoExpireTime 称号过期时间（分钟）
	ChenghaoExpireTime int32 `gorm:"column:chenghao_expire_time" json:"chenghaoExpireTime"`

	// Daoju 道具
	Daoju string `gorm:"size:255;column:daoju" json:"daoju"`

	// EquipAttr 装备
	EquipAttr string `gorm:"size:255;column:equip_attr" json:"equipAttr"`

	// Shoushi 首饰
	Shoushi string `gorm:"size:255;column:shoushi" json:"shoushi"`

	// Pet 宠物
	Pet string `gorm:"size:255;column:pet" json:"pet"`

	// Type 类型（0:胜利 1:死亡）
	Type int32 `gorm:"column:type" json:"type"`

	// AddTime 添加时间
	AddTime time.Time `gorm:"autoCreateTime;column:add_time" json:"addTime"`

	// Deleted 是否删除（0:未删除 1:已删除）
	Deleted int32 `gorm:"column:deleted" json:"deleted"`

	// Wuxue 武学
	Wuxue int32 `gorm:"column:wuxue" json:"wuxue"`
}

// TableName 显式指定表名
func (VictoryDieReward) TableName() string {
	return "victory_die_reward"
}
