package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type GrowthConfigDAO = BaseDAO[model.GrowthConfig]

func NewGrowthConfigDAO() *GrowthConfigDAO {
	return NewBaseDAO[model.GrowthConfig](
		db.GORM(),
		db.Cache(),
		"growth_config",
		func(t *model.GrowthConfig) int64 {
			return int64(t.ID)
		},
	)
}
