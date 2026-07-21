package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type CharaStatueDAO = BaseDAO[model.CharaStatue]

func NewCharaStatueDAO() *CharaStatueDAO {
	return NewBaseDAO[model.CharaStatue](
		db.GORM(),
		db.Cache(),
		"chara_statue",
		func(t *model.CharaStatue) int64 {
			return int64(t.ID)
		},
	)
}
