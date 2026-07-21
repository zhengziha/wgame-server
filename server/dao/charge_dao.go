package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type ChargeDAO = BaseDAO[model.Charge]

func NewChargeDAO() *ChargeDAO {
	return NewBaseDAO[model.Charge](
		db.AuthGORM(),
		db.Cache(),
		"charge",
		func(t *model.Charge) int64 {
			return int64(t.ID)
		},
	)
}
