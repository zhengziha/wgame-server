package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type RenwuDAO = BaseDAO[model.Renwu]

func NewRenwuDAO() *RenwuDAO {
	return NewBaseDAO[model.Renwu](
		db.GORM(),
		db.Cache(),
		"renwu",
		func(t *model.Renwu) int64 {
			return int64(t.ID)
		},
	)
}
