package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type GoodShopDAO = BaseDAO[model.GoodShop]

func NewGoodShopDAO() *GoodShopDAO {
	return NewBaseDAO[model.GoodShop](
		db.GORM(),
		db.Cache(),
		"good_shop",
		func(t *model.GoodShop) int64 {
			return int64(t.ID)
		},
	)
}
