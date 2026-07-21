package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type DaySignPrizeDAO = BaseDAO[model.DaySignPrize]

func NewDaySignPrizeDAO() *DaySignPrizeDAO {
	return NewBaseDAO[model.DaySignPrize](
		db.GORM(),
		db.Cache(),
		"day_sign_prize",
		func(t *model.DaySignPrize) int64 {
			return int64(t.ID)
		},
	)
}
