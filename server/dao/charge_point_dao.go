package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type ChargePointDAO = BaseDAO[model.ChargePoint]

func NewChargePointDAO() *ChargePointDAO {
	return NewBaseDAO[model.ChargePoint](
		db.GORM(),
		db.Cache(),
		"charge_point",
		func(t *model.ChargePoint) int64 {
			return int64(t.ID)
		},
	)
}
