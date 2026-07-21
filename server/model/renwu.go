package model

import "time"

// Renwu 任务表。
// 部分字段（taskType、camp、extra、nextCallType、nextCallId、limitCount）
// 在 Java 实体中声明但未出现在 wd-game-18.sql 建表语句中，按普通列保留。
type Renwu struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// Uncontent 任务npc菜单
	Uncontent string `gorm:"size:255;column:uncontent" json:"uncontent"`

	// NpcName npc名字，查询菜单用
	NpcName string `gorm:"size:255;column:npc_name" json:"npcName"`

	// CurrentTask 当前任务标识
	CurrentTask string `gorm:"size:255;column:current_task" json:"currentTask"`

	// TaskType 任务类型
	TaskType string `gorm:"size:255;column:task_type" json:"taskType"`

	// ShowName 任务显示名字
	ShowName string `gorm:"size:255;column:show_name" json:"showName"`

	// TaskPrompt 任务提示
	TaskPrompt string `gorm:"size:255;column:task_prompt" json:"taskPrompt"`

	// TaskDesc 任务描述
	TaskDesc string `gorm:"size:255;column:task_desc" json:"taskDesc"`

	// Reward 服务器处理奖励
	Reward string `gorm:"size:255;column:reward" json:"reward"`

	// Camp 阵营名字
	Camp string `gorm:"size:255;column:camp" json:"camp"`

	// ShowReward 显示奖励
	ShowReward string `gorm:"size:255;column:show_reward" json:"showReward"`

	// Attrib 属性标记（0 不可放弃，1 可放弃等，详见 Java 注释）
	Attrib int32 `gorm:"column:attrib" json:"attrib"`

	// TaskEndTime 任务结束时间（秒级时间戳）
	TaskEndTime int32 `gorm:"column:task_end_time" json:"taskEndTime"`

	// TaskState 任务状态
	TaskState string `gorm:"size:255;column:task_state" json:"taskState"`

	// Extra 额外参数
	Extra string `gorm:"size:255;column:extra" json:"extra"`

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

	// LimitCount 任务限制完成次数
	LimitCount int32 `gorm:"column:limit_count" json:"limitCount"`
}

// TableName 显式指定表名
func (Renwu) TableName() string {
	return "renwu"
}
