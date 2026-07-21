package db

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// newSQLiteDialector 构造 SQLite 的 gorm.Dialector。
// 使用 glebarez/sqlite（modernc.org/sqlite 纯 Go 实现，无需 CGO）。
//
// DSN 是文件路径或 ":memory:"；为空时默认 "data/game.db"。
// 若 DSN 指向的父目录不存在会自动创建（避免 "out of memory (14)" 这类错误）。
func newSQLiteDialector(cfg *DBConfig) gorm.Dialector {
	dsn := cfg.DSN
	if dsn == "" {
		dsn = "data/game.db"
	}
	ensureSQLiteDir(dsn)
	return sqlite.Open(dsn)
}

// ensureSQLiteDir 当 DSN 是文件路径时，确保其父目录存在。
// 跳过 ":memory:" 和带 query string 的内存 dsn。
func ensureSQLiteDir(dsn string) {
	if dsn == "" || dsn == ":memory:" {
		return
	}
	// 去掉可能存在的 query string（例如 "data/game.db?cache=shared"）
	path := dsn
	if idx := strings.IndexByte(path, '?'); idx >= 0 {
		path = path[:idx]
	}
	// 仅处理有目录前缀的路径
	dir := filepath.Dir(path)
	if dir == "" || dir == "." || dir == "/" {
		return
	}
	_ = os.MkdirAll(dir, 0o755)
}
