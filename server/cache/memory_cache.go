package cache

import (
	"context"
	"sync"
	"time"
)

// MemoryCache 纯内存实现的 Cache。
//
// 主要用途：
//   - 单元测试：无需启动 Redis 即可验证 DAO 逻辑
//   - 本地开发降级：临时关闭 Redis 时启用
//   - 单机场景的轻量缓存
//
// 不持久化，进程退出即失效；并发安全。
type MemoryCache struct {
	mu   sync.RWMutex
	data map[string]memEntry
}

type memEntry struct {
	val      []byte
	expireAt time.Time // zero 表示永不过期
}

// NewMemoryCache 构造一个空内存缓存
func NewMemoryCache() *MemoryCache {
	return &MemoryCache{data: make(map[string]memEntry, 64)}
}

// 编译期接口检查：确保 *MemoryCache 实现了 Cache 接口
// 相当于 Java 中的: class MemoryCache implements Cache { ... }
var _ Cache = (*MemoryCache)(nil)

func (m *MemoryCache) GetRaw(ctx context.Context, key string) ([]byte, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	e, ok := m.data[key]
	if !ok {
		return nil, ErrNil
	}
	if !e.expireAt.IsZero() && time.Now().After(e.expireAt) {
		// 惰性过期：命中过期键视为不存在
		return nil, ErrNil
	}
	// 复制一份，避免外部修改污染内部
	out := make([]byte, len(e.val))
	copy(out, e.val)
	return out, nil
}

func (m *MemoryCache) SetRaw(ctx context.Context, key string, val []byte, ttl time.Duration) error {
	var expireAt time.Time
	if ttl > 0 {
		expireAt = time.Now().Add(ttl)
	}
	stored := make([]byte, len(val))
	copy(stored, val)

	m.mu.Lock()
	m.data[key] = memEntry{val: stored, expireAt: expireAt}
	m.mu.Unlock()
	return nil
}

func (m *MemoryCache) Del(ctx context.Context, keys ...string) error {
	if len(keys) == 0 {
		return nil
	}
	m.mu.Lock()
	for _, k := range keys {
		delete(m.data, k)
	}
	m.mu.Unlock()
	return nil
}

func (m *MemoryCache) Exists(ctx context.Context, key string) (bool, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	e, ok := m.data[key]
	if !ok {
		return false, nil
	}
	if !e.expireAt.IsZero() && time.Now().After(e.expireAt) {
		return false, nil
	}
	return true, nil
}

func (m *MemoryCache) Close() error { return nil }
