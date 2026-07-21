package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type NpcDialogueDAO = BaseDAO[model.NpcDialogue]

func NewNpcDialogueDAO() *NpcDialogueDAO {
	return NewBaseDAO[model.NpcDialogue](
		db.GORM(),
		db.Cache(),
		"npc_dialogue",
		func(t *model.NpcDialogue) int64 {
			return int64(t.ID)
		},
	)
}
