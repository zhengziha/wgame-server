package model

import "time"

// TaskFightInfo 任务战斗信息表。
// 注意：该表未出现在 wd-game-18.sql，表名与列结构依据
// wd-game-test.sql 中的 task_fight_info 建表语句确定。
type TaskFightInfo struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// CurrentTask 当前任务标识
	CurrentTask string `gorm:"size:255;column:current_task" json:"currentTask"`

	// TaskType 任务类型
	TaskType string `gorm:"size:255;column:task_type" json:"taskType"`

	// MonsterList 怪物列表
	MonsterList string `gorm:"size:255;column:monster_list" json:"monsterList"`

	// ShouHu 是否守护
	ShouHu bool `gorm:"column:shou_hu" json:"shouHu"`

	// TeamSize 队伍人数
	TeamSize int32 `gorm:"column:team_size" json:"teamSize"`

	// NextCallType 下一步调用类型
	NextCallType string `gorm:"size:255;column:next_call_type" json:"nextCallType"`

	// NextCallId 下一步调用 id
	NextCallId string `gorm:"size:255;column:next_call_id" json:"nextCallId"`

	// AddTime 创建时间
	AddTime time.Time `gorm:"autoCreateTime;column:add_time" json:"addTime"`

	// UpdateTime 更新时间
	UpdateTime time.Time `gorm:"autoUpdateTime;column:update_time" json:"updateTime"`

	// Deleted 逻辑删除标记（tinyint(1)）
	Deleted bool `gorm:"column:deleted" json:"deleted"`
}

// TableName 显式指定表名
func (TaskFightInfo) TableName() string {
	return "task_fight_info"
}
