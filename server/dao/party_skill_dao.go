package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type PartySkillDAO = BaseDAO[model.PartySkill]

func NewPartySkillDAO() *PartySkillDAO {
	return NewBaseDAO[model.PartySkill](
		db.GORM(),
		db.Cache(),
		"party_skill",
		func(t *model.PartySkill) int64 {
			return int64(t.ID)
		},
	)
}
