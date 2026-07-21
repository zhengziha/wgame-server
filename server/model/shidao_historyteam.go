package model

import "time"

// ShidaoHistoryteam 试道历史队伍信息表。
// level 为 SQL 保留字，column tag 直接写列名，GORM 会自动加反引号。
// drawReward/drawDateTime 字段在 wd-game-test.sql 中出现，wd-game-18.sql 未包含，
// 此处按 Java 实体保留以兼容业务逻辑。
type ShidaoHistoryteam struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// ShidaoHistoryId 试道历史记录表主键
	ShidaoHistoryId int32 `gorm:"column:shidao_history_id" json:"shidaoHistoryId"`

	// Name 名称
	Name string `gorm:"size:15;column:name" json:"name"`

	// Level 等级
	Level int32 `gorm:"column:level" json:"level"`

	// Family 是否有家
	Family int32 `gorm:"column:family" json:"family"`

	// Gid UUID
	Gid string `gorm:"size:255;column:gid" json:"gid"`

	// Icon 头像
	Icon int32 `gorm:"column:icon" json:"icon"`

	// Tao 道行
	Tao int32 `gorm:"column:tao" json:"tao"`

	// DrawReward 是否领取奖励（null:无奖励 0:未领取 1:已领取）
	DrawReward int32 `gorm:"column:draw_reward" json:"drawReward"`

	// DrawDateTime 奖励领取时间
	DrawDateTime time.Time `gorm:"column:draw_date_time" json:"drawDateTime"`
}

// TableName 显式指定表名
func (ShidaoHistoryteam) TableName() string {
	return "shidao_history_team"
}
