package model

// GrowthTask 成长任务表。
type GrowthTask struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// TaskName 任务名称
	TaskName string `gorm:"size:128;column:task_name" json:"taskName"`

	// TaskDesc 任务详情
	TaskDesc string `gorm:"size:1024;column:task_desc" json:"taskDesc"`

	// Exp 任务经验
	Exp int32 `gorm:"column:exp" json:"exp"`

	// Num 任务数量
	Num int32 `gorm:"column:num" json:"num"`
}

// TableName 显式指定表名
func (GrowthTask) TableName() string {
	return "growth_task"
}
