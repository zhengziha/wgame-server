package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type FriendDAO = BaseDAO[model.Friend]

func NewFriendDAO() *FriendDAO {
	return NewBaseDAO[model.Friend](
		db.GORM(),
		db.Cache(),
		"friend",
		func(t *model.Friend) int64 {
			return int64(t.ID)
		},
	)
}
