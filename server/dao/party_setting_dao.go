package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type PartySettingDAO = BaseDAO[model.PartySetting]

func NewPartySettingDAO() *PartySettingDAO {
	return NewBaseDAO[model.PartySetting](
		db.GORM(),
		db.Cache(),
		"party_setting",
		func(t *model.PartySetting) int64 {
			return int64(t.ID)
		},
	)
}
