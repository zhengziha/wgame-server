package model

// PartyShop 帮派商店物品表。
// 对应建表语句位于 wd-game-test.sql。
type PartyShop struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// Name 物品名称
	Name string `gorm:"size:255;column:name" json:"name"`

	// Num 数量
	Num int32 `gorm:"column:num" json:"num"`

	// Price 价格
	Price int32 `gorm:"column:price" json:"price"`
}

// TableName 显式指定表名
func (PartyShop) TableName() string {
	return "party_shop"
}
