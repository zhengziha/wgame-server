package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type WeddingListDAO = BaseDAO[model.WeddingList]

func NewWeddingListDAO() *WeddingListDAO {
	return NewBaseDAO[model.WeddingList](
		db.GORM(),
		db.Cache(),
		"wedding_list",
		func(t *model.WeddingList) int64 {
			return int64(t.ID)
		},
	)
}
