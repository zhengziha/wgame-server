package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type StallRecordDAO = BaseDAO[model.StallRecord]

func NewStallRecordDAO() *StallRecordDAO {
	return NewBaseDAO[model.StallRecord](
		db.GORM(),
		db.Cache(),
		"stall_record",
		func(t *model.StallRecord) int64 {
			return int64(t.ID)
		},
	)
}
