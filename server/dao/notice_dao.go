package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type NoticeDAO = BaseDAO[model.Notice]

func NewNoticeDAO() *NoticeDAO {
	return NewBaseDAO[model.Notice](
		db.GORM(),
		db.Cache(),
		"notice",
		func(t *model.Notice) int64 {
			return int64(t.ID)
		},
	)
}
