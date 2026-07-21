package model

// PartySkill 帮派技能表。
// level 为 SQL 保留字，column tag 直接写列名，GORM 会自动加反引号。
type PartySkill struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// PartyId 帮派 id
	PartyId string `gorm:"size:255;column:party_id" json:"partyId"`

	// No 技能编号
	No int32 `gorm:"column:no" json:"no"`

	// Name 技能名称
	Name string `gorm:"size:255;column:name" json:"name"`

	// Level 等级（列名 level 为 SQL 保留字）
	Level int32 `gorm:"column:level" json:"level"`

	// CurrentScore 当前分
	CurrentScore int32 `gorm:"column:current_score" json:"currentScore"`

	// LevelupScore 下一级分
	LevelupScore int32 `gorm:"column:levelup_score" json:"levelupScore"`
}

// TableName 显式指定表名
func (PartySkill) TableName() string {
	return "party_skill"
}
