package model

import "time"

// SaleGood 交易商品表。
// Level/Status 列名为 SQL 保留字。
type SaleGood struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// GoodsID 物品 ID
	GoodsID string `gorm:"size:225;column:goods_id" json:"goodsId"`

	// Name 物品名字
	Name string `gorm:"size:255;column:name" json:"name"`

	// Alias 别名
	Alias string `gorm:"size:255;column:alias" json:"alias"`

	// Price 价格
	Price int32 `gorm:"column:price" json:"price"`

	// ReqLevel 要求等级
	ReqLevel int32 `gorm:"column:req_level" json:"reqLevel"`

	// Gid 玩家 gid
	Gid string `gorm:"size:255;column:gid" json:"gid"`

	// AddTime 创建时间
	AddTime time.Time `gorm:"autoCreateTime;column:add_time" json:"addTime"`

	// UpdateTime 更新时间
	UpdateTime time.Time `gorm:"autoUpdateTime;column:update_time" json:"updateTime"`

	// Deleted 逻辑删除标记（tinyint(1)）
	Deleted bool `gorm:"column:deleted" json:"deleted"`

	// Level 等级（列名 level 为 SQL 保留字）
	Level int32 `gorm:"column:level" json:"level"`

	// Type 类型（1:道具, 2:宠物）
	Type int32 `gorm:"column:type" json:"type"`

	// Extra 扩展字段
	Extra string `gorm:"size:255;column:extra" json:"extra"`

	// Status 状态（列名 status 为 SQL 保留字）
	Status int32 `gorm:"column:status" json:"status"`

	// StartTime 开始时间
	StartTime int32 `gorm:"column:start_time" json:"startTime"`

	// EndTime 结束时间
	EndTime int32 `gorm:"column:end_time" json:"endTime"`

	// Goods 装备和物品（mediumtext）
	Goods string `gorm:"type:mediumtext;column:goods" json:"goods"`

	// Icon 图标
	Icon int32 `gorm:"column:icon" json:"icon"`

	// Unidentified 是否未鉴定
	Unidentified int32 `gorm:"column:unidentified" json:"unidentified"`

	// ItemPolar 相性
	ItemPolar int32 `gorm:"column:item_polar" json:"itemPolar"`

	// CgPriceCount 价格计算配置
	CgPriceCount int32 `gorm:"column:cg_price_count" json:"cgPriceCount"`

	// SgID 物品分类 ID
	SgID int32 `gorm:"column:sg_id" json:"sgId"`
}

// TableName 显式指定表名
func (SaleGood) TableName() string {
	return "sale_good"
}
