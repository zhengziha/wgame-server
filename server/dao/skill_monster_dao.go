package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type SkillMonsterDAO = BaseDAO[model.SkillMonster]

func NewSkillMonsterDAO() *SkillMonsterDAO {
	return NewBaseDAO[model.SkillMonster](
		db.GORM(),
		db.Cache(),
		"skill_monster",
		func(t *model.SkillMonster) int64 {
			return int64(t.ID)
		},
	)
}
