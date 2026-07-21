package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type LuckDrawTimesRewardDAO = BaseDAO[model.LuckDrawTimesReward]

func NewLuckDrawTimesRewardDAO() *LuckDrawTimesRewardDAO {
	return NewBaseDAO[model.LuckDrawTimesReward](
		db.GORM(),
		db.Cache(),
		"luck_draw_times_reward",
		func(t *model.LuckDrawTimesReward) int64 {
			return int64(t.ID)
		},
	)
}
