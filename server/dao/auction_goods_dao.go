package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type AuctionGoodsDAO = BaseDAO[model.AuctionGoods]

func NewAuctionGoodsDAO() *AuctionGoodsDAO {
	return NewBaseDAO[model.AuctionGoods](
		db.GORM(),
		db.Cache(),
		"auction_goods",
		func(t *model.AuctionGoods) int64 {
			return int64(t.ID)
		},
	)
}
