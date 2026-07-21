package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type FasionDAO = BaseDAO[model.Fasion]

func NewFasionDAO() *FasionDAO {
	return NewBaseDAO[model.Fasion](
		db.GORM(),
		db.Cache(),
		"pack_modification",
		func(t *model.Fasion) int64 {
			return int64(t.ID)
		},
	)
}
