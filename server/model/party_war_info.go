package model

import "time"

// PartyWarInfo 帮派战赛程信息表。
// group、time 为 SQL 保留字，column tag 直接写列名，GORM 会自动加反引号。
// Java 实体字段使用 snake_case 命名（comp_area/lunkong_result），Go 字段转为驼峰。
type PartyWarInfo struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// No 第几届
	No int32 `gorm:"column:no" json:"no"`

	// CompArea 赛区 1-4
	CompArea string `gorm:"size:255;column:comp_area" json:"compArea"`

	// Stage 赛程阶段 1小组赛 2第一场半决赛 3第二场半决赛 4三四名半决赛 5决赛
	Stage int32 `gorm:"column:stage" json:"stage"`

	// Group 组 A/B，空为淘汰赛（列名 group 为 SQL 保留字）
	Group string `gorm:"size:255;column:group" json:"group"`

	// Result 结果 attacker_win/defenser_win/lunkong
	Result string `gorm:"size:255;column:result" json:"result"`

	// Time 结束时间（列名 time 为 SQL 保留字）
	Time time.Time `gorm:"column:time" json:"time"`

	// LunkongResult 轮空结果 attacker_win/defenser_win
	LunkongResult string `gorm:"size:255;column:lunkong_result" json:"lunkongResult"`

	// Attacker 进攻方
	Attacker string `gorm:"size:255;column:attacker" json:"attacker"`

	// Defenser 防守方
	Defenser string `gorm:"size:255;column:defenser" json:"defenser"`
}

// TableName 显式指定表名
func (PartyWarInfo) TableName() string {
	return "party_war_info"
}
