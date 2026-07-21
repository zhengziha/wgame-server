package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type ShengwangShopDAO = BaseDAO[model.ShengwangShop]

func NewShengwangShopDAO() *ShengwangShopDAO {
	return NewBaseDAO[model.ShengwangShop](
		db.GORM(),
		db.Cache(),
		"t_shengwang_shop",
		func(t *model.ShengwangShop) int64 {
			return int64(t.ID)
		},
	)
}
