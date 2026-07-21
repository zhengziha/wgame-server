package model

import "time"

// Dialog 对话/申请表。
type Dialog struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// PeerName 申请人
	PeerName string `gorm:"size:255;column:peer_name" json:"peerName"`

	// AskType 申请类型
	AskType string `gorm:"size:255;column:ask_type" json:"askType"`

	// ApplyGid 申请 gid
	ApplyGid string `gorm:"size:255;column:apply_gid" json:"applyGid"`

	// CreateTime 创建时间
	CreateTime time.Time `gorm:"autoCreateTime;column:create_time" json:"createTime"`

	// ExtJson 拓展 json 信息
	ExtJson string `gorm:"type:longtext;column:ext_json" json:"extJson"`
}

// TableName 显式指定表名
func (Dialog) TableName() string {
	return "dialog"
}
