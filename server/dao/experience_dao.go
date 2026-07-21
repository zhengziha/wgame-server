package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type ExperienceDAO = BaseDAO[model.Experience]

func NewExperienceDAO() *ExperienceDAO {
	return NewBaseDAO[model.Experience](
		db.GORM(),
		db.Cache(),
		"experience",
		func(t *model.Experience) int64 {
			return int64(t.Attrib)
		},
	)
}
