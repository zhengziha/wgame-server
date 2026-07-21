package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type TaskFightInfoDAO = BaseDAO[model.TaskFightInfo]

func NewTaskFightInfoDAO() *TaskFightInfoDAO {
	return NewBaseDAO[model.TaskFightInfo](
		db.GORM(),
		db.Cache(),
		"task_fight_info",
		func(t *model.TaskFightInfo) int64 {
			return int64(t.ID)
		},
	)
}
