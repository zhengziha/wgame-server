package model

import "time"

// ConfigInfo 配置信息表。
type ConfigInfo struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// Uuid 唯一 ID
	Uuid string `gorm:"size:255;column:uuid" json:"uuid"`

	// KeyName 配置名称
	KeyName string `gorm:"size:168;column:key_name" json:"keyName"`

	// AliasName 别名
	AliasName string `gorm:"size:255;column:alias_name" json:"aliasName"`

	// AddTime 创建时间
	AddTime time.Time `gorm:"autoCreateTime;column:add_time" json:"addTime"`

	// Data 配置数据（longtext）
	Data string `gorm:"type:longtext;column:data" json:"data"`
}

// TableName 显式指定表名
func (ConfigInfo) TableName() string {
	return "config_info"
}
