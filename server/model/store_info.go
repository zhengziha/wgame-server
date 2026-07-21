package model

import "time"

// StoreInfo 商店物品信息表。
// 注意：Combine 为非持久化字段，数据库中不存在对应列。
type StoreInfo struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// Quality 品质
	Quality string `gorm:"size:255;column:quality" json:"quality"`

	// Value 物品 value
	Value int32 `gorm:"column:value" json:"value"`

	// Type 物品类型
	Type int32 `gorm:"column:type" json:"type"`

	// Name 物品名称
	Name string `gorm:"size:255;column:name" json:"name"`

	// TotalScore item_type
	TotalScore int32 `gorm:"column:total_score" json:"totalScore"`

	// RecognizeRecognized gift 标识
	RecognizeRecognized int32 `gorm:"column:recognize_recognized" json:"recognizeRecognized"`

	// RebuildLevel 售卖金额
	RebuildLevel int32 `gorm:"column:rebuild_level" json:"rebuildLevel"`

	// SilverCoin 灵气
	SilverCoin int32 `gorm:"column:silver_coin" json:"silverCoin"`

	// Combine 叠加数量（非持久化字段，数据库无此列）
	Combine int32 `gorm:"-" json:"combine"`

	// AddTime 创建时间
	AddTime time.Time `gorm:"autoCreateTime;column:add_time" json:"addTime"`

	// UpdateTime 更新时间
	UpdateTime time.Time `gorm:"autoUpdateTime;column:update_time" json:"updateTime"`

	// Deleted 逻辑删除标记（tinyint(1)）
	Deleted bool `gorm:"column:deleted" json:"deleted"`
}

// TableName 显式指定表名
func (StoreInfo) TableName() string {
	return "store_info"
}
