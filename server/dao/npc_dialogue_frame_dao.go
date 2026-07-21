package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type NpcDialogueFrameDAO = BaseDAO[model.NpcDialogueFrame]

func NewNpcDialogueFrameDAO() *NpcDialogueFrameDAO {
	return NewBaseDAO[model.NpcDialogueFrame](
		db.GORM(),
		db.Cache(),
		"npc_dialogue_frame",
		func(t *model.NpcDialogueFrame) int64 {
			return int64(t.ID)
		},
	)
}
