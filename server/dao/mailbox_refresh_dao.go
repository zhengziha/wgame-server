package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type MailboxRefreshDAO = BaseDAO[model.MailboxRefresh]

func NewMailboxRefreshDAO() *MailboxRefreshDAO {
	return NewBaseDAO[model.MailboxRefresh](
		db.GORM(),
		db.Cache(),
		"mailbox_refresh",
		func(t *model.MailboxRefresh) int64 {
			return int64(t.ID)
		},
	)
}
