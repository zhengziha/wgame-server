package model

import "time"

// Charge 充值记录表。
// 注意：orderNo、onceGiveStatus 在 Java 实体中声明但未出现在
// wd-game-18.sql 建表语句中，按普通列保留。
type Charge struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// Accountname 玩家账号
	Accountname string `gorm:"size:255;not null;column:accountname" json:"accountname"`

	// Coin 充值金额（游戏币）
	Coin int32 `gorm:"not null;column:coin" json:"coin"`

	// State 0 已充值未领取,1 已充值已领取
	State int32 `gorm:"not null;column:state" json:"state"`

	// Type 充值类型（0 模拟,1 真实,2 赠送,3 物品兑换,4 红包,5 礼包码）
	Type int32 `gorm:"column:type" json:"type"`

	// Gid 角色 id
	Gid string `gorm:"size:128;column:gid" json:"gid"`

	// AddTime 创建时间
	AddTime time.Time `gorm:"autoCreateTime;column:add_time" json:"addTime"`

	// UpdateTime 更新时间
	UpdateTime time.Time `gorm:"autoUpdateTime;column:update_time" json:"updateTime"`

	// Deleted 逻辑删除标记（tinyint(1)）
	Deleted bool `gorm:"column:deleted" json:"deleted"`

	// Money 充值金额（人民币）
	Money int32 `gorm:"column:money" json:"money"`

	// Code 注册码
	Code string `gorm:"size:255;column:code" json:"code"`

	// Remark 备注
	Remark string `gorm:"size:255;column:remark" json:"remark"`

	// OrderNo 订单号
	OrderNo string `gorm:"size:255;column:order_no" json:"orderNo"`

	// OnceGiveStatus 单笔领取状态
	OnceGiveStatus int32 `gorm:"column:once_give_status" json:"onceGiveStatus"`
}

// TableName 显式指定表名
func (Charge) TableName() string {
	return "charge"
}
