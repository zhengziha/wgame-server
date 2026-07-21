package model

import "time"

// CharaStatue 人物雕像表（对应 Java Chara_Statue）。
// 注意：SQL 中 serverId 列为驼峰命名，npc_name 上有唯一索引。
type CharaStatue struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// Serverid 服务器 id（SQL 列名为驼峰 serverId）
	Serverid string `gorm:"size:128;column:serverId" json:"serverid"`

	// NpcName npc 名字（SQL 列有唯一索引）
	NpcName string `gorm:"size:128;uniqueIndex;column:npc_name" json:"npcName"`

	// AddTime 创建时间
	AddTime time.Time `gorm:"autoCreateTime;column:add_time" json:"addTime"`

	// UpdateTime 更新时间
	UpdateTime time.Time `gorm:"autoUpdateTime;column:update_time" json:"updateTime"`

	// Deleted 逻辑删除标记（tinyint(1)）
	Deleted bool `gorm:"column:deleted" json:"deleted"`

	// Data 雕像数据（mediumtext）
	Data string `gorm:"type:text;column:data" json:"data"`
}

// TableName 显式指定表名
func (CharaStatue) TableName() string {
	return "chara_statue"
}
