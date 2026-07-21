package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type ZhuangbeiInfoDAO = BaseDAO[model.ZhuangbeiInfo]

func NewZhuangbeiInfoDAO() *ZhuangbeiInfoDAO {
	return NewBaseDAO[model.ZhuangbeiInfo](
		db.GORM(),
		db.Cache(),
		"zhuangbei_info",
		func(t *model.ZhuangbeiInfo) int64 {
			return int64(t.ID)
		},
	)
}
