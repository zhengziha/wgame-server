package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type NpcDAO = BaseDAO[model.Npc]

func NewNpcDAO() *NpcDAO {
	return NewBaseDAO[model.Npc](
		db.GORM(),
		db.Cache(),
		"npc",
		func(t *model.Npc) int64 {
			return int64(t.ID)
		},
	)
}
