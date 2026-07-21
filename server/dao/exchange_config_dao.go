package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type ExchangeConfigDAO = BaseDAO[model.ExchangeConfig]

func NewExchangeConfigDAO() *ExchangeConfigDAO {
	return NewBaseDAO[model.ExchangeConfig](
		db.GORM(),
		db.Cache(),
		"exchange_config",
		func(t *model.ExchangeConfig) int64 {
			return int64(t.ID)
		},
	)
}
