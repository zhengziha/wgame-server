package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type HouseFurnitureDAO = BaseDAO[model.HouseFurniture]

func NewHouseFurnitureDAO() *HouseFurnitureDAO {
	return NewBaseDAO[model.HouseFurniture](
		db.GORM(),
		db.Cache(),
		"house_furniture",
		func(t *model.HouseFurniture) int64 {
			return int64(t.ID)
		},
	)
}
