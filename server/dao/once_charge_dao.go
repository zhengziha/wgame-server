package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type OnceChargeDAO = BaseDAO[model.OnceCharge]

func NewOnceChargeDAO() *OnceChargeDAO {
	return NewBaseDAO[model.OnceCharge](
		db.GORM(),
		db.Cache(),
		"one_charge",
		func(t *model.OnceCharge) int64 {
			return int64(t.ID)
		},
	)
}
