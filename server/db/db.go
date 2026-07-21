// Package db 提供数据库 (GORM) 与 Cache 的统一初始化与全局访问入口。
//
// 设计要点：
//   - 数据库通过 GORM 的 Dialector 抽象，支持 sqlite / mysql / postgres / 自定义
//   - 缓存通过 cache.Cache 接口抽象，支持 redis / memory / 自定义
//   - 全局单例 + Init/Close 生命周期，避免业务层到处传 *gorm.DB
//   - AutoMigrate 单独暴露，由 main 在初始化后按需调用
//
// 解耦策略：
//   - DAO 只依赖 cache.Cache 接口和 db.GORM()，不依赖具体实现
//   - 切换数据库/缓存后端只需在 main 中传入不同 driver + DSN
package db

import (
	"errors"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"wgame-server/server/cache"
)

// 支持的数据库 driver 常量（用于 NewDialector）
const (
	DriverSQLite   = "sqlite"
	DriverMySQL    = "mysql"
	DriverPostgres = "postgres"
)

// DBConfig 数据库统一配置。
//
// 不同 driver 的差异集中在 DSN 格式上：
//   - sqlite:   "data/game.db" 或 ":memory:"
//   - mysql:    "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
//   - postgres: "host=127.0.0.1 user=postgres password=pass dbname=game port=5432 sslmode=disable TimeZone=Asia/Shanghai"
type DBConfig struct {
	// Driver 数据库类型：sqlite | mysql | postgres
	// （自定义 driver 需直接调用 InitWithDialector）
	Driver string

	// DSN 数据源名称，格式随 driver 变化
	DSN string

	// LogLevel GORM 日志级别：1 Silent, 2 Error, 3 Warn, 4 Info
	LogLevel int

	// MaxOpenConns 连接池最大连接数（<=0 时按 driver 默认值）
	MaxOpenConns int

	// MaxIdleConns 连接池最大空闲连接数
	MaxIdleConns int

	// ConnMaxLifetime 连接最长生命周期
	ConnMaxLifetime time.Duration
}

// RedisConfig Redis 配置。
type RedisConfig struct {
	Addr     string
	Password string
	DB       int
	PoolSize int
}

// 全局实例
var (
	gormDB  *gorm.DB
	cacheDB cache.Cache

	rw        sync.RWMutex
	initOnce  sync.Once
	initErr   error
	closeOnce sync.Once
)

// Init 同时初始化数据库和 Cache。
//
// 参数：
//   - dbCfg   数据库配置；传 nil 跳过数据库层
//   - c       已构造好的 Cache 实现；传 nil 跳过缓存层
//
// 任一失败都会返回 error，并清理已建立的资源。
//
// 使用示例：
//
//	// SQLite + Redis
//	db.Init(&db.DBConfig{Driver: "sqlite", DSN: "data/game.db"},
//	    cache.NewRedisCache(redis.NewClient(...)))
//
//	// MySQL + Memory
//	db.Init(&db.DBConfig{Driver: "mysql", DSN: "user:pass@tcp(...)/db"},
//	    cache.NewMemoryCache())
func Init(dbCfg *DBConfig, c cache.Cache) error {
	initOnce.Do(func() {
		if dbCfg != nil {
			dl, err := NewDialector(dbCfg)
			if err != nil {
				initErr = err
				return
			}
			initErr = initGORM(dbCfg, dl)
			if initErr != nil {
				return
			}
		}
		if c != nil {
			cacheDB = c
		}
	})
	return initErr
}

// InitWithDialector 使用自定义 gorm.Dialector 初始化数据库。
// 用于支持尚未封装到 NewDialector 的数据库（如 SQL Server、ClickHouse）。
func InitWithDialector(dbCfg *DBConfig, dl gorm.Dialector, c cache.Cache) error {
	initOnce.Do(func() {
		if dbCfg != nil && dl != nil {
			initErr = initGORM(dbCfg, dl)
			if initErr != nil {
				return
			}
		}
		if c != nil {
			cacheDB = c
		}
	})
	return initErr
}

// initGORM 用给定 Dialector 打开 GORM 并配置连接池
func initGORM(cfg *DBConfig, dl gorm.Dialector) error {
	logLevel := logger.Info
	if cfg.LogLevel > 0 && cfg.LogLevel <= 4 {
		logLevel = logger.LogLevel(cfg.LogLevel)
	}
	gdb, err := gorm.Open(dl, &gorm.Config{
		Logger:                 logger.Default.LogMode(logLevel),
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		return err
	}
	sqlDB, err := gdb.DB()
	if err != nil {
		return err
	}

	// 根据 driver 给出合理的连接池默认值
	maxOpen, maxIdle, maxLifetime := defaultPoolSize(cfg.Driver)
	if cfg.MaxOpenConns > 0 {
		maxOpen = cfg.MaxOpenConns
	}
	if cfg.MaxIdleConns > 0 {
		maxIdle = cfg.MaxIdleConns
	}
	if cfg.ConnMaxLifetime > 0 {
		maxLifetime = cfg.ConnMaxLifetime
	}
	sqlDB.SetMaxOpenConns(maxOpen)
	sqlDB.SetMaxIdleConns(maxIdle)
	sqlDB.SetConnMaxLifetime(maxLifetime)

	gormDB = gdb
	return nil
}

// defaultPoolSize 不同 driver 的默认连接池
func defaultPoolSize(driver string) (maxOpen, maxIdle int, maxLifetime time.Duration) {
	switch driver {
	case DriverSQLite:
		// SQLite 单写多读，连接池设小一些避免写锁竞争
		return 5, 2, time.Hour
	default:
		// MySQL / Postgres 等网络数据库可放宽
		return 32, 8, time.Hour
	}
}

// NewRedisClient 工厂方法：基于 RedisConfig 创建 go-redis 客户端。
func NewRedisClient(cfg *RedisConfig) *redis.Client {
	poolSize := cfg.PoolSize
	if poolSize <= 0 {
		poolSize = 16
	}
	return redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
		PoolSize: poolSize,
	})
}

// GORM 返回全局 *gorm.DB；未初始化时返回 nil。
func GORM() *gorm.DB {
	rw.RLock()
	defer rw.RUnlock()
	return gormDB
}

// Cache 返回全局 Cache 实现接口；未注入时返回 nil。
func Cache() cache.Cache {
	rw.RLock()
	defer rw.RUnlock()
	return cacheDB
}

// SetCache 在 Init 之后替换缓存实现。
func SetCache(c cache.Cache) {
	rw.Lock()
	cacheDB = c
	rw.Unlock()
}

// Close 关闭所有资源（幂等）。通常在进程退出时调用。
func Close() error {
	closeOnce.Do(func() {
		if gormDB != nil {
			if sqlDB, err := gormDB.DB(); err == nil {
				_ = sqlDB.Close()
			}
			gormDB = nil
		}
		if cacheDB != nil {
			_ = cacheDB.Close()
			cacheDB = nil
		}
	})
	return nil
}

// AutoMigrate 自动建表。传入需要迁移的模型指针。
func AutoMigrate(models ...interface{}) error {
	if gormDB == nil {
		return ErrNotInitialized
	}
	return gormDB.AutoMigrate(models...)
}

// ErrNotInitialized 表示未调用 Init 成功就访问了 DB
var ErrNotInitialized = errors.New("db: not initialized, call db.Init first")

// ErrUnsupportedDriver 表示 NewDialector 收到未识别的 driver
var ErrUnsupportedDriver = errors.New("db: unsupported driver")
