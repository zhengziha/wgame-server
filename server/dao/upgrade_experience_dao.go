package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type UpgradeExperienceDAO = BaseDAO[model.UpgradeExperience]

func NewUpgradeExperienceDAO() *UpgradeExperienceDAO {
	return NewBaseDAO[model.UpgradeExperience](
		db.GORM(),
		db.Cache(),
		"upgrade_experience",
		func(t *model.UpgradeExperience) int64 {
			return int64(t.Level)
		},
	)
}
