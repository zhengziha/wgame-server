package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type DailiDAO = BaseDAO[model.Daili]

func NewDailiDAO() *DailiDAO {
	return NewBaseDAO[model.Daili](
		db.AuthGORM(),
		db.Cache(),
		"daili",
		func(t *model.Daili) int64 {
			return int64(t.ID)
		},
	)
}
