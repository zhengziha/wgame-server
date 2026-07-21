package model

import "time"

// GroceriesShop 杂货店表。
// Level/Type 列名为 SQL 保留字；Itemcount 列名在数据库中为驼峰 itemCount。
type GroceriesShop struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// GoodsNo 物品编号
	GoodsNo int32 `gorm:"column:goods_no" json:"goodsNo"`

	// PayType 支付类型（8 只能金币，其它代金券或金币）
	PayType int32 `gorm:"column:pay_type" json:"payType"`

	// Name 物品名称
	Name string `gorm:"size:255;column:name" json:"name"`

	// Value 支付金额
	Value int32 `gorm:"column:value" json:"value"`

	// Level 等级（列名 level 为 SQL 保留字）
	Level int32 `gorm:"column:level" json:"level"`

	// Type 类型（2 妖石, 1 道具）（列名 type 为 SQL 保留字）
	Type int32 `gorm:"column:type" json:"type"`

	// Itemcount 物品数量（数据库列名为驼峰 itemCount）
	Itemcount int32 `gorm:"column:itemCount" json:"itemcount"`

	// AddTime 创建时间
	AddTime time.Time `gorm:"autoCreateTime;column:add_time" json:"addTime"`

	// UpdateTime 更新时间
	UpdateTime time.Time `gorm:"autoUpdateTime;column:update_time" json:"updateTime"`

	// Deleted 逻辑删除标记（tinyint(1)）
	Deleted bool `gorm:"column:deleted" json:"deleted"`
}

// TableName 显式指定表名
func (GroceriesShop) TableName() string {
	return "groceries_shop"
}
