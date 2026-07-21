package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type ChargeConfigDAO = BaseDAO[model.ChargeConfig]

func NewChargeConfigDAO() *ChargeConfigDAO {
	return NewBaseDAO[model.ChargeConfig](
		db.GORM(),
		db.Cache(),
		"charge_config",
		func(t *model.ChargeConfig) int64 {
			return int64(t.ID)
		},
	)
}
