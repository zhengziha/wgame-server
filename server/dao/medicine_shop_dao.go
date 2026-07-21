package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type MedicineShopDAO = BaseDAO[model.MedicineShop]

func NewMedicineShopDAO() *MedicineShopDAO {
	return NewBaseDAO[model.MedicineShop](
		db.GORM(),
		db.Cache(),
		"medicine_shop",
		func(t *model.MedicineShop) int64 {
			return int64(t.ID)
		},
	)
}
