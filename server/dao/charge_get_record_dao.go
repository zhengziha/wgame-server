package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type ChargeGetRecordDAO = BaseDAO[model.ChargeGetRecord]

func NewChargeGetRecordDAO() *ChargeGetRecordDAO {
	return NewBaseDAO[model.ChargeGetRecord](
		db.GORM(),
		db.Cache(),
		"charge_get_record",
		func(t *model.ChargeGetRecord) int64 {
			return int64(t.ID)
		},
	)
}
