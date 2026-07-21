package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type FirstKillDAO = BaseDAO[model.FirstKill]

func NewFirstKillDAO() *FirstKillDAO {
	return NewBaseDAO[model.FirstKill](
		db.GORM(),
		db.Cache(),
		"first_kill",
		func(t *model.FirstKill) int64 {
			return int64(t.ID)
		},
	)
}
