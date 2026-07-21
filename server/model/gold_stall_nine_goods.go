package model

import "time"

// GoldStallNineGoods 珍宝阁九宫格物品（珍宝摆摊）表。
// status、level 为 SQL 保留字，column tag 直接写列名，GORM 会自动加反引号。
type GoldStallNineGoods struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// Name 物品名字
	Name string `gorm:"size:50;column:name" json:"name"`

	// GoodsId 物品 id
	GoodsId string `gorm:"size:255;column:goods_id" json:"goodsId"`

	// Price 价格
	Price int32 `gorm:"column:price" json:"price"`

	// Status 状态（列名 status 为 SQL 保留字）
	// 0无状态 1公示中 2出售中 3已下架 4冻结中 5审核中 6已审核 7审核不通过
	Status int32 `gorm:"column:status" json:"status"`

	// Alias 物品 key
	Alias string `gorm:"size:255;column:alias" json:"alias"`

	// ReqLevel 要求等级
	ReqLevel int32 `gorm:"column:req_level" json:"reqLevel"`

	// Gid 卖家 gid
	Gid string `gorm:"size:255;column:gid" json:"gid"`

	// UpdateTime 更新时间
	UpdateTime time.Time `gorm:"autoUpdateTime;column:update_time" json:"updateTime"`

	// Level 等级（列名 level 为 SQL 保留字）
	Level int32 `gorm:"column:level" json:"level"`

	// StallItemType 类型 0装备道具 2宠物
	StallItemType int32 `gorm:"column:stall_item_type" json:"stallItemType"`

	// Extra 扩展字段
	Extra string `gorm:"size:255;column:extra" json:"extra"`

	// ItemPolar 相性
	ItemPolar int32 `gorm:"column:item_polar" json:"itemPolar"`

	// CgPriceCount 剩余改价次数
	CgPriceCount int32 `gorm:"column:cg_price_count" json:"cgPriceCount"`

	// Unidentified 是否未鉴定 0默认 1未鉴定
	Unidentified int32 `gorm:"column:unidentified" json:"unidentified"`

	// StartTime 开始时间
	StartTime int32 `gorm:"column:start_time" json:"startTime"`

	// EndTime 结束时间
	EndTime int32 `gorm:"column:end_time" json:"endTime"`

	// InitPrice 初始价格
	InitPrice int32 `gorm:"column:init_price" json:"initPrice"`

	// FlagNum 标识
	FlagNum int32 `gorm:"column:flag_num" json:"flagNum"`

	// BuyoutPrice 一口价（拍卖价格）
	BuyoutPrice int32 `gorm:"column:buyout_price" json:"buyoutPrice"`

	// SellType 指定类型 0未指定 1指定 5拍卖
	SellType int32 `gorm:"column:sell_type" json:"sellType"`

	// AppointeeName 指定人名称
	AppointeeName string `gorm:"size:255;column:appointee_name" json:"appointeeName"`

	// Deleted 逻辑删除标记（tinyint(1)）
	Deleted bool `gorm:"column:deleted" json:"deleted"`

	// AddTime 创建时间
	AddTime time.Time `gorm:"autoCreateTime;column:add_time" json:"addTime"`

	// Goods 装备和物品 JSON
	Goods string `gorm:"type:mediumtext;column:goods" json:"goods"`

	// Master 物品所属人
	Master string `gorm:"size:20;column:master" json:"master"`
}

// TableName 显式指定表名
func (GoldStallNineGoods) TableName() string {
	return "gold_stall_nine_goods"
}
