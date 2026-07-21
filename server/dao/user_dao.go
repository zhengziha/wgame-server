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
	"wgame-server/server/model"
)

// UserDAO 用户数据访问对象
type UserDAO struct {
	gormDB *gorm.DB
	cache  cache.Cache
}

// NewUserDAO 构造一个 UserDAO，使用全局 GORM + Cache。
func NewUserDAO() *UserDAO {
	return &UserDAO{
		gormDB: db.GORM(),
		cache:  db.Cache(),
	}
}

// NewUserDAOWith 构造一个 UserDAO，传入自定义 DB/Cache。
// 适合单元测试或特殊业务路径。
func NewUserDAOWith(g *gorm.DB, c cache.Cache) *UserDAO {
	return &UserDAO{gormDB: g, cache: c}
}

// userCacheKey 生成 User 缓存 key
func userCacheKey(id int64) string {
	return fmt.Sprintf("wgame:user:%d", id)
}

// userCacheKeyByName 按账号名生成缓存 key（用于登录路径）
func userCacheKeyByName(name string) string {
	return fmt.Sprintf("wgame:user:name:%s", name)
}

// GetByID 按 id 查询用户，先走 Cache 再回源 SQLite。
func (d *UserDAO) GetByID(ctx context.Context, id int64) (*model.User, error) {
	// 1) 查缓存
	key := userCacheKey(id)
	var u model.User
	if d.cache != nil {
		if err := cache.Get(ctx, d.cache, key, &u); err == nil {
			return &u, nil
		}
	}

	// 2) 查 DB
	if d.gormDB == nil {
		return nil, db.ErrNotInitialized
	}
	if err := d.gormDB.WithContext(ctx).First(&u, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	// 3) 回写缓存（best-effort）
	if d.cache != nil {
		_ = cache.Set(ctx, d.cache, key, &u, 10*time.Minute)
	}
	return &u, nil
}

// GetByAccountName 按账号名查询用户（登录常用路径）
func (d *UserDAO) GetByAccountName(ctx context.Context, name string) (*model.User, error) {
	if d.gormDB == nil {
		return nil, db.ErrNotInitialized
	}
	var u model.User
	if err := d.gormDB.WithContext(ctx).Where("account_name = ?", name).First(&u).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	// 回写 id 索引缓存
	if d.cache != nil {
		_ = cache.Set(ctx, d.cache, userCacheKey(u.ID), &u, 10*time.Minute)
	}
	return &u, nil
}

// Create 创建用户；成功后会写一份缓存副本
func (d *UserDAO) Create(ctx context.Context, u *model.User) error {
	if d.gormDB == nil {
		return db.ErrNotInitialized
	}
	if err := d.gormDB.WithContext(ctx).Create(u).Error; err != nil {
		return err
	}
	if d.cache != nil {
		_ = cache.Set(ctx, d.cache, userCacheKey(u.ID), u, 10*time.Minute)
	}
	return nil
}

// Update 更新用户（全字段）；先写 DB 再删缓存。
func (d *UserDAO) Update(ctx context.Context, u *model.User) error {
	if d.gormDB == nil {
		return db.ErrNotInitialized
	}
	if err := d.gormDB.WithContext(ctx).Save(u).Error; err != nil {
		return err
	}
	if d.cache != nil {
		_ = cache.Del(ctx, d.cache, userCacheKey(u.ID))
	}
	return nil
}

// UpdateLevel 仅更新等级字段（演示增量更新 + 失效缓存）
func (d *UserDAO) UpdateLevel(ctx context.Context, id int64, level int) error {
	if d.gormDB == nil {
		return db.ErrNotInitialized
	}
	if err := d.gormDB.WithContext(ctx).Model(&model.User{}).
		Where("id = ?", id).
		Update("level", level).Error; err != nil {
		return err
	}
	if d.cache != nil {
		_ = cache.Del(ctx, d.cache, userCacheKey(id))
	}
	return nil
}

// Delete 删除用户；DB 删完后清缓存
func (d *UserDAO) Delete(ctx context.Context, id int64) error {
	if d.gormDB == nil {
		return db.ErrNotInitialized
	}
	if err := d.gormDB.WithContext(ctx).Delete(&model.User{}, id).Error; err != nil {
		return err
	}
	if d.cache != nil {
		_ = cache.Del(ctx, d.cache, userCacheKey(id))
	}
	return nil
}
