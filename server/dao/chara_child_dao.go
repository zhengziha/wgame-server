package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type CharaChildDAO = BaseDAO[model.CharaChild]

func NewCharaChildDAO() *CharaChildDAO {
	return NewBaseDAO[model.CharaChild](
		db.GORM(),
		db.Cache(),
		"chara_child",
		func(t *model.CharaChild) int64 {
			return int64(t.ID)
		},
	)
}
