package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type BlackListDAO = BaseDAO[model.BlackList]

func NewBlackListDAO() *BlackListDAO {
	return NewBaseDAO[model.BlackList](
		db.AuthGORM(),
		db.Cache(),
		"black_list",
		func(t *model.BlackList) int64 {
			return int64(t.ID)
		},
	)
}
