# 任务：注册所有 model 到 GORM AutoMigrate

## 项目背景
- 项目根目录：/Users/zhengzihang/wgame-server
- 模型目录：/Users/zhengzihang/wgame-server/server/model/ （98 个 .go 文件，每个文件一个 GORM 模型）
- 现有 main.go 中 AutoMigrate 只注册了 model.User：

```go
if err := db.AutoMigrate(&model.User{}); err != nil { ... }
```

- 位置在 main.go 第 60-63 行附近

## 必读参考文件
- /Users/zhengzihang/wgame-server/main.go （找到现有 AutoMigrate 调用位置）
- /Users/zhengzihang/wgame-server/server/model/user.go （模型样例）
- /Users/zhengzihang/wgame-server/server/db/db.go （AutoMigrate 函数签名）
- 遍历 /Users/zhengzihang/wgame-server/server/model/*.go 获取所有结构体名

## 目标
把 main.go 中 AutoMigrate 调用扩展为注册全部 model（98 个），让启动时自动建表/同步表结构到 SQLite/MySQL/Postgres。

## 实现要求

### 方案 A（推荐）：集中维护一个 AllModels 切片
在 server/model/ 下新增 `all_models.go`：

```go
package model

// AllModels 返回所有需要 AutoMigrate 的模型指针。
// 新增模型时在此处追加一行即可。
func AllModels() []interface{} {
    return []interface{}{
        &User{},
        &Accounts{},
        &Characters{},
        // ... 遍历 model 目录所有结构体
    }
}
```

然后 main.go 改为：

```go
if err := db.AutoMigrate(model.AllModels()...); err != nil { ... }
```

### 排除规则
以下模型不要加入 AllModels：
- 无主键模型（如果 gorm tag 没有 `primaryKey`）：experience.go、experience_treasure.go
- 仔细检查每个模型，如果发现没有 primaryKey tag 的，跳过并在汇报里列出

### 命名与顺序
- 按字母序排序，便于 diff review
- User 放第一个（作为示例/历史遗留）

## 限制
- 只新增 `all_models.go`，只修改 main.go 一行
- 不要修改其它任何 model 文件
- 不要改 `db.AutoMigrate` 签名

## 验证
完成后执行：

```bash
cd /Users/zhengzihang/wgame-server && go vet ./... && go build ./...
```

然后用内存数据库快速跑一次验证建表无错：

```bash
/tmp/wgame-server -addr=:18900 -db-driver=sqlite -db-dsn=:memory: -cache-driver=memory
```

预期日志出现 `automigrate done`，无 SQL 报错。
（如果 /tmp/wgame-server 不存在，先 `go build -o /tmp/wgame-server .`）

最后用一条消息汇报：
- all_models.go 中注册的模型总数
- 被排除的模型列表及原因
- 是否有 AutoMigrate 报错（例如某些表字段类型在 SQLite 上不兼容）
