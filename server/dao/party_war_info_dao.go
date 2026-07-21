package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type PartyWarInfoDAO = BaseDAO[model.PartyWarInfo]

func NewPartyWarInfoDAO() *PartyWarInfoDAO {
	return NewBaseDAO[model.PartyWarInfo](
		db.GORM(),
		db.Cache(),
		"party_war_info",
		func(t *model.PartyWarInfo) int64 {
			return int64(t.ID)
		},
	)
}
