package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type PartyWarBidDAO = BaseDAO[model.PartyWarBid]

func NewPartyWarBidDAO() *PartyWarBidDAO {
	return NewBaseDAO[model.PartyWarBid](
		db.GORM(),
		db.Cache(),
		"party_war_bid",
		func(t *model.PartyWarBid) int64 {
			return int64(t.ID)
		},
	)
}
