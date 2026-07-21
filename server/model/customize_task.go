package model

// CustomizeTask 自定义任务表。
// taskPrompt 在 Java 实体中声明但未出现在 wd-game-test.sql 建表语句中，按普通列保留。
// taskTimeLimit 在 Java 中为 Integer，SQL 中为 varchar，按 Java 类型映射为 int32。
type CustomizeTask struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// TaskName 任务名称
	TaskName string `gorm:"size:1024;column:task_name" json:"taskName"`

	// TaskPrompt 任务提示
	TaskPrompt string `gorm:"column:task_prompt" json:"taskPrompt"`

	// TaskDesc 任务详情
	TaskDesc string `gorm:"size:4096;column:task_desc" json:"taskDesc"`

	// TaskRewardDesc 任务奖励描述
	TaskRewardDesc string `gorm:"size:1024;column:task_reward_desc" json:"taskRewardDesc"`

	// NpcName 展示任务的 NPC 名字
	NpcName string `gorm:"size:64;column:npc_name" json:"npcName"`

	// NpcMenu 展示任务的 NPC 选项
	NpcMenu string `gorm:"size:1024;column:npc_menu" json:"npcMenu"`

	// ReceiveWeek 任务可领取的星期
	ReceiveWeek string `gorm:"size:128;column:receive_week" json:"receiveWeek"`

	// ReceiveMaxSize 任务每日领取最大次数
	ReceiveMaxSize int32 `gorm:"column:receive_max_size" json:"receiveMaxSize"`

	// ReceiveConsume 领取任务的消耗品
	ReceiveConsume string `gorm:"size:1024;column:receive_consume" json:"receiveConsume"`

	// TaskType 任务类型
	TaskType string `gorm:"size:32;column:task_type" json:"taskType"`

	// TaskCondition 任务完成条件
	TaskCondition string `gorm:"size:1024;column:task_condition" json:"taskCondition"`

	// TaskTriggerCondition 任务触发条件
	TaskTriggerCondition string `gorm:"size:512;column:task_trigger_condition" json:"taskTriggerCondition"`

	// TaskTimeLimit 任务时限
	TaskTimeLimit int32 `gorm:"column:task_time_limit" json:"taskTimeLimit"`

	// TaskReward 任务完成奖励
	TaskReward string `gorm:"size:1024;column:task_reward" json:"taskReward"`
}

// TableName 显式指定表名
func (CustomizeTask) TableName() string {
	return "customize_task"
}
