package model

import "time"

// AuctionGoods 拍卖物品表。
// maxPrice 在 Java 中为 Integer，SQL 中为 decimal(10,2)，按 Java 类型映射为 int32。
type AuctionGoods struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// Name 拍卖的物品名称
	Name string `gorm:"size:255;not null;column:name" json:"name"`

	// Price 拍卖价格
	Price int32 `gorm:"not null;column:price" json:"price"`

	// AuctionTime 竞拍时间，单位分钟
	AuctionTime int32 `gorm:"column:auction_time" json:"auctionTime"`

	// GoodStatus 物品状态
	GoodStatus int32 `gorm:"column:good_status" json:"goodStatus"`

	// Uuid 最高价角色 UUID
	Uuid string `gorm:"size:64;column:uuid" json:"uuid"`

	// MaxPrice 最高价
	MaxPrice int32 `gorm:"column:max_price" json:"maxPrice"`

	// AuctionType 支付类型
	AuctionType string `gorm:"size:255;column:auction_type" json:"auctionType"`

	// AddTime 创建时间
	AddTime time.Time `gorm:"autoCreateTime;column:add_time" json:"addTime"`
}

// TableName 显式指定表名
func (AuctionGoods) TableName() string {
	return "auction_goods"
}
