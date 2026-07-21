package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type PartyDAO = BaseDAO[model.Party]

func NewPartyDAO() *PartyDAO {
	return NewBaseDAO[model.Party](
		db.GORM(),
		db.Cache(),
		"party",
		func(t *model.Party) int64 {
			return int64(t.ID)
		},
	)
}
