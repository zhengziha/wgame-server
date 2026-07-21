package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type StoreInfoDAO = BaseDAO[model.StoreInfo]

func NewStoreInfoDAO() *StoreInfoDAO {
	return NewBaseDAO[model.StoreInfo](
		db.GORM(),
		db.Cache(),
		"store_info",
		func(t *model.StoreInfo) int64 {
			return int64(t.ID)
		},
	)
}
