package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type AnchorRandomBoxItemDAO = BaseDAO[model.AnchorRandomBoxItem]

func NewAnchorRandomBoxItemDAO() *AnchorRandomBoxItemDAO {
	return NewBaseDAO[model.AnchorRandomBoxItem](
		db.GORM(),
		db.Cache(),
		"anchor_random_box_item",
		func(t *model.AnchorRandomBoxItem) int64 {
			return int64(t.ID)
		},
	)
}
