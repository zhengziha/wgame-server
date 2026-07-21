package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type ArenaRankDAO = BaseDAO[model.ArenaRank]

func NewArenaRankDAO() *ArenaRankDAO {
	return NewBaseDAO[model.ArenaRank](
		db.GORM(),
		db.Cache(),
		"arena_rank",
		func(t *model.ArenaRank) int64 {
			return int64(t.ID)
		},
	)
}
