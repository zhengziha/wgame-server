// Package main 是 wgame-server 的入口。
//
// 启动流程：
//  1. 加载配置文件（支持命令行参数覆盖）
//  2. 根据配置初始化数据库（game + auth 双数据源）
//  3. 根据配置初始化缓存
//  4. AutoMigrate 业务模型
//  5. 触发 demo handler 包的 init() 自注册
//  6. 启动 TCP socket server
//
// 配置优先级：命令行参数 > 环境变量 > 配置文件 > 默认值
//
// 使用示例：
//
//	# 使用默认配置
//	./wgame-server
//
//	# 指定配置文件
//	./wgame-server -config config.yml
//
//	# 配置文件 + 命令行覆盖
//	./wgame-server -config config.yml -addr :9000 -db-driver mysql
package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"wgame-server/comm/log"
	"wgame-server/config"
	"wgame-server/server/cache"
	"wgame-server/server/core"
	"wgame-server/server/db"
	"wgame-server/server/network/socket"

	// 匿名导入以触发 init() 自注册
	_ "wgame-server/server/demo/handlers"
	_ "wgame-server/server/handler/auth"
	_ "wgame-server/server/handler/map"
	_ "wgame-server/server/handler/system"
)

func main() {
	// 1) 命令行参数定义
	configPath := flag.String("config", "", "path to config file (yaml)")

	// 以下参数可用于覆盖配置文件中的对应值
	addr := flag.String("addr", "", "TCP listen address (override config)")
	dbDriver := flag.String("db-driver", "", "game database driver: sqlite | mysql | postgres")
	dbDSN := flag.String("db-dsn", "", "game database DSN")
	dbLogLevel := flag.Int("db-log", 0, "game DB GORM log level (0=use config)")
	dbMaxOpen := flag.Int("db-max-open", 0, "game DB max open conns (0=use config)")
	dbMaxIdle := flag.Int("db-max-idle", 0, "game DB max idle conns (0=use config)")
	cacheDriver := flag.String("cache-driver", "", "cache driver: redis | memory | none")
	redisAddr := flag.String("redis-addr", "", "Redis address")
	redisPassword := flag.String("redis-pass", "", "Redis password")
	redisDB := flag.Int("redis-db", -1, "Redis logical db (-1=use config)")
	redisPool := flag.Int("redis-pool", 0, "Redis pool size (0=use config)")
	flag.Parse()

	// 2) 加载配置文件
	cfg, err := config.Load(*configPath)
	if err != nil {
		log.Error("[main] load config failed: %v", err)
		os.Exit(1)
	}

	// 3) 命令行参数覆盖配置文件
	applyFlags(cfg, addr, dbDriver, dbDSN, dbLogLevel, dbMaxOpen, dbMaxIdle,
		cacheDriver, redisAddr, redisPassword, redisDB, redisPool)

	log.Info("[main] config loaded: server.addr=%s, game_db.driver=%s, cache.driver=%s",
		cfg.Server.Addr, cfg.GameDB.Driver, cfg.Cache.Driver)

	// 4) 初始化游戏数据库
	if err := db.Init(&db.DBConfig{
		Driver:       cfg.GameDB.Driver,
		DSN:          cfg.GameDB.DSN,
		LogLevel:     cfg.GameDB.LogLevel,
		MaxOpenConns: cfg.GameDB.MaxOpenConns,
		MaxIdleConns: cfg.GameDB.MaxIdleConns,
	}, nil); err != nil {
		log.Error("[main] game db init failed: %v", err)
		os.Exit(1)
	}
	defer db.Close()
	log.Info("[main] game db initialized: driver=%s", cfg.GameDB.Driver)

	// 5) 初始化认证数据库（如果配置了）
	if cfg.AuthDB.Driver != "" && cfg.AuthDB.DSN != "" {
		if err := db.InitAuth(&db.DBConfig{
			Driver:       cfg.AuthDB.Driver,
			DSN:          cfg.AuthDB.DSN,
			LogLevel:     cfg.AuthDB.LogLevel,
			MaxOpenConns: cfg.AuthDB.MaxOpenConns,
			MaxIdleConns: cfg.AuthDB.MaxIdleConns,
		}); err != nil {
			log.Error("[main] auth db init failed: %v", err)
			os.Exit(1)
		}
		log.Info("[main] auth db initialized: driver=%s", cfg.AuthDB.Driver)
	}

	// 6) 初始化缓存
	cacheImpl, err := buildCache(cfg.Cache)
	if err != nil {
		log.Error("[main] cache build failed: %v", err)
		os.Exit(1)
	}
	if cacheImpl != nil {
		db.SetCache(cacheImpl)
		log.Info("[main] cache initialized: driver=%s", cfg.Cache.Driver)
	} else {
		log.Info("[main] cache disabled (driver=none)")
	}

	// 7) AutoMigrate 所有业务模型
	// if err := db.AutoMigrate(model.AllModels()...); err != nil {
	// 	log.Error("[main] automigrate failed: %v", err)
	// 	os.Exit(1)
	// }
	// log.Info("[main] automigrate done: %d models", len(model.AllModels()))

	// 8) 初始化 GameCore（线路和地图）
	core.Instance().SetLineNum(1) // 默认 1 条线路，可配置
	if err := core.Instance().InitLineAndMap(db.GORM()); err != nil {
		log.Error("[main] game core init failed: %v", err)
		os.Exit(1)
	}
	core.Instance().SetReady(true)
	log.Info("[main] game core initialized")

	// 9) 启动 socket server
	srv := socket.NewServer(socket.Config{Addr: cfg.Server.Addr})
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

	// 10) 阻塞等待信号
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	log.Info("[main] recv signal %v, shutting down...", <-sig)
	srv.Stop()
	db.Close()
	log.Info("[main] bye")
}

// applyFlags 将命令行参数应用到配置（仅覆盖非零值）
func applyFlags(cfg *config.Config,
	addr, dbDriver, dbDSN *string,
	dbLogLevel, dbMaxOpen, dbMaxIdle *int,
	cacheDriver, redisAddr, redisPassword *string,
	redisDB, redisPool *int) {

	if *addr != "" {
		cfg.Server.Addr = *addr
	}
	if *dbDriver != "" {
		cfg.GameDB.Driver = *dbDriver
	}
	if *dbDSN != "" {
		cfg.GameDB.DSN = *dbDSN
	}
	if *dbLogLevel > 0 {
		cfg.GameDB.LogLevel = *dbLogLevel
	}
	if *dbMaxOpen > 0 {
		cfg.GameDB.MaxOpenConns = *dbMaxOpen
	}
	if *dbMaxIdle > 0 {
		cfg.GameDB.MaxIdleConns = *dbMaxIdle
	}
	if *cacheDriver != "" {
		cfg.Cache.Driver = *cacheDriver
	}
	if *redisAddr != "" {
		cfg.Cache.Redis.Addr = *redisAddr
	}
	if *redisPassword != "" {
		cfg.Cache.Redis.Password = *redisPassword
	}
	if *redisDB >= 0 {
		cfg.Cache.Redis.DB = *redisDB
	}
	if *redisPool > 0 {
		cfg.Cache.Redis.PoolSize = *redisPool
	}
}

// buildCache 根据配置构造 cache.Cache 实现。
// 返回 (nil, nil) 表示禁用缓存。
func buildCache(cfg config.CacheConfig) (cache.Cache, error) {
	switch cfg.Driver {
	case "", "none":
		return nil, nil
	case "memory":
		return cache.NewMemoryCache(), nil
	case "redis":
		cli := db.NewRedisClient(&db.RedisConfig{
			Addr:     cfg.Redis.Addr,
			Password: cfg.Redis.Password,
			DB:       cfg.Redis.DB,
			PoolSize: cfg.Redis.PoolSize,
		})
		return cache.NewRedisCache(cli), nil
	default:
		return nil, fmt.Errorf("unknown cache driver: %s", cfg.Driver)
	}
}
