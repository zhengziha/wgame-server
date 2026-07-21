package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type SysUserDAO = BaseDAO[model.SysUser]

func NewSysUserDAO() *SysUserDAO {
	return NewBaseDAO[model.SysUser](
		db.GORM(),
		db.Cache(),
		"sys_user",
		func(t *model.SysUser) int64 {
			return t.ID
		},
	)
}
