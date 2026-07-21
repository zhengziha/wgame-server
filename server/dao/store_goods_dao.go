package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type StoreGoodsDAO = BaseDAO[model.StoreGoods]

func NewStoreGoodsDAO() *StoreGoodsDAO {
	return NewBaseDAO[model.StoreGoods](
		db.GORM(),
		db.Cache(),
		"store_goods",
		func(t *model.StoreGoods) int64 {
			return int64(t.ID)
		},
	)
}
