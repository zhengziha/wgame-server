package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type AccessibilityMapDAO = BaseDAO[model.AccessibilityMap]

func NewAccessibilityMapDAO() *AccessibilityMapDAO {
	return NewBaseDAO[model.AccessibilityMap](
		db.GORM(),
		db.Cache(),
		"accessibility_map",
		func(t *model.AccessibilityMap) int64 {
			return int64(t.ID)
		},
	)
}
