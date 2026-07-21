package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type PetHelpTypeDAO = BaseDAO[model.PetHelpType]

func NewPetHelpTypeDAO() *PetHelpTypeDAO {
	return NewBaseDAO[model.PetHelpType](
		db.GORM(),
		db.Cache(),
		"pet_help_type",
		func(t *model.PetHelpType) int64 {
			return int64(t.ID)
		},
	)
}
