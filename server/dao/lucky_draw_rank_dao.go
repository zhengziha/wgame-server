package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type LuckyDrawRankDAO = BaseDAO[model.LuckyDrawRank]

func NewLuckyDrawRankDAO() *LuckyDrawRankDAO {
	return NewBaseDAO[model.LuckyDrawRank](
		db.GORM(),
		db.Cache(),
		"lucky_draw_rank",
		func(t *model.LuckyDrawRank) int64 {
			return int64(t.ID)
		},
	)
}
