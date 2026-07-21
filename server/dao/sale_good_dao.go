package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type SaleGoodDAO = BaseDAO[model.SaleGood]

func NewSaleGoodDAO() *SaleGoodDAO {
	return NewBaseDAO[model.SaleGood](
		db.GORM(),
		db.Cache(),
		"sale_good",
		func(t *model.SaleGood) int64 {
			return int64(t.ID)
		},
	)
}
