package model

// GoodShop 好心值商店物品表。
// 对应建表语句位于 wd-game-test.sql。
// descript 在 Java 实体中声明，wd-game-test.sql 建表语句未出现该列，按普通列保留以兼容 Java 逻辑。
type GoodShop struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// Type 商店类型（0天技秘笈，2好心值商店，3阴气之尘商店）
	Type int32 `gorm:"column:type" json:"type"`

	// Name 物品名称
	Name string `gorm:"size:255;column:name" json:"name"`

	// Descript 道具描述
	Descript string `gorm:"type:text;column:descript" json:"descript"`

	// Num 数量
	Num int32 `gorm:"column:num" json:"num"`

	// Price 价格
	Price int32 `gorm:"column:price" json:"price"`
}

// TableName 显式指定表名
func (GoodShop) TableName() string {
	return "good_shop"
}
