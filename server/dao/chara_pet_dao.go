package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type CharaPetDAO = BaseDAO[model.CharaPet]

func NewCharaPetDAO() *CharaPetDAO {
	return NewBaseDAO[model.CharaPet](
		db.GORM(),
		db.Cache(),
		"chara_pet",
		func(t *model.CharaPet) int64 {
			return int64(t.ID)
		},
	)
}
