package model

import "time"

// ShidaoHistory 试道历史表。
// level/rank 为 SQL 保留字，column tag 直接写列名，GORM 会自动加反引号。
// isMonth 字段仅出现在 wd-game-test.sql 建表语句中，wd-game-18.sql 未包含，
// 此处按 Java 实体保留以兼容业务逻辑。
type ShidaoHistory struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// Leader 队长名字
	Leader string `gorm:"size:15;column:leader" json:"leader"`

	// LeaderUuid 队长 uuid
	LeaderUuid string `gorm:"size:255;column:leader_uuid" json:"leaderUuid"`

	// Level 等级阶段
	Level int32 `gorm:"column:level" json:"level"`

	// Rank 排名
	Rank int32 `gorm:"column:rank" json:"rank"`

	// Score 队伍试道总积分
	Score int32 `gorm:"column:score" json:"score"`

	// TotalTao 队伍总道行
	TotalTao int32 `gorm:"column:total_tao" json:"totalTao"`

	// ShidaoTime 试道时间（int 时间戳）
	ShidaoTime int32 `gorm:"column:shidao_time" json:"shidaoTime"`

	// CreateTime 创建时间
	CreateTime time.Time `gorm:"autoCreateTime;column:create_time" json:"createTime"`

	// IsMonth 是否月试道（0:否 1:是）
	IsMonth int32 `gorm:"column:is_month" json:"isMonth"`
}

// TableName 显式指定表名
func (ShidaoHistory) TableName() string {
	return "shidao_history"
}
