package db

import "gorm.io/gorm"

// NewDialector 根据 DBConfig.Driver 创建对应的 gorm.Dialector。
// 这是 driver 的统一工厂入口；具体实现分别在
// dialect_sqlite.go / dialect_mysql.go / dialect_postgres.go。
//
// 未识别的 driver 返回 ErrUnsupportedDriver。
func NewDialector(cfg *DBConfig) (gorm.Dialector, error) {
	switch cfg.Driver {
	case DriverSQLite:
		return newSQLiteDialector(cfg), nil
	case DriverMySQL:
		return newMySQLDialector(cfg), nil
	case DriverPostgres:
		return newPostgresDialector(cfg), nil
	default:
		return nil, ErrUnsupportedDriver
	}
}
