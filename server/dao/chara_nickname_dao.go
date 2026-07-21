package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type CharaNicknameDAO = BaseDAO[model.CharaNickname]

func NewCharaNicknameDAO() *CharaNicknameDAO {
	return NewBaseDAO[model.CharaNickname](
		db.GORM(),
		db.Cache(),
		"chara_nickname",
		func(t *model.CharaNickname) int64 {
			return t.ID
		},
	)
}
