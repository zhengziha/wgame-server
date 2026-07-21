package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type LuckDrawItemDAO = BaseDAO[model.LuckDrawItem]

func NewLuckDrawItemDAO() *LuckDrawItemDAO {
	return NewBaseDAO[model.LuckDrawItem](
		db.GORM(),
		db.Cache(),
		"luck_draw_item",
		func(t *model.LuckDrawItem) int64 {
			return int64(t.ID)
		},
	)
}
