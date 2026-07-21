package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type ChargeRankDAO = BaseDAO[model.ChargeRank]

func NewChargeRankDAO() *ChargeRankDAO {
	return NewBaseDAO[model.ChargeRank](
		db.GORM(),
		db.Cache(),
		"charge_rank",
		func(t *model.ChargeRank) int64 {
			return int64(t.ID)
		},
	)
}
