// Package cache 定义与具体后端无关的缓存抽象。
//
// 设计要点：
//   - Cache 是接口；RedisCache / MemoryCache 是两种实现，可按需替换
//   - 业务层（DAO）依赖接口而非具体实现，便于单测注入 MemoryCache
//   - JSON 序列化在抽象层完成，实现只负责原始字节存取
//   - 所有 key 由调用方决定前缀，避免不同业务冲突
package cache

import (
	"context"
	"encoding/json"
	"errors"
	"time"
)

// ErrNil key 不存在（等价 redigo.ErrNil / redis.Nil）
var ErrNil = errors.New("cache: key not found")

// DefaultTTL 默认缓存有效期，可被 Set 的参数覆盖
const DefaultTTL = 10 * time.Minute

// Cache 与具体后端无关的缓存接口。
//
// 最小化契约：
//   - GetRaw: 取原始字节，未命中返回 ErrNil
//   - SetRaw: 存原始字节
//   - Del:    删除若干 key
//   - Exists: 判断 key 是否存在
//   - Close:  释放底层资源（无则 no-op）
//
// 上层 Get/Set 包装 JSON 序列化后调用底层方法。
type Cache interface {
	GetRaw(ctx context.Context, key string) ([]byte, error)
	SetRaw(ctx context.Context, key string, val []byte, ttl time.Duration) error
	Del(ctx context.Context, keys ...string) error
	Exists(ctx context.Context, key string) (bool, error)
	Close() error
}

// Set 把任意可 JSON 序列化的对象写入缓存。
// ttl <= 0 时使用 DefaultTTL。
func Set(ctx context.Context, c Cache, key string, val interface{}, ttl time.Duration) error {
	if c == nil {
		return ErrNilCache
	}
	if ttl <= 0 {
		ttl = DefaultTTL
	}
	buf, err := json.Marshal(val)
	if err != nil {
		return err
	}
	return c.SetRaw(ctx, key, buf, ttl)
}

// Get 从缓存取回 JSON 并 Unmarshal 到 out（out 必须是指针）。
// key 不存在时返回 ErrNil。
func Get(ctx context.Context, c Cache, key string, out interface{}) error {
	if c == nil {
		return ErrNilCache
	}
	buf, err := c.GetRaw(ctx, key)
	if err != nil {
		return err
	}
	return json.Unmarshal(buf, out)
}

// Del 删除一个或多个 key
func Del(ctx context.Context, c Cache, keys ...string) error {
	if c == nil {
		return ErrNilCache
	}
	if len(keys) == 0 {
		return nil
	}
	return c.Del(ctx, keys...)
}

// Exists 判断 key 是否存在
func Exists(ctx context.Context, c Cache, key string) (bool, error) {
	if c == nil {
		return false, ErrNilCache
	}
	return c.Exists(ctx, key)
}

// ErrNilCache 表示未注入任何 Cache 实现
var ErrNilCache = errors.New("cache: no cache implementation injected")
