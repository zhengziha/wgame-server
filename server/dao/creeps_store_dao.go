package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type CreepsStoreDAO = BaseDAO[model.CreepsStore]

func NewCreepsStoreDAO() *CreepsStoreDAO {
	return NewBaseDAO[model.CreepsStore](
		db.GORM(),
		db.Cache(),
		"creeps_store",
		func(t *model.CreepsStore) int64 {
			return int64(t.ID)
		},
	)
}
