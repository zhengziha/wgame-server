package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type ChengweiDAO = BaseDAO[model.Chengwei]

func NewChengweiDAO() *ChengweiDAO {
	return NewBaseDAO[model.Chengwei](
		db.GORM(),
		db.Cache(),
		"chengwei",
		func(t *model.Chengwei) int64 {
			return int64(t.ID)
		},
	)
}
