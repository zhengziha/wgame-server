package model

import "time"

// Notice 公告表。
// Time 列名为 SQL 保留字。
type Notice struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// Message 公告消息内容
	Message string `gorm:"size:1024;column:message" json:"message"`

	// Time 轮询时间（分钟）（列名 time 为 SQL 保留字）
	Time int32 `gorm:"column:time" json:"time"`

	// AddTime 创建时间
	AddTime time.Time `gorm:"autoCreateTime;column:add_time" json:"addTime"`

	// UpdateTime 更新时间
	UpdateTime time.Time `gorm:"autoUpdateTime;column:update_time" json:"updateTime"`

	// Deleted 逻辑删除标记（tinyint(1)）
	Deleted bool `gorm:"column:deleted" json:"deleted"`
}

// TableName 显式指定表名
func (Notice) TableName() string {
	return "notice"
}
