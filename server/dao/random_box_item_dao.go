package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type RandomBoxItemDAO = BaseDAO[model.RandomBoxItem]

func NewRandomBoxItemDAO() *RandomBoxItemDAO {
	return NewBaseDAO[model.RandomBoxItem](
		db.GORM(),
		db.Cache(),
		"random_box_item",
		func(t *model.RandomBoxItem) int64 {
			return int64(t.ID)
		},
	)
}
