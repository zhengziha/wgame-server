package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type AutoShuaguaiDAO = BaseDAO[model.AutoShuaguai]

func NewAutoShuaguaiDAO() *AutoShuaguaiDAO {
	return NewBaseDAO[model.AutoShuaguai](
		db.GORM(),
		db.Cache(),
		"auto_shuaguai",
		func(t *model.AutoShuaguai) int64 {
			return int64(t.ID)
		},
	)
}
