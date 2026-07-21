package model

import "time"

// StallRecord 摆摊（交易）记录表。
// status/level/type 为 SQL 保留字，column tag 直接写列名，GORM 会自动加反引号。
type StallRecord struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// Gid 用户 gid
	Gid string `gorm:"size:55;column:gid" json:"gid"`

	// Cid 用户 id
	Cid int32 `gorm:"column:cid" json:"cid"`

	// OwnerName 所属人名字
	OwnerName string `gorm:"size:15;column:owner_name" json:"ownerName"`

	// Status 状态（0:交易完成 5:审核中 6:已审核 7:审核不通过 11:拍卖公示期 12:拍卖出售期 13:拍卖付款期）
	Status int32 `gorm:"column:status" json:"status"`

	// StallRecordType 记录类型（0:集市 1:珍宝）
	StallRecordType int32 `gorm:"column:stall_record_type" json:"stallRecordType"`

	// GoodsUuid 商品 uid
	GoodsUuid string `gorm:"size:55;column:goods_uuid" json:"goodsUuid"`

	// GoodsName 商品名称
	GoodsName string `gorm:"size:25;column:goods_name" json:"goodsName"`

	// ItemType 商品类型
	ItemType int32 `gorm:"column:item_type" json:"itemType"`

	// EndTime 结束时间
	EndTime time.Time `gorm:"column:end_time" json:"endTime"`

	// Price 价格
	Price int32 `gorm:"column:price" json:"price"`

	// ReqLevel 要求等级
	ReqLevel int32 `gorm:"column:req_level" json:"reqLevel"`

	// ItemPolar 商品属性
	ItemPolar int32 `gorm:"column:item_polar" json:"itemPolar"`

	// BuyType 购买类型
	BuyType int32 `gorm:"column:buy_type" json:"buyType"`

	// Level 等级
	Level int32 `gorm:"column:level" json:"level"`

	// AddTime 订单时间
	AddTime time.Time `gorm:"autoCreateTime;column:add_time" json:"addTime"`

	// Data 数据 JSON
	Data string `gorm:"type:mediumtext;column:data" json:"data"`

	// Type 类型（0:出售记录 1:购买记录）
	Type int32 `gorm:"column:type" json:"type"`

	// BuyName 购买人名字
	BuyName string `gorm:"size:15;column:buy_name" json:"buyName"`

	// BuyGid 购买人的 gid
	BuyGid string `gorm:"size:55;column:buy_gid" json:"buyGid"`
}

// TableName 显式指定表名
func (StallRecord) TableName() string {
	return "stall_record"
}
