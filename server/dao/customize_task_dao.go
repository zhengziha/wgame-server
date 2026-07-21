package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type CustomizeTaskDAO = BaseDAO[model.CustomizeTask]

func NewCustomizeTaskDAO() *CustomizeTaskDAO {
	return NewBaseDAO[model.CustomizeTask](
		db.GORM(),
		db.Cache(),
		"customize_task",
		func(t *model.CustomizeTask) int64 {
			return int64(t.ID)
		},
	)
}
