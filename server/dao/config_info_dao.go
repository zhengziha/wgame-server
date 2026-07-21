package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type ConfigInfoDAO = BaseDAO[model.ConfigInfo]

func NewConfigInfoDAO() *ConfigInfoDAO {
	return NewBaseDAO[model.ConfigInfo](
		db.GORM(),
		db.Cache(),
		"config_info",
		func(t *model.ConfigInfo) int64 {
			return int64(t.ID)
		},
	)
}
