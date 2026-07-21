package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type GrowthTaskDAO = BaseDAO[model.GrowthTask]

func NewGrowthTaskDAO() *GrowthTaskDAO {
	return NewBaseDAO[model.GrowthTask](
		db.GORM(),
		db.Cache(),
		"growth_task",
		func(t *model.GrowthTask) int64 {
			return int64(t.ID)
		},
	)
}
