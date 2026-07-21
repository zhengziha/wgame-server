package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type ChoujiangDAO = BaseDAO[model.Choujiang]

func NewChoujiangDAO() *ChoujiangDAO {
	return NewBaseDAO[model.Choujiang](
		db.GORM(),
		db.Cache(),
		"choujiang",
		func(t *model.Choujiang) int64 {
			return int64(t.ID)
		},
	)
}
