package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type ConnDAO = BaseDAO[model.Conn]

func NewConnDAO() *ConnDAO {
	return NewBaseDAO[model.Conn](
		db.AuthGORM(),
		db.Cache(),
		"conn",
		func(t *model.Conn) int64 {
			return int64(t.ID)
		},
	)
}
