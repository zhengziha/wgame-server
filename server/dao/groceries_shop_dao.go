package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type GroceriesShopDAO = BaseDAO[model.GroceriesShop]

func NewGroceriesShopDAO() *GroceriesShopDAO {
	return NewBaseDAO[model.GroceriesShop](
		db.GORM(),
		db.Cache(),
		"groceries_shop",
		func(t *model.GroceriesShop) int64 {
			return int64(t.ID)
		},
	)
}
