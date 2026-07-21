# DAO 层实施计划

## 一、现状分析

### 1. 模型统计
- 总模型数：99 个
- 有主键模型：97 个（int32 类型 96 个，int64 类型 3 个）
- 无主键模型：2 个（Experience、ExperienceTreasure）
- 需要补全主键的模型：2 个（补全后所有模型都有主键）

### 2. 主键类型分布
| 主键类型 | 数量 | 代表模型 |
|----------|------|----------|
| int32    | 96   | Characters、Pet、Npc、Experience、ExperienceTreasure 等 |
| int64    | 3    | User、SysUser、CharaNickname |

**统一处理**：
- BaseDAO 的 `pkExtractor` 函数返回 `int64` 类型
- int32 主键在提取时自动转换为 int64
- 所有缓存 key 使用统一的 int64 主键格式

### 3. 特殊情况
- `UpgradeExperience` 主键字段名为 `Level`（非 `ID`）
- `Experience` 和 `ExperienceTreasure` 需要补全主键 tag

## 二、实施步骤

### 步骤 0：补全模型主键
**文件**：
- `/Users/zhengzihang/wgame-server/server/model/experience.go`
- `/Users/zhengzihang/wgame-server/server/model/experience_treasure.go`

**修改内容**：
为 `Attrib` 字段添加 `gorm:"primaryKey"` tag：

```go
// Experience 升级经验配置表。
type Experience struct {
    // Attrib 等级（SQL 中为 PRIMARY KEY）
    Attrib int32 `gorm:"primaryKey;autoIncrement;column:attrib" json:"attrib"`
    // ... 其他字段
}

// ExperienceTreasure 升级经验宝箱配置表。
type ExperienceTreasure struct {
    // Attrib 等级（SQL 中为 PRIMARY KEY）
    Attrib int32 `gorm:"primaryKey;autoIncrement;column:attrib" json:"attrib"`
    // ... 其他字段
}
```

### 步骤 1：创建 BaseDAO 泛型实现
**文件**：`/Users/zhengzihang/wgame-server/server/dao/base_dao.go`

**设计要点**：
```go
type BaseDAO[T any] struct {
    gormDB         *gorm.DB
    cache          cache.Cache
    tableName      string
    cacheKeyPrefix string
    pkExtractor    func(*T) int64
}

func NewBaseDAO[T any](
    gormDB *gorm.DB, 
    cache cache.Cache, 
    tableName string,
    pkExtractor func(*T) int64,
) *BaseDAO[T]
```

**通用方法**：
- `GetByID(ctx, id int64) (*T, error)` — Cache-Aside 模式
- `Create(ctx, *T) error` — 成功后写缓存
- `Update(ctx, *T) error` — 更新后删缓存
- `Delete(ctx, id int64) error` — 删除后清缓存
- `ListAll(ctx) ([]T, error)` — 不走缓存
- `FindBy(ctx, field string, val interface{}) ([]T, error)` — 通用条件查询

**缓存策略**：
- Key 格式：`wgame:{tableName}:{id}`
- TTL：10 分钟
- 未命中/缓存出错时降级到 DB
- `gorm.ErrRecordNotFound` 转为 `(nil, nil)`

### 步骤 2：创建 model_id.go
**文件**：`/Users/zhengzihang/wgame-server/server/model/model_id.go`

**实现策略**：
```go
package model

// GetModelID 统一提取模型主键（int64）
// 对于 int32 主键的模型，会自动转换为 int64
func GetModelID(m any) int64 {
    switch v := m.(type) {
    case *User:
        return v.ID  // int64
    case *SysUser:
        return v.ID  // int64
    case *CharaNickname:
        return v.ID  // int64
    case *Characters:
        return int64(v.ID)  // int32 -> int64
    case *Experience:
        return int64(v.Attrib)  // int32 -> int64
    case *ExperienceTreasure:
        return int64(v.Attrib)  // int32 -> int64
    // ... 为所有有主键的模型添加 case
    }
    return 0
}
```

**主键类型说明**：
- int64 类型：User、SysUser、CharaNickname（3个）
- int32 类型：其余 96 个模型（包括 Experience、ExperienceTreasure）
- 所有主键最终统一转换为 int64 返回

### 步骤 3：为每个模型生成 DAO 文件
**位置**：`/Users/zhengzihang/wgame-server/server/dao/`

**生成规则**：
1. 文件名：`{snake_case_model}_dao.go`（如 `characters_dao.go`）
2. 代码模板：
```go
package dao

import (
    "wgame-server/server/db"
    "wgame-server/server/model"
)

type {Model}DAO = BaseDAO[model.{Model}]

func New{Model}DAO() *{Model}DAO {
    return NewBaseDAO[model.{Model}](
        db.GORM(), 
        db.Cache(), 
        "{table_name}",
        func(t *model.{Model}) int64 { 
            return int64(t.ID)  // 根据实际主键类型调整
        },
    )
}
```

**需要生成的模型列表（99个）**：
- 所有有主键的模型（包括补全主键后的 Experience 和 ExperienceTreasure）

### 步骤 4：重构 UserDAO
**文件**：`/Users/zhengzihang/wgame-server/server/dao/user_dao.go`

**重构方案**：
- 保留 `UserDAO` 结构体
- 内部使用 `BaseDAO[model.User]` 作为嵌入字段或委托
- 保持 `NewUserDAO()` 签名兼容
- 保留 `GetByAccountName` 等业务特有方法

## 三、补全主键的模型

需要为以下模型补上 `gorm:"primaryKey"` tag：

| 模型名 | 文件名 | 补全字段 | 原因 |
|--------|--------|----------|------|
| Experience | experience.go | `Attrib` 字段添加 `gorm:"primaryKey"` | SQL DDL 中有主键，Go struct 未标注 |
| ExperienceTreasure | experience_treasure.go | `Attrib` 字段添加 `gorm:"primaryKey"` | SQL DDL 中有主键，Go struct 未标注 |

## 四、验证步骤

1. 运行 `go vet ./...` 检查代码规范
2. 运行 `go build ./...` 确保编译通过
3. 检查 UserDAO 向后兼容性

## 五、预期产出

1. **补全主键**：修改 2 个模型文件（experience.go、experience_treasure.go）
2. `base_dao.go` — 1 个泛型基础 DAO
3. `model_id.go` — 1 个主键提取函数
4. `*_dao.go` — 99 个具体模型 DAO（包括 Experience 和 ExperienceTreasure）
5. 重构 `user_dao.go` — 保持向后兼容

总计：103 个文件变更（含重构的 user_dao.go 和补全的 2 个模型文件）