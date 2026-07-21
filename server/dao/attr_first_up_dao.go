package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type AttrFirstUpDAO = BaseDAO[model.AttrFirstUp]

func NewAttrFirstUpDAO() *AttrFirstUpDAO {
	return NewBaseDAO[model.AttrFirstUp](
		db.GORM(),
		db.Cache(),
		"attr_first_up",
		func(t *model.AttrFirstUp) int64 {
			return int64(t.ID)
		},
	)
}
