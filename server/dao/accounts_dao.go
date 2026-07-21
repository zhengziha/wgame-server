package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type AccountsDAO = BaseDAO[model.Accounts]

func NewAccountsDAO() *AccountsDAO {
	return NewBaseDAO[model.Accounts](
		db.AuthGORM(),
		db.Cache(),
		"accounts",
		func(t *model.Accounts) int64 {
			return int64(t.ID)
		},
	)
}
