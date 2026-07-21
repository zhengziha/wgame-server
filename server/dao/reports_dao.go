package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type ReportsDAO = BaseDAO[model.Reports]

func NewReportsDAO() *ReportsDAO {
	return NewBaseDAO[model.Reports](
		db.GORM(),
		db.Cache(),
		"reports",
		func(t *model.Reports) int64 {
			return int64(t.ID)
		},
	)
}
