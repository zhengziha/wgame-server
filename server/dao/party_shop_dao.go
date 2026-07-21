package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type PartyShopDAO = BaseDAO[model.PartyShop]

func NewPartyShopDAO() *PartyShopDAO {
	return NewBaseDAO[model.PartyShop](
		db.GORM(),
		db.Cache(),
		"party_shop",
		func(t *model.PartyShop) int64 {
			return int64(t.ID)
		},
	)
}
