package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type ShidaoHistoryDAO = BaseDAO[model.ShidaoHistory]

func NewShidaoHistoryDAO() *ShidaoHistoryDAO {
	return NewBaseDAO[model.ShidaoHistory](
		db.GORM(),
		db.Cache(),
		"shidao_history",
		func(t *model.ShidaoHistory) int64 {
			return int64(t.ID)
		},
	)
}
