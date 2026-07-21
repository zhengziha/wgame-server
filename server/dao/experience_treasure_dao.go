package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type ExperienceTreasureDAO = BaseDAO[model.ExperienceTreasure]

func NewExperienceTreasureDAO() *ExperienceTreasureDAO {
	return NewBaseDAO[model.ExperienceTreasure](
		db.GORM(),
		db.Cache(),
		"experience_treasure",
		func(t *model.ExperienceTreasure) int64 {
			return int64(t.Attrib)
		},
	)
}
