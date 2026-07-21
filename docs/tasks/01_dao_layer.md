# 任务：生成 DAO 层（泛型 BaseDAO + 各模型 DAO）

## 项目背景
- 项目根目录：/Users/zhengzihang/wgame-server
- Go module 名：wgame-server
- 已有 98 个 GORM 模型在 server/model/ 目录下（每个模型一个 .go 文件）
- 已有缓存抽象层：server/cache/cache.go（Cache 接口）、server/db/db.go（GORM + Cache 全局访问）
- 现有一个手写的 DAO 样例：server/dao/user_dao.go

## 目标
1. 先实现一个泛型 BaseDAO[T]，消除 UserDAO 中重复的 CRUD + Cache-Aside 模板代码
2. 然后把 server/model/ 下其余 97 个模型按 BaseDAO 模式补齐 DAO（无需手写每个）

## 必读参考文件（开始前务必 Read）
- /Users/zhengzihang/wgame-server/server/dao/user_dao.go （现有 DAO 模板，Cache-Aside 模式）
- /Users/zhengzihang/wgame-server/server/cache/cache.go （Cache 接口 + Get/Set/Del 通用函数）
- /Users/zhengzihang/wgame-server/server/db/db.go （GORM() / Cache() 全局访问）
- /Users/zhengzihang/wgame-server/server/model/user.go （模型样例，含主键自增）
- /Users/zhengzihang/wgame-server/server/model/characters.go （复杂模型样例）
- /Users/zhengzihang/wgame-server/server/model/experience.go （无主键模型样例）

## 设计要求

### 1. 泛型 BaseDAO[T]
位置：/Users/zhengzihang/wgame-server/server/dao/base_dao.go
要求：
- 使用 Go 1.21+ 泛型：`type BaseDAO[T any] struct`
- 持有 `gormDB *gorm.DB`、`cache cache.Cache`、`cacheKeyPrefix string`
- 约束：T 必须实现一个接口约定来暴露主键与表名。可以定义接口：
    ```go
    type Model interface {
        GetID() int64   // 或 interface{} 兼容 int32/int64 主键
    }
    ```
  或者用泛型 + 函数回调方式传入主键提取函数（更灵活，推荐）：
    ```go
    func NewBaseDAO[T any](gormDB *gorm.DB, cache cache.Cache, tableName string, pkExtractor func(*T) int64) *BaseDAO[T]
    ```
- 通用方法：
    - `GetByID(ctx, id) (*T, error)` — Cache-Aside：先 cache.Get 未命中再 DB.First 再回写
    - `Create(ctx, *T) error` — 成功后 cache.Set
    - `Update(ctx, *T) error` — Save 后 cache.Del
    - `Delete(ctx, id) error` — DB 删除后 cache.Del
    - `ListAll(ctx) ([]T, error)` — 不走缓存
    - `FindBy(ctx, field string, val interface{}) ([]T, error)` — 通用条件查询，不走缓存
- 缓存 key 格式：`wgame:{tableName}:{id}`，TTL 10 分钟
- 未命中/缓存出错时降级到 DB（参考 user_dao.go 的处理）
- 注意 `gorm.ErrRecordNotFound` 要转成返回 `(nil, nil)`

### 2. 模型接口约定
为了让 BaseDAO 能提取主键，请为每个模型补一个 `GetID()` 方法。
推荐做法：在 server/model/ 下新增一个 `model_id.go` 文件，用类型开关集中实现，避免改动 98 个文件：

```go
package model

func GetID(m any) int64 {
    switch v := m.(type) {
    case *User:         return int64(v.ID)
    case *Characters:   return int64(v.ID)
    case *CharaNickname: return v.ID   // 已是 int64
    // ...
    }
    return 0
}
```

（遍历 server/model/*.go 读取每个模型的主键字段类型生成此文件）

### 3. 具体 DAO 文件
位置：/Users/zhengzihang/wgame-server/server/dao/
对每个 model 生成对应的 dao 文件（如 `characters_dao.go`），内容非常薄：

```go
type CharactersDAO = BaseDAO[model.Characters]

func NewCharactersDAO() *CharactersDAO {
    return NewBaseDAO[model.Characters](db.GORM(), db.Cache(), "characters",
        func(t *model.Characters) int64 { return int64(t.ID) })
}
```

对于无主键模型（experience.go、experience_treasure.go），不生成 DAO，在汇报里列出。

### 4. 用户既有 DAO 保留
不要删除 user_dao.go，但要重构它基于 BaseDAO（让它成为 `BaseDAO[model.User]` 的别名或薄封装），保持 `NewUserDAO` 签名兼容。

## 翻译/实现规则
- 缓存 key 前缀统一用 model 的 `TableName()`
- 所有方法接收 `context.Context` 作为第一参数
- 错误处理：DB 错误原样返回；缓存错误只记日志（可用 fmt.Println 或留 TODO，项目暂无统一日志到 DAO 层的约定）
- 不写测试，不改 main.go，不改 model 文件（除新增 model_id.go）

## 验证
完成所有文件后执行：

```bash
cd /Users/zhengzihang/wgame-server && go vet ./... && go build ./...
```

必须全部通过。最后用一条消息汇报：
- BaseDAO 实现位置与关键设计
- 生成的具体 DAO 文件数
- 跳过的无主键模型列表
- model_id.go 的实现策略
