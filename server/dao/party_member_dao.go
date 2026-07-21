package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type PartyMemberDAO = BaseDAO[model.PartyMember]

func NewPartyMemberDAO() *PartyMemberDAO {
	return NewBaseDAO[model.PartyMember](
		db.GORM(),
		db.Cache(),
		"party_member",
		func(t *model.PartyMember) int64 {
			return int64(t.ID)
		},
	)
}
