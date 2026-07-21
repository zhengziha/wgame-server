package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// newPostgresDialector 构造 PostgreSQL 的 gorm.Dialector。
//
// DSN 示例：
//
//	host=127.0.0.1 user=postgres password=pass dbname=game port=5432 sslmode=disable TimeZone=Asia/Shanghai
//
// 为空时会用本地默认值，便于本地开发快速启动。
func newPostgresDialector(cfg *DBConfig) gorm.Dialector {
	dsn := cfg.DSN
	if dsn == "" {
		dsn = "host=127.0.0.1 user=postgres password=postgres dbname=game port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	}
	return postgres.Open(dsn)
}
