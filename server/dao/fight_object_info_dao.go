package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type FightObjectInfoDAO = BaseDAO[model.FightObjectInfo]

func NewFightObjectInfoDAO() *FightObjectInfoDAO {
	return NewBaseDAO[model.FightObjectInfo](
		db.GORM(),
		db.Cache(),
		"fight_object_info",
		func(t *model.FightObjectInfo) int64 {
			return int64(t.ID)
		},
	)
}
