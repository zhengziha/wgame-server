package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type JieBaiDAO = BaseDAO[model.JieBai]

func NewJieBaiDAO() *JieBaiDAO {
	return NewBaseDAO[model.JieBai](
		db.GORM(),
		db.Cache(),
		"jie_bai",
		func(t *model.JieBai) int64 {
			return int64(t.ID)
		},
	)
}
