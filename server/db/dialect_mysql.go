package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// newMySQLDialector 构造 MySQL 的 gorm.Dialector。
//
// DSN 示例：
//
//	user:pass@tcp(127.0.0.1:3306)/game?charset=utf8mb4&parseTime=True&loc=Local
//
// 为空时会用本地默认值，便于本地开发快速启动。
func newMySQLDialector(cfg *DBConfig) gorm.Dialector {
	dsn := cfg.DSN
	if dsn == "" {
		dsn = "root:root@tcp(127.0.0.1:3306)/game?charset=utf8mb4&parseTime=True&loc=Local"
	}
	return mysql.Open(dsn)
}
