// Package dao 数据访问层：组合 GORM（持久化）+ cache.Cache（缓存）。
//
// 缓存策略：Cache-Aside（旁路缓存）
//   - 读：先查 Cache，命中则返回；未命中查 DB，回写 Cache
//   - 写：先写 DB，再删除 Cache（避免并发更新导致脏数据）
//   - Key 命名：wgame:{table}:{pk}，例如 wgame:user:123
//
// 解耦：
//   - DAO 通过 db.Cache() 拿到 cache.Cache 接口实例，
//     任何实现该接口的缓存（Redis/Memory/自定义）都可以无缝替换。
//   - 单测可注入 cache.NewMemoryCache()，无需启动 Redis。
package dao

import (
	"context"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"

	"wgame-server/server/cache"
	"wgame-server/server/db"
)

// BaseDAO 泛型数据访问对象，提供通用 CRUD + Cache-Aside 实现。
// T 为任意 GORM 模型类型。
type BaseDAO[T any] struct {
	gormDB      *gorm.DB
	cache       cache.Cache
	tableName   string
	pkExtractor func(*T) int64
}

// NewBaseDAO 创建一个 BaseDAO 实例。
// 参数：
//   - gormDB: GORM 数据库实例
//   - cache:  缓存接口实例（可为 nil，禁用缓存）
//   - tableName: 表名（用于生成缓存 key）
//   - pkExtractor: 主键提取函数，从模型中提取 int64 主键
func NewBaseDAO[T any](
	gormDB *gorm.DB,
	cache cache.Cache,
	tableName string,
	pkExtractor func(*T) int64,
) *BaseDAO[T] {
	return &BaseDAO[T]{
		gormDB:      gormDB,
		cache:       cache,
		tableName:   tableName,
		pkExtractor: pkExtractor,
	}
}

// cacheKey 生成缓存 key
func (d *BaseDAO[T]) cacheKey(id int64) string {
	return fmt.Sprintf("wgame:%s:%d", d.tableName, id)
}

// GetByID 按主键查询，先走 Cache 再回源 DB（Cache-Aside 模式）。
// 未命中或缓存出错时降级到 DB 查询。
func (d *BaseDAO[T]) GetByID(ctx context.Context, id int64) (*T, error) {
	// 1) 查缓存
	key := d.cacheKey(id)
	var result T
	if d.cache != nil {
		if err := cache.Get(ctx, d.cache, key, &result); err == nil {
			return &result, nil
		}
		// 缓存未命中或出错，降级到 DB
	}

	// 2) 查 DB
	if d.gormDB == nil {
		return nil, db.ErrNotInitialized
	}
	if err := d.gormDB.WithContext(ctx).First(&result, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	// 3) 回写缓存（best-effort）
	if d.cache != nil {
		_ = cache.Set(ctx, d.cache, key, &result, 10*time.Minute)
	}
	return &result, nil
}

// Create 创建记录；成功后会写一份缓存副本。
func (d *BaseDAO[T]) Create(ctx context.Context, record *T) error {
	if d.gormDB == nil {
		return db.ErrNotInitialized
	}
	if err := d.gormDB.WithContext(ctx).Create(record).Error; err != nil {
		return err
	}
	// 回写缓存
	if d.cache != nil {
		id := d.pkExtractor(record)
		_ = cache.Set(ctx, d.cache, d.cacheKey(id), record, 10*time.Minute)
	}
	return nil
}

// Update 更新记录（全字段）；先写 DB 再删缓存。
func (d *BaseDAO[T]) Update(ctx context.Context, record *T) error {
	if d.gormDB == nil {
		return db.ErrNotInitialized
	}
	if err := d.gormDB.WithContext(ctx).Save(record).Error; err != nil {
		return err
	}
	// 删除缓存
	if d.cache != nil {
		id := d.pkExtractor(record)
		_ = cache.Del(ctx, d.cache, d.cacheKey(id))
	}
	return nil
}

// Delete 按主键删除记录；DB 删完后清缓存。
func (d *BaseDAO[T]) Delete(ctx context.Context, id int64) error {
	if d.gormDB == nil {
		return db.ErrNotInitialized
	}
	// 使用模型零值删除
	var zero T
	if err := d.gormDB.WithContext(ctx).Delete(&zero, id).Error; err != nil {
		return err
	}
	// 删除缓存
	if d.cache != nil {
		_ = cache.Del(ctx, d.cache, d.cacheKey(id))
	}
	return nil
}

// ListAll 查询所有记录，不走缓存。
func (d *BaseDAO[T]) ListAll(ctx context.Context) ([]T, error) {
	if d.gormDB == nil {
		return nil, db.ErrNotInitialized
	}
	var results []T
	if err := d.gormDB.WithContext(ctx).Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

// FindBy 通用条件查询，不走缓存。
// field 为列名，val 为查询值。
func (d *BaseDAO[T]) FindBy(ctx context.Context, field string, val interface{}) ([]T, error) {
	if d.gormDB == nil {
		return nil, db.ErrNotInitialized
	}
	var results []T
	if err := d.gormDB.WithContext(ctx).Where(fmt.Sprintf("%s = ?", field), val).Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}
