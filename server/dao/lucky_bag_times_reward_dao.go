package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type LuckyBagTimesRewardDAO = BaseDAO[model.LuckyBagTimesReward]

func NewLuckyBagTimesRewardDAO() *LuckyBagTimesRewardDAO {
	return NewBaseDAO[model.LuckyBagTimesReward](
		db.GORM(),
		db.Cache(),
		"lucky_bag_times_reward",
		func(t *model.LuckyBagTimesReward) int64 {
			return int64(t.ID)
		},
	)
}
