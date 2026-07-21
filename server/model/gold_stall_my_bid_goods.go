package model

// GoldStallMyBidGoods 珍宝阁我竞拍的物品表。
type GoldStallMyBidGoods struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// Gid 角色 uuid
	Gid string `gorm:"size:64;column:gid" json:"gid"`

	// GoodsId 物品 uuid
	GoodsId string `gorm:"size:64;column:goods_id" json:"goodsId"`

	// BuyoutPrice 竞价价格
	BuyoutPrice int32 `gorm:"column:buyout_price" json:"buyoutPrice"`

	// Deleted 逻辑删除标记（tinyint）
	Deleted bool `gorm:"column:deleted" json:"deleted"`
}

// TableName 显式指定表名
func (GoldStallMyBidGoods) TableName() string {
	return "gold_stall_my_bid_goods"
}
