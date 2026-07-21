package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type RenwuMonsterDAO = BaseDAO[model.RenwuMonster]

func NewRenwuMonsterDAO() *RenwuMonsterDAO {
	return NewBaseDAO[model.RenwuMonster](
		db.GORM(),
		db.Cache(),
		"renwu_monster",
		func(t *model.RenwuMonster) int64 {
			return int64(t.ID)
		},
	)
}
