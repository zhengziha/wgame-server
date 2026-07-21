package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type LuckyDrawRecordDAO = BaseDAO[model.LuckyDrawRecord]

func NewLuckyDrawRecordDAO() *LuckyDrawRecordDAO {
	return NewBaseDAO[model.LuckyDrawRecord](
		db.GORM(),
		db.Cache(),
		"lucky_draw_record",
		func(t *model.LuckyDrawRecord) int64 {
			return int64(t.ID)
		},
	)
}
