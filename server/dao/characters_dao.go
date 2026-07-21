package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type CharactersDAO = BaseDAO[model.Characters]

func NewCharactersDAO() *CharactersDAO {
	return NewBaseDAO[model.Characters](
		db.GORM(),
		db.Cache(),
		"characters",
		func(t *model.Characters) int64 {
			return int64(t.ID)
		},
	)
}
