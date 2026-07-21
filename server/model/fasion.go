package model

import "time"

// Fasion 时装/包装外观表（对应数据库表 pack_modification）。
type Fasion struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// Alias 组合名字带日期
	Alias string `gorm:"size:32;column:alias" json:"alias"`

	// FasionType 身上显示
	FasionType string `gorm:"size:32;column:fasion_type" json:"fasionType"`

	// Str 名字
	Str string `gorm:"size:32;column:str" json:"str"`

	// Type 穿戴栏显示/商品图标 icon
	Type string `gorm:"size:255;column:type" json:"type"`

	// FoodNum 未使用字段
	FoodNum int32 `gorm:"column:food_num" json:"foodNum"`

	// GoodsPrice 价格
	GoodsPrice int32 `gorm:"column:goods_price" json:"goodsPrice"`

	// Sex 0男1女（跟宠:1可飞行,0不可需购买腾云）
	Sex int32 `gorm:"column:sex" json:"sex"`

	// Position 背包位置
	Position int32 `gorm:"column:position" json:"position"`

	// Category 部位（外观特效跟宠物对应 1/2/3）
	Category int32 `gorm:"column:category" json:"category"`

	// AddTime 创建时间
	AddTime time.Time `gorm:"autoCreateTime;column:add_time" json:"addTime"`

	// UpdateTime 更新时间
	UpdateTime time.Time `gorm:"autoUpdateTime;column:update_time" json:"updateTime"`

	// Deleted 逻辑删除标记（tinyint(1)）
	Deleted bool `gorm:"column:deleted" json:"deleted"`
}

// TableName 显式指定表名
func (Fasion) TableName() string {
	return "pack_modification"
}
