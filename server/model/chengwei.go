package model

import "time"

// Chengwei 称谓表。
// time 为 SQL 保留字，column tag 直接写列名，GORM 会自动加反引号。
// openUpLineNotice 在 Java 实体中声明，wd-game-18.sql 建表语句未出现该列，按普通列保留以兼容 Java 逻辑。
type Chengwei struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// Name 称谓名称（唯一索引）
	Name string `gorm:"size:155;uniqueIndex;column:name" json:"name"`

	// Money 累计充值金额
	Money int32 `gorm:"column:money" json:"money"`

	// Time 过期时间（列名 time 为 SQL 保留字）
	Time int32 `gorm:"column:time" json:"time"`

	// Color 文字颜色
	Color string `gorm:"size:255;column:color" json:"color"`

	// AddTime 创建时间
	AddTime time.Time `gorm:"autoCreateTime;column:add_time" json:"addTime"`

	// Attr 属性 JSON
	Attr string `gorm:"type:text;column:attr" json:"attr"`

	// Icon 图标
	Icon int32 `gorm:"column:icon" json:"icon"`

	// OpenUpLineNotice 是否开启上线全服公告（0关闭 1开启）
	OpenUpLineNotice int32 `gorm:"column:open_up_line_notice" json:"openUpLineNotice"`
}

// TableName 显式指定表名
func (Chengwei) TableName() string {
	return "chengwei"
}
