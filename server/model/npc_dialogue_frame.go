package model

import "time"

// NpcDialogueFrame NPC 对话框表。
// 注意：NextCallType/NextCallID/Camp 为非持久化字段，数据库中不存在对应列。
// Next 列名为 SQL 保留字。
type NpcDialogueFrame struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// Content 上级菜单内容
	Content string `gorm:"size:255;column:content" json:"content"`

	// Portrait 头像编号
	Portrait int32 `gorm:"column:portrait" json:"portrait"`

	// Name 名称
	Name string `gorm:"size:255;column:name" json:"name"`

	// CurrentTask 特殊任务名称
	CurrentTask string `gorm:"size:255;column:current_task" json:"currentTask"`

	// Uncontent 菜单内容
	Uncontent string `gorm:"type:text;column:uncontent" json:"uncontent"`

	// NextCallType 下一步调用类型（非持久化字段，数据库无此列）
	NextCallType string `gorm:"-" json:"nextCallType"`

	// NextCallID 下一步调用 ID（非持久化字段，数据库无此列）
	NextCallID string `gorm:"-" json:"nextCallId"`

	// Camp 阵营（非持久化字段，数据库无此列）
	Camp string `gorm:"-" json:"camp"`

	// Next 剧本 ids（列名 next 为 SQL 保留字）
	Next string `gorm:"size:255;column:next" json:"next"`

	// Attrib 属性
	Attrib int32 `gorm:"column:attrib" json:"attrib"`

	// AddTime 创建时间
	AddTime time.Time `gorm:"autoCreateTime;column:add_time" json:"addTime"`

	// Deleted 逻辑删除标记（tinyint(1)）
	Deleted bool `gorm:"column:deleted" json:"deleted"`

	// UpdateTime 更新时间
	UpdateTime time.Time `gorm:"autoUpdateTime;column:update_time" json:"updateTime"`
}

// TableName 显式指定表名
func (NpcDialogueFrame) TableName() string {
	return "npc_dialogue_frame"
}
