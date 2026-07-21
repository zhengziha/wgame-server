package model

// ShengwangShop 声望商店表。
// 对应 Java 实体 T_shengwang_shop，Go 结构体命名按规范去掉 T_ 前缀。
type ShengwangShop struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// Key1 商品 key
	Key1 string `gorm:"size:255;column:key1" json:"key1"`

	// Name 商品名称
	Name string `gorm:"size:255;column:name" json:"name"`

	// Price 价格
	Price int32 `gorm:"column:price" json:"price"`
}

// TableName 显式指定表名（保留原表前缀 t_）
func (ShengwangShop) TableName() string {
	return "t_shengwang_shop"
}
