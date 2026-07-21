package model

import "time"

// CustomPetSkill 自定义宠物技能表。
type CustomPetSkill struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// PetName 宠物名称
	PetName string `gorm:"size:255;column:pet_name" json:"petName"`

	// SkillName 技能名称
	SkillName string `gorm:"size:255;column:skill_name" json:"skillName"`

	// SkillLevel 技能等级
	SkillLevel int32 `gorm:"default:1;column:skill_level" json:"skillLevel"`

	// SkillRange 技能范围
	SkillRange int32 `gorm:"default:1;column:skill_range" json:"skillRange"`

	// SkillRound 回合数
	SkillRound int32 `gorm:"default:1;column:skill_round" json:"skillRound"`

	// AddTime 创建时间
	AddTime time.Time `gorm:"autoCreateTime;column:add_time" json:"addTime"`
}

// TableName 显式指定表名
func (CustomPetSkill) TableName() string {
	return "custom_pet_skill"
}
