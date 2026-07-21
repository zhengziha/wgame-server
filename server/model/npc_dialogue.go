package model

import "time"

// NpcDialogue NPC 对话表。
// 注意：GroupID/NextCallType/NextCallID 为非持久化字段，数据库中不存在对应列。
type NpcDialogue struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// Name 说话者名称
	Name string `gorm:"size:255;column:name" json:"name"`

	// Portranit 头像编号
	Portranit int32 `gorm:"column:portranit" json:"portranit"`

	// PicNo 图片编号
	PicNo int32 `gorm:"column:pic_no" json:"picNo"`

	// Content 对话内容
	Content string `gorm:"size:255;column:content" json:"content"`

	// Isconmlete 是否完成
	Isconmlete int32 `gorm:"column:isconmlete" json:"isconmlete"`

	// Isincombat 是否在战斗中
	Isincombat int32 `gorm:"column:isincombat" json:"isincombat"`

	// Palytime 播放时间
	Palytime int32 `gorm:"column:palytime" json:"palytime"`

	// TaskType 任务类型
	TaskType string `gorm:"size:255;column:task_type" json:"taskType"`

	// GroupID 分组 ID（非持久化字段，数据库无此列）
	GroupID string `gorm:"-" json:"groupId"`

	// NextCallType 下一步调用类型（非持久化字段，数据库无此列）
	NextCallType string `gorm:"-" json:"nextCallType"`

	// NextCallID 下一步调用 ID（非持久化字段，数据库无此列）
	NextCallID string `gorm:"-" json:"nextCallId"`

	// AddTime 创建时间
	AddTime time.Time `gorm:"autoCreateTime;column:add_time" json:"addTime"`

	// UpdateTime 更新时间
	UpdateTime time.Time `gorm:"autoUpdateTime;column:update_time" json:"updateTime"`

	// Deleted 逻辑删除标记（tinyint(1)）
	Deleted bool `gorm:"column:deleted" json:"deleted"`

	// Idname 关联 ID 名称
	Idname string `gorm:"size:255;column:idname" json:"idname"`
}

// TableName 显式指定表名
func (NpcDialogue) TableName() string {
	return "npc_dialogue"
}
