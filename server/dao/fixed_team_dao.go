package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type FixedTeamDAO = BaseDAO[model.FixedTeam]

func NewFixedTeamDAO() *FixedTeamDAO {
	return NewBaseDAO[model.FixedTeam](
		db.GORM(),
		db.Cache(),
		"fixed_team",
		func(t *model.FixedTeam) int64 {
			return int64(t.ID)
		},
	)
}
