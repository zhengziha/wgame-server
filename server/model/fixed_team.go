package model

import "time"

// FixedTeam 固定队伍信息表。
type FixedTeam struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// Name 固定队伍名字
	Name string `gorm:"size:50;column:name" json:"name"`

	// LeaderUid 固定队长 uid
	LeaderUid string `gorm:"size:50;column:leader_uid" json:"leaderUid"`

	// Uid 唯一 id（唯一索引）
	Uid string `gorm:"size:50;uniqueIndex;column:uid" json:"uid"`

	// Level 等级
	Level int32 `gorm:"column:level" json:"level"`

	// Intimacy 亲密度
	Intimacy int32 `gorm:"column:intimacy" json:"intimacy"`

	// AddTime 创建时间
	AddTime time.Time `gorm:"autoCreateTime;column:add_time" json:"addTime"`

	// Members 队员 JSON 信息
	Members string `gorm:"type:text;column:members" json:"members"`
}

// TableName 显式指定表名
func (FixedTeam) TableName() string {
	return "fixed_team"
}
