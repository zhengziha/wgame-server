package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type PetDAO = BaseDAO[model.Pet]

func NewPetDAO() *PetDAO {
	return NewBaseDAO[model.Pet](
		db.GORM(),
		db.Cache(),
		"pet",
		func(t *model.Pet) int64 {
			return int64(t.ID)
		},
	)
}
