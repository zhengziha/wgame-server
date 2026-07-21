package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type GoldStallNineGoodsDAO = BaseDAO[model.GoldStallNineGoods]

func NewGoldStallNineGoodsDAO() *GoldStallNineGoodsDAO {
	return NewBaseDAO[model.GoldStallNineGoods](
		db.GORM(),
		db.Cache(),
		"gold_stall_nine_goods",
		func(t *model.GoldStallNineGoods) int64 {
			return int64(t.ID)
		},
	)
}
