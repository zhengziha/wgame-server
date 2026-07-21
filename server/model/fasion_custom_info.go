package model

import "time"

// FasionCustomInfo 自定义时装信息表。
type FasionCustomInfo struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// EquipPos 组合名字带日期
	EquipPos int32 `gorm:"column:equip_pos" json:"equipPos"`

	// FasionPart 身上显示
	FasionPart int32 `gorm:"column:fasion_part" json:"fasionPart"`

	// FasionDye 穿戴栏显示/商品图标 icon
	FasionDye int32 `gorm:"column:fasion_dye" json:"fasionDye"`

	// Name 名字
	Name string `gorm:"size:32;column:name" json:"name"`

	// Gift 是否赠品
	Gift int32 `gorm:"column:gift" json:"gift"`

	// Icon 图标
	Icon int32 `gorm:"column:icon" json:"icon"`

	// GoodsPrice 价格
	GoodsPrice int32 `gorm:"column:goods_price" json:"goodsPrice"`

	// Sex 男女（0男 1女 2不限制）
	Sex int32 `gorm:"column:sex" json:"sex"`

	// Position 显示位置
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
func (FasionCustomInfo) TableName() string {
	return "fasion_custom_info"
}
