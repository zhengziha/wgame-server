package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type PartyWarScoreDAO = BaseDAO[model.PartyWarScore]

func NewPartyWarScoreDAO() *PartyWarScoreDAO {
	return NewBaseDAO[model.PartyWarScore](
		db.GORM(),
		db.Cache(),
		"party_war_score",
		func(t *model.PartyWarScore) int64 {
			return int64(t.ID)
		},
	)
}
