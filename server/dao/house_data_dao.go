package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type HouseDataDAO = BaseDAO[model.HouseData]

func NewHouseDataDAO() *HouseDataDAO {
	return NewBaseDAO[model.HouseData](
		db.GORM(),
		db.Cache(),
		"house_data",
		func(t *model.HouseData) int64 {
			return int64(t.ID)
		},
	)
}
