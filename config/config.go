// Package config 提供统一的配置管理，支持 YAML 文件 + 命令行参数覆盖。
//
// 优先级：命令行参数 > 环境变量 > 配置文件 > 默认值
//
// 使用方式：
//
//	cfg, err := config.Load("config.yml")
//	// 或通过命令行指定配置文件：
//	// ./wgame-server -config path/to/config.yml
package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

// Config 应用全局配置
type Config struct {
	Server ServerConfig `mapstructure:"server"`
	GameDB DBConfig     `mapstructure:"game_db"`
	AuthDB DBConfig     `mapstructure:"auth_db"`
	Cache  CacheConfig  `mapstructure:"cache"`
	Log    LogConfig    `mapstructure:"log"`
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Addr string `mapstructure:"addr"`
}

// DBConfig 数据库配置
type DBConfig struct {
	Driver          string `mapstructure:"driver"`
	DSN             string `mapstructure:"dsn"`
	LogLevel        int    `mapstructure:"log_level"`
	MaxOpenConns    int    `mapstructure:"max_open_conns"`
	MaxIdleConns    int    `mapstructure:"max_idle_conns"`
	ConnMaxLifetime int    `mapstructure:"conn_max_lifetime"`
}

// CacheConfig 缓存配置
type CacheConfig struct {
	Driver string      `mapstructure:"driver"`
	Redis  RedisConfig `mapstructure:"redis"`
}

// RedisConfig Redis 配置
type RedisConfig struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

// LogConfig 日志配置
type LogConfig struct {
	Level string `mapstructure:"level"`
}

// DefaultConfig 返回默认配置
func DefaultConfig() *Config {
	return &Config{
		Server: ServerConfig{
			Addr: ":8800",
		},
		GameDB: DBConfig{
			Driver:       "sqlite",
			DSN:          "data/game.db",
			LogLevel:     2,
			MaxOpenConns: 0,
			MaxIdleConns: 0,
		},
		AuthDB: DBConfig{
			Driver:       "sqlite",
			DSN:          "data/auth.db",
			LogLevel:     2,
			MaxOpenConns: 0,
			MaxIdleConns: 0,
		},
		Cache: CacheConfig{
			Driver: "redis",
			Redis: RedisConfig{
				Addr:     "127.0.0.1:6379",
				Password: "",
				DB:       0,
				PoolSize: 16,
			},
		},
		Log: LogConfig{
			Level: "info",
		},
	}
}

// Load 加载配置文件并应用命令行参数覆盖。
// configPath: 配置文件路径，为空时使用默认配置。
func Load(configPath string) (*Config, error) {
	cfg := DefaultConfig()

	v := viper.New()

	// 设置默认值
	setDefaults(v)

	// 读取配置文件
	if configPath != "" {
		v.SetConfigFile(configPath)
	} else {
		// 尝试从默认位置读取
		v.SetConfigName("config")
		v.SetConfigType("yml")
		v.AddConfigPath(".")
		v.AddConfigPath("./config")
		v.AddConfigPath("/etc/wgame-server")
	}

	// 尝试读取配置文件（如果不存在则忽略）
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("read config file: %w", err)
		}
		// 配置文件不存在，使用默认值
	}

	// 支持环境变量覆盖（前缀 WGAME）
	v.SetEnvPrefix("WGAME")
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// 解析配置到结构体
	if err := v.Unmarshal(cfg); err != nil {
		return nil, fmt.Errorf("unmarshal config: %w", err)
	}

	// 命令行参数覆盖（通过 pflag 在 main.go 中处理）
	// 这里提供 ApplyFlags 方法供 main.go 调用

	return cfg, nil
}

// setDefaults 设置默认配置值
func setDefaults(v *viper.Viper) {
	// Server
	v.SetDefault("server.addr", ":8800")

	// Game DB
	v.SetDefault("game_db.driver", "sqlite")
	v.SetDefault("game_db.dsn", "")
	v.SetDefault("game_db.log_level", 2)
	v.SetDefault("game_db.max_open_conns", 0)
	v.SetDefault("game_db.max_idle_conns", 0)

	// Auth DB
	v.SetDefault("auth_db.driver", "sqlite")
	v.SetDefault("auth_db.dsn", "")
	v.SetDefault("auth_db.log_level", 2)
	v.SetDefault("auth_db.max_open_conns", 0)
	v.SetDefault("auth_db.max_idle_conns", 0)

	// Cache
	v.SetDefault("cache.driver", "redis")
	v.SetDefault("cache.redis.addr", "127.0.0.1:6379")
	v.SetDefault("cache.redis.password", "")
	v.SetDefault("cache.redis.db", 0)
	v.SetDefault("cache.redis.pool_size", 16)

	// Log
	v.SetDefault("log.level", "info")
}

// ConfigFromEnv 从环境变量加载配置（用于容器化部署）
func ConfigFromEnv() *Config {
	cfg := DefaultConfig()

	if v := os.Getenv("WGAME_SERVER_ADDR"); v != "" {
		cfg.Server.Addr = v
	}
	if v := os.Getenv("WGAME_GAME_DB_DRIVER"); v != "" {
		cfg.GameDB.Driver = v
	}
	if v := os.Getenv("WGAME_GAME_DB_DSN"); v != "" {
		cfg.GameDB.DSN = v
	}
	if v := os.Getenv("WGAME_AUTH_DB_DRIVER"); v != "" {
		cfg.AuthDB.Driver = v
	}
	if v := os.Getenv("WGAME_AUTH_DB_DSN"); v != "" {
		cfg.AuthDB.DSN = v
	}
	if v := os.Getenv("WGAME_CACHE_DRIVER"); v != "" {
		cfg.Cache.Driver = v
	}
	if v := os.Getenv("WGAME_CACHE_REDIS_ADDR"); v != "" {
		cfg.Cache.Redis.Addr = v
	}
	if v := os.Getenv("WGAME_CACHE_REDIS_PASSWORD"); v != "" {
		cfg.Cache.Redis.Password = v
	}

	return cfg
}
