package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type FasionCustomInfoDAO = BaseDAO[model.FasionCustomInfo]

func NewFasionCustomInfoDAO() *FasionCustomInfoDAO {
	return NewBaseDAO[model.FasionCustomInfo](
		db.GORM(),
		db.Cache(),
		"fasion_custom_info",
		func(t *model.FasionCustomInfo) int64 {
			return int64(t.ID)
		},
	)
}
