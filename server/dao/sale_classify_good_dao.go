package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type SaleClassifyGoodDAO = BaseDAO[model.SaleClassifyGood]

func NewSaleClassifyGoodDAO() *SaleClassifyGoodDAO {
	return NewBaseDAO[model.SaleClassifyGood](
		db.GORM(),
		db.Cache(),
		"sale_classify_good",
		func(t *model.SaleClassifyGood) int64 {
			return int64(t.ID)
		},
	)
}
