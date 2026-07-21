package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type MapInfoDAO = BaseDAO[model.MapInfo]

func NewMapInfoDAO() *MapInfoDAO {
	return NewBaseDAO[model.MapInfo](
		db.GORM(),
		db.Cache(),
		"map",
		func(t *model.MapInfo) int64 {
			return int64(t.ID)
		},
	)
}
