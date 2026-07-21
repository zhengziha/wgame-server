package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type CustomPetSkillDAO = BaseDAO[model.CustomPetSkill]

func NewCustomPetSkillDAO() *CustomPetSkillDAO {
	return NewBaseDAO[model.CustomPetSkill](
		db.GORM(),
		db.Cache(),
		"custom_pet_skill",
		func(t *model.CustomPetSkill) int64 {
			return int64(t.ID)
		},
	)
}
