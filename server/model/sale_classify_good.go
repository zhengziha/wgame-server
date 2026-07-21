package model

import "time"

// SaleClassifyGood 交易分类商品表。
type SaleClassifyGood struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// Icon 图标
	Icon int32 `gorm:"column:icon" json:"icon"`

	// Name 商品名称
	Name string `gorm:"size:32;column:name" json:"name"`

	// Price 价格
	Price int32 `gorm:"column:price" json:"price"`

	// Compose 组合描述
	Compose string `gorm:"size:64;column:compose" json:"compose"`

	// PublicityTime 公示时间
	PublicityTime int32 `gorm:"column:publicity_time" json:"publicityTime"`

	// AddTime 创建时间
	AddTime time.Time `gorm:"autoCreateTime;column:add_time" json:"addTime"`

	// UpdateTime 更新时间
	UpdateTime time.Time `gorm:"autoUpdateTime;column:update_time" json:"updateTime"`

	// Deleted 逻辑删除标记（tinyint(1)）
	Deleted bool `gorm:"column:deleted" json:"deleted"`
}

// TableName 显式指定表名
func (SaleClassifyGood) TableName() string {
	return "sale_classify_good"
}
