// Package main 是 wgame-server 的入口。
//
// 启动流程：
//  1. 根据 -db-driver 选择数据库实现（sqlite / mysql / postgres）
//  2. 根据 -cache-driver 选择缓存实现（redis / memory / none）
//  3. AutoMigrate 业务模型
//  4. 触发 demo handler 包的 init() 自注册
//  5. 启动 TCP socket server
//
// 整个进程只暴露一个 TCP 监听端口，使用自定义 10 字节头协议。
package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"wgame-server/comm/log"
	"wgame-server/server/cache"
	"wgame-server/server/db"
	"wgame-server/server/model"
	"wgame-server/server/network/socket"

	// 匿名导入以触发 init() 自注册
	_ "wgame-server/server/demo/handlers"
)

func main() {
	addr := flag.String("addr", ":8800", "TCP listen address")

	// 数据库层：通过 -db-driver 切换 sqlite / mysql / postgres
	dbDriver := flag.String("db-driver", "sqlite", "database driver: sqlite | mysql | postgres")
	dbDSN := flag.String("db-dsn", "", "database DSN (empty=driver default)")
	dbLogLevel := flag.Int("db-log", 2, "GORM log level: 1 Silent, 2 Error, 3 Warn, 4 Info")
	dbMaxOpen := flag.Int("db-max-open", 0, "max open conns (0=driver default)")
	dbMaxIdle := flag.Int("db-max-idle", 0, "max idle conns (0=driver default)")

	// 缓存层：通过 -cache-driver 切换 redis / memory / none
	cacheDriver := flag.String("cache-driver", "redis", "cache driver: redis | memory | none")
	redisAddr := flag.String("redis-addr", "127.0.0.1:6379", "Redis address (only used when cache-driver=redis)")
	redisPassword := flag.String("redis-pass", "", "Redis password")
	redisDB := flag.Int("redis-db", 0, "Redis logical db")
	redisPool := flag.Int("redis-pool", 16, "Redis pool size")
	flag.Parse()

	// 1) 初始化数据库
	dbCfg := &db.DBConfig{
		Driver:       *dbDriver,
		DSN:          *dbDSN,
		LogLevel:     *dbLogLevel,
		MaxOpenConns: *dbMaxOpen,
		MaxIdleConns: *dbMaxIdle,
	}
	if err := db.Init(dbCfg, nil); err != nil {
		log.Error("[main] db init failed: %v", err)
		os.Exit(1)
	}
	defer db.Close()
	log.Info("[main] db initialized: driver=%s", *dbDriver)

	// 2) 构造 Cache 实现（与 DB 初始化解耦）
	cacheImpl, err := buildCache(*cacheDriver, *redisAddr, *redisPassword, *redisDB, *redisPool)
	if err != nil {
		log.Error("[main] cache build failed: %v", err)
		os.Exit(1)
	}
	if cacheImpl != nil {
		db.SetCache(cacheImpl)
		log.Info("[main] cache initialized: driver=%s", *cacheDriver)
	} else {
		log.Info("[main] cache disabled (driver=none)")
	}

	// 3) AutoMigrate 业务模型
	if err := db.AutoMigrate(&model.User{}); err != nil {
		log.Error("[main] automigrate failed: %v", err)
		os.Exit(1)
	}
	log.Info("[main] automigrate done")

	// 4) 启动 socket server
	srv := socket.NewServer(socket.Config{Addr: *addr})
	srv.OnConnect(func(_ *socket.SocketCmdContext) {})
	srv.OnDisconnect(func(c *socket.SocketCmdContext) {
		log.Info("[main] disconnect sid=%d uid=%d", c.GetSessionId(), c.GetUserId())
	})

	go func() {
		if err := srv.Start(); err != nil {
			log.Error("[main] server start failed: %v", err)
			os.Exit(1)
		}
	}()

	// 5) 阻塞等待信号
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	log.Info("[main] recv signal %v, shutting down...", <-sig)
	srv.Stop()
	db.Close()
	log.Info("[main] bye")
}

// buildCache 根据 driver 名称构造 cache.Cache 实现。
// 返回 (nil, nil) 表示禁用缓存。
func buildCache(driver, redisAddr, redisPassword string, redisDB, redisPool int) (cache.Cache, error) {
	switch driver {
	case "", "none":
		return nil, nil
	case "memory":
		return cache.NewMemoryCache(), nil
	case "redis":
		cli := db.NewRedisClient(&db.RedisConfig{
			Addr:     redisAddr,
			Password: redisPassword,
			DB:       redisDB,
			PoolSize: redisPool,
		})
		return cache.NewRedisCache(cli), nil
	default:
		return nil, fmt.Errorf("unknown cache driver: %s", driver)
	}
}
