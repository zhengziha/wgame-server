package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type LivenessRewardsDAO = BaseDAO[model.LivenessRewards]

func NewLivenessRewardsDAO() *LivenessRewardsDAO {
	return NewBaseDAO[model.LivenessRewards](
		db.GORM(),
		db.Cache(),
		"liveness_rewards",
		func(t *model.LivenessRewards) int64 {
			return int64(t.ID)
		},
	)
}
