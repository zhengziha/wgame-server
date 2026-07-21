package model

import "time"

// StoreGoods 商店商品表。
type StoreGoods struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// Name 商品名称
	Name string `gorm:"size:255;column:name" json:"name"`

	// Barcode 条形码
	Barcode string `gorm:"size:255;column:barcode" json:"barcode"`

	// ForSale 购买类型（0:银元宝,1:金元宝,2:通用）
	ForSale int32 `gorm:"not null;column:for_sale" json:"forSale"`

	// ShowPos 展示位置
	ShowPos int32 `gorm:"column:show_pos" json:"showPos"`

	// Rpos 推荐位置
	Rpos int32 `gorm:"column:rpos" json:"rpos"`

	// SaleQuota 销售配额
	SaleQuota int32 `gorm:"column:sale_quota" json:"saleQuota"`

	// Recommend 是否推荐
	Recommend int32 `gorm:"column:recommend" json:"recommend"`

	// Coin 金币
	Coin int32 `gorm:"column:coin" json:"coin"`

	// Discount 折扣
	Discount int32 `gorm:"column:discount" json:"discount"`

	// Type 商品类型
	Type int32 `gorm:"column:type" json:"type"`

	// QuotaLimit 配额上限
	QuotaLimit int32 `gorm:"column:quota_limit" json:"quotaLimit"`

	// MustVip 是否需要 VIP
	MustVip int32 `gorm:"column:must_vip" json:"mustVip"`

	// IsGift 是否赠品
	IsGift int32 `gorm:"column:is_gift" json:"isGift"`

	// FollowPetType 跟宠类型
	FollowPetType int32 `gorm:"column:follow_pet_type" json:"followPetType"`

	// AddTime 创建时间
	AddTime time.Time `gorm:"autoCreateTime;column:add_time" json:"addTime"`

	// UpdateTime 更新时间
	UpdateTime time.Time `gorm:"autoUpdateTime;column:update_time" json:"updateTime"`

	// Deleted 逻辑删除标记（tinyint(1)）
	Deleted bool `gorm:"column:deleted" json:"deleted"`
}

// TableName 显式指定表名
func (StoreGoods) TableName() string {
	return "store_goods"
}
