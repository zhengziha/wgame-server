package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type DialogDAO = BaseDAO[model.Dialog]

func NewDialogDAO() *DialogDAO {
	return NewBaseDAO[model.Dialog](
		db.GORM(),
		db.Cache(),
		"dialog",
		func(t *model.Dialog) int64 {
			return int64(t.ID)
		},
	)
}
