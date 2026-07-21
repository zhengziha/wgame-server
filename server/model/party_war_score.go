package model

// PartyWarScore 帮派战积分表。
// group 为 SQL 保留字，column tag 直接写列名，GORM 会自动加反引号。
// Java 实体字段使用 snake_case 命名（party_name/war_score），Go 字段转为驼峰。
// outlet 在 Java 中为 String，SQL 中为 int，按 Java 类型映射为 string。
type PartyWarScore struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// No 第几届
	No int32 `gorm:"column:no" json:"no"`

	// PartyName 帮派名字
	PartyName string `gorm:"size:255;column:party_name" json:"partyName"`

	// Group 组 A/B，空为淘汰赛（列名 group 为 SQL 保留字）
	Group string `gorm:"size:255;column:group" json:"group"`

	// Outlet 出线 1出线 0不出线
	Outlet string `gorm:"column:outlet" json:"outlet"`

	// WarScore 积分
	WarScore int32 `gorm:"column:war_score" json:"warScore"`
}

// TableName 显式指定表名
func (PartyWarScore) TableName() string {
	return "party_war_score"
}
