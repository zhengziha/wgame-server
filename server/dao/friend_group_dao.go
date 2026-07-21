package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type FriendGroupDAO = BaseDAO[model.FriendGroup]

func NewFriendGroupDAO() *FriendGroupDAO {
	return NewBaseDAO[model.FriendGroup](
		db.GORM(),
		db.Cache(),
		"friend_group",
		func(t *model.FriendGroup) int64 {
			return int64(t.ID)
		},
	)
}
