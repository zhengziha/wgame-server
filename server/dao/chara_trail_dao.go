package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type CharaTrailDAO = BaseDAO[model.CharaTrail]

func NewCharaTrailDAO() *CharaTrailDAO {
	return NewBaseDAO[model.CharaTrail](
		db.GORM(),
		db.Cache(),
		"chara_trail",
		func(t *model.CharaTrail) int64 {
			return int64(t.ID)
		},
	)
}
