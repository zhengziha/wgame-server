package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type MysticalRechargeDAO = BaseDAO[model.MysticalRecharge]

func NewMysticalRechargeDAO() *MysticalRechargeDAO {
	return NewBaseDAO[model.MysticalRecharge](
		db.GORM(),
		db.Cache(),
		"mystical_recharge",
		func(t *model.MysticalRecharge) int64 {
			return int64(t.ID)
		},
	)
}
