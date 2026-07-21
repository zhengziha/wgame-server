package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type ShidaoHistoryteamDAO = BaseDAO[model.ShidaoHistoryteam]

func NewShidaoHistoryteamDAO() *ShidaoHistoryteamDAO {
	return NewBaseDAO[model.ShidaoHistoryteam](
		db.GORM(),
		db.Cache(),
		"shidao_history_team",
		func(t *model.ShidaoHistoryteam) int64 {
			return int64(t.ID)
		},
	)
}
