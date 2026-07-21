package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type GiftPackConfigDAO = BaseDAO[model.GiftPackConfig]

func NewGiftPackConfigDAO() *GiftPackConfigDAO {
	return NewBaseDAO[model.GiftPackConfig](
		db.GORM(),
		db.Cache(),
		"gift_pack_config",
		func(t *model.GiftPackConfig) int64 {
			return int64(t.ID)
		},
	)
}
