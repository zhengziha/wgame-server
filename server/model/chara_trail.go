package model

import "time"

// CharaTrail 玩家活动轨迹表。
type CharaTrail struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// Data 数据
	Data string `gorm:"size:255;column:data" json:"data"`

	// Remarks 标记（例如：增加积分、经验、道行、武学）
	Remarks string `gorm:"size:15;column:remarks" json:"remarks"`

	// Source 类型（例如：积分、道行）
	Source string `gorm:"size:15;column:source" json:"source"`

	// CharaName 角色名（带等级）
	CharaName string `gorm:"size:55;column:chara_name" json:"charaName"`

	// Cid 角色 id
	Cid int32 `gorm:"column:cid" json:"cid"`

	// Gid 角色 uuid
	Gid string `gorm:"size:255;column:gid" json:"gid"`

	// AddTime 创建时间
	AddTime time.Time `gorm:"autoCreateTime;column:add_time" json:"addTime"`
}

// TableName 显式指定表名
func (CharaTrail) TableName() string {
	return "chara_trail"
}
