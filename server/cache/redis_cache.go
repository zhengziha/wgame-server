package cache

import (
	"context"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
)

// RedisCache 基于 go-redis 的 Cache 实现。
//
// 配置字段参考 RedisConfig；为方便 main 注入，结构公开，
// 调用方可以直接 &RedisCache{Cli: client} 自定义。
type RedisCache struct {
	Cli *redis.Client
}

// NewRedisCache 从 *redis.Client 构造一个 RedisCache。
// 传入 nil 也允许，运行期调用会返回 ErrNilCache。
func NewRedisCache(cli *redis.Client) *RedisCache {
	return &RedisCache{Cli: cli}
}

// Compile-time check
var _ Cache = (*RedisCache)(nil)

func (c *RedisCache) GetRaw(ctx context.Context, key string) ([]byte, error) {
	if c.Cli == nil {
		return nil, ErrNilCache
	}
	cmd := c.Cli.Get(ctx, key)
	if err := cmd.Err(); err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, ErrNil
		}
		return nil, err
	}
	return []byte(cmd.Val()), nil
}

func (c *RedisCache) SetRaw(ctx context.Context, key string, val []byte, ttl time.Duration) error {
	if c.Cli == nil {
		return ErrNilCache
	}
	return c.Cli.Set(ctx, key, val, ttl).Err()
}

func (c *RedisCache) Del(ctx context.Context, keys ...string) error {
	if c.Cli == nil {
		return ErrNilCache
	}
	if len(keys) == 0 {
		return nil
	}
	return c.Cli.Del(ctx, keys...).Err()
}

func (c *RedisCache) Exists(ctx context.Context, key string) (bool, error) {
	if c.Cli == nil {
		return false, ErrNilCache
	}
	n, err := c.Cli.Exists(ctx, key).Result()
	return n > 0, err
}

func (c *RedisCache) Close() error {
	if c.Cli == nil {
		return nil
	}
	return c.Cli.Close()
}
