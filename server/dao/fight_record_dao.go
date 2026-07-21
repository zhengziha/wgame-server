package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type FightRecordDAO = BaseDAO[model.FightRecord]

func NewFightRecordDAO() *FightRecordDAO {
	return NewBaseDAO[model.FightRecord](
		db.GORM(),
		db.Cache(),
		"fight_record",
		func(t *model.FightRecord) int64 {
			return int64(t.ID)
		},
	)
}
