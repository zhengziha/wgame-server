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

// UserDAO 用户数据访问对象，基于 BaseDAO 实现通用 CRUD。
// 同时保留 User 特有的业务方法（如 GetByAccountName）。
type UserDAO struct {
	*BaseDAO[model.User]
}

// NewUserDAO 构造一个 UserDAO，使用全局 GORM + Cache。
func NewUserDAO() *UserDAO {
	return &UserDAO{
		BaseDAO: NewBaseDAO[model.User](
			db.GORM(),
			db.Cache(),
			"t_user",
			func(t *model.User) int64 { return t.ID },
		),
	}
}

// NewUserDAOWith 构造一个 UserDAO，传入自定义 DB/Cache。
// 适合单元测试或特殊业务路径。
func NewUserDAOWith(g *gorm.DB, c cache.Cache) *UserDAO {
	return &UserDAO{
		BaseDAO: NewBaseDAO[model.User](
			g,
			c,
			"t_user",
			func(t *model.User) int64 { return t.ID },
		),
	}
}

// userCacheKeyByName 按账号名生成缓存 key（用于登录路径）
func userCacheKeyByName(name string) string {
	return fmt.Sprintf("wgame:user:name:%s", name)
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
		_ = cache.Set(ctx, d.cache, d.cacheKey(u.ID), &u, 10*time.Minute)
	}
	return &u, nil
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
		_ = cache.Del(ctx, d.cache, d.cacheKey(id))
	}
	return nil
}
