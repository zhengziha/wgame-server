package model

import "time"

// Daili 代理（管理员）表，建表语句位于 wd-auth-18.sql。
// chargeMoney/maxChargeMoney 在 Java 实体中声明但未出现在建表语句中，
// 按 charge.go / characters.go 的处理方式，保留为普通列以兼容 Java 逻辑。
type Daili struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// Account 登录账号
	Account string `gorm:"size:32;not null;column:account" json:"account"`

	// Passwd 密码
	Passwd string `gorm:"size:255;not null;column:passwd" json:"passwd"`

	// Code 注册码
	Code string `gorm:"size:32;not null;column:code" json:"code"`

	// Token 令牌
	Token string `gorm:"size:255;column:token" json:"token"`

	// ChargeLink 充值链接
	ChargeLink string `gorm:"size:255;column:charge_link" json:"chargeLink"`

	// Zdsl 身份标识（0:子代 1:总代 2:管理员）
	Zdsl int32 `gorm:"default:0;column:zdsl" json:"zdsl"`

	// ChargeMoney 充值金额
	ChargeMoney int32 `gorm:"column:charge_money" json:"chargeMoney"`

	// MaxChargeMoney 最大充值金额
	MaxChargeMoney int32 `gorm:"column:max_charge_money" json:"maxChargeMoney"`

	// AddTime 创建时间
	AddTime time.Time `gorm:"autoCreateTime;column:add_time" json:"addTime"`

	// UpdateTime 更新时间
	UpdateTime time.Time `gorm:"autoUpdateTime;column:update_time" json:"updateTime"`

	// Deleted 逻辑删除标记（tinyint(1)）
	Deleted bool `gorm:"column:deleted" json:"deleted"`
}

// TableName 显式指定表名
func (Daili) TableName() string {
	return "daili"
}
