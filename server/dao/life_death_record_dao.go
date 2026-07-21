package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type LifeDeathRecordDAO = BaseDAO[model.LifeDeathRecord]

func NewLifeDeathRecordDAO() *LifeDeathRecordDAO {
	return NewBaseDAO[model.LifeDeathRecord](
		db.GORM(),
		db.Cache(),
		"life_death_record",
		func(t *model.LifeDeathRecord) int64 {
			return int64(t.ID)
		},
	)
}
