package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type LuckyBagRandomBoxItemDAO = BaseDAO[model.LuckyBagRandomBoxItem]

func NewLuckyBagRandomBoxItemDAO() *LuckyBagRandomBoxItemDAO {
	return NewBaseDAO[model.LuckyBagRandomBoxItem](
		db.GORM(),
		db.Cache(),
		"lucky_bag_random_box_item",
		func(t *model.LuckyBagRandomBoxItem) int64 {
			return int64(t.ID)
		},
	)
}
