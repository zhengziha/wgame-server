// Package model 定义 GORM 持久化模型。
//
// 所有模型使用大写表名（GORM 默认规则）+ int64 主键，
// 字段 tag 同时包含 gorm/json/redis cache 元信息，便于 DAO 层统一处理。
package model

import "time"

// User 用户基础表，仅作为 ORM 演示模型。
// 真实业务可在同一目录继续追加 model 文件。
type User struct {
	// ID 主键，业务层用雪花算法/自增均可
	ID int64 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// AccountName 登录账号（唯一）
	AccountName string `gorm:"uniqueIndex;size:64;not null;column:account_name" json:"accountName"`

	// Nickname 游戏内昵称
	Nickname string `gorm:"size:64;not null;column:nickname" json:"nickname"`

	// Level 玩家等级
	Level int `gorm:"default:1;column:level" json:"level"`

	// CreatedAt 创建时间
	CreatedAt time.Time `gorm:"autoCreateTime;column:created_at" json:"createdAt"`

	// UpdatedAt 更新时间
	UpdatedAt time.Time `gorm:"autoUpdateTime;column:updated_at" json:"updatedAt"`
}

// TableName 显式指定表名，避免 GORM 复数规则把 User -> users
func (User) TableName() string {
	return "t_user"
}
