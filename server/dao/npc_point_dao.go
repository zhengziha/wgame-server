package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type NpcPointDAO = BaseDAO[model.NpcPoint]

func NewNpcPointDAO() *NpcPointDAO {
	return NewBaseDAO[model.NpcPoint](
		db.GORM(),
		db.Cache(),
		"npc_point",
		func(t *model.NpcPoint) int64 {
			return int64(t.ID)
		},
	)
}
