package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type GoldStallMyBidGoodsDAO = BaseDAO[model.GoldStallMyBidGoods]

func NewGoldStallMyBidGoodsDAO() *GoldStallMyBidGoodsDAO {
	return NewBaseDAO[model.GoldStallMyBidGoods](
		db.GORM(),
		db.Cache(),
		"gold_stall_my_bid_goods",
		func(t *model.GoldStallMyBidGoods) int64 {
			return int64(t.ID)
		},
	)
}
