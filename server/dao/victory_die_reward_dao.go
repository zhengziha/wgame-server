package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type VictoryDieRewardDAO = BaseDAO[model.VictoryDieReward]

func NewVictoryDieRewardDAO() *VictoryDieRewardDAO {
	return NewBaseDAO[model.VictoryDieReward](
		db.GORM(),
		db.Cache(),
		"victory_die_reward",
		func(t *model.VictoryDieReward) int64 {
			return int64(t.ID)
		},
	)
}
