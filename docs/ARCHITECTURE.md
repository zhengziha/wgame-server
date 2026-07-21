# wgame-server 架构文档

> 基于 hero_story.go_server 框架结构，使用 Go 语言重写 wd-server-fl（Java Netty）项目的 socket 版游戏服务器基础框架。

## 1. 项目概览

### 1.1 项目定位

- **目标**：用 Go 重写 wd-server-fl Java 项目，复用 hero_story 的并发模型与业务骨架
- **传输层**：raw TCP socket（替代 hero_story 的 WebSocket，也替代 Java Netty）
- **协议**：自定义二进制协议（10 字节头 + 变长 body + 替换式加密），与 wd-server-fl 客户端完全兼容
- **业务范围**：98 个 GORM 模型已翻译，DAO 层已生成，handler 层待补齐（详见 [docs/tasks/](./tasks/)）

### 1.2 技术栈

| 层 | 选型 | 说明 |
|---|---|---|
| 语言 | Go 1.25 | module: `wgame-server` |
| ORM | gorm.io/gorm v1.31 | 通过 Dialector 抽象 |
| SQLite Driver | github.com/glebarez/sqlite | 纯 Go，**无需 CGO** |
| MySQL Driver | gorm.io/driver/mysql | 生产推荐 |
| Postgres Driver | gorm.io/driver/postgres | 可选 |
| Redis | github.com/redis/go-redis/v9 | 缓存后端 |
| GBK 编码 | golang.org/x/text | 对齐 Java 的 GBK 字符串 |
| 配置 | github.com/spf13/viper | YAML + 环境变量 + flag |

## 2. 顶层目录结构

```
wgame-server/
├── main.go                      # 进程入口：配置加载 → DB/Cache 初始化 → AutoMigrate → 启动 TCP server
├── config/                      # 配置层（viper）
│   ├── config.go
│   └── default.yml              # 内置默认配置（可被外部 config.yml 覆盖）
├── comm/                        # 通用基础库（无业务依赖）
│   ├── log/                     #   按天分文件的日志
│   ├── main_thread/             #   主线程串行队列（业务同步骨架）
│   └── async_op/                #   分片异步 worker 池
├── server/                      # 服务器主体
│   ├── codec/                   #   协议编解码（frame + 加密表 + GameReader/Writer）
│   ├── context/                 #   MyCmdContext 抽象接口
│   ├── session/                 #   GameSession（会话状态）
│   ├── msg/                     #   InMessage / OutMessage 消息抽象
│   ├── cache/                   #   Cache 接口 + Redis/Memory 实现
│   ├── db/                      #   GORM + Dialector 工厂
│   ├── model/                   #   98 个 GORM 模型 + AllModels 注册表
│   ├── dao/                     #   BaseDAO 泛型 + 98 个具体 DAO
│   ├── network/                 #   网络层
│   │   ├── socket/              #     raw TCP server 实现
│   │   ├── firewall/            #     入站防火墙（限流/WPE）
│   │   ├── handler/             #     cmd 自注册派发
│   │   └── broadcaster/         #     在线连接索引 + 广播
│   └── demo/                    #   示例 handler 与消息
├── cmd/client/                  # 测试客户端（验证编解码与 echo）
└── docs/                        # 文档
    ├── ARCHITECTURE.md          # 本文件
    └── tasks/                   # 后续任务提示词
```

## 3. 分层架构

```
┌────────────────────────────────────────────────────────────┐
│                      main.go（入口）                        │
│   配置加载 → DB/Cache 初始化 → AutoMigrate → Start TCP      │
└────────────────────────────┬───────────────────────────────┘
                             │
        ┌────────────────────┼────────────────────┐
        │                    │                    │
        ▼                    ▼                    ▼
┌──────────────┐    ┌────────────────┐   ┌──────────────┐
│  网络层       │    │   数据层        │   │  并发模型     │
│              │    │                │   │              │
│ socket/      │    │ db/ (GORM)     │   │ main_thread/ │
│ firewall/    │    │ cache/         │   │ async_op/    │
│ handler/     │    │ dao/           │   │              │
│ broadcaster/ │    │ model/         │   │              │
└──────┬───────┘    └────────────────┘   └──────────────┘
       │
       ▼
┌──────────────────────────────────────────────────────────┐
│                      业务 handler 层                      │
│   （通过 init() 自注册到 handler.Register，无需改 main）  │
└──────────────────────────────────────────────────────────┘
       │
       ▼
┌──────────────────────────────────────────────────────────┐
│             编解码层 codec/ + msg/                       │
│   10 字节头 / 35 张加密表 / GBK / GameReader / GameWriter │
└──────────────────────────────────────────────────────────┘
```

## 4. 核心组件详解

### 4.1 协议编解码层（server/codec/）

完整对应 Java wd-server-fl 的 `GameReadTool` / `GameWriteTool` / `EncryptTable`。

#### 4.1.1 帧格式（10 字节头）

| 偏移 | 长度 | 字段 | 说明 |
|---|---|---|---|
| 0 | 2 | magic | uint16 BE = **19802 (0x4D6A)** |
| 2 | 2 | tableIndex | uint16 BE；1..35 = 加密表索引；0 = 不加密 |
| 4 | 4 | tickCount | int32 BE；客户端 tick；服务端写 0 |
| 8 | 2 | length | uint16 BE = `2 + len(payload)` |
| 10 | 2 | cmd | uint16 BE |
| 12 | ... | payload | 业务消息体 |

- 单帧最大 **10240** 字节（对齐 Java `LengthFieldBasedFrameDecoder(10240, 8, 2, 0, 4)`）
- 加密范围：`bytes[10 : 10+length]`（cmd + payload）
- 字节序：**大端**
- 字符串编码：**GBK**

参见 [frame.go](../server/codec/frame.go)。

#### 4.1.2 加密

- 35 张 256 项的替换表（`[35][256]byte`），对应 Java `EncryptTable.java`
- 出站：`EncryptBody(tableIndex, body)` 按字节查表替换
- 入站：当前**不解密**（与 Java `ServerHandler` 行为一致）
- 原始数据分两个文件存放：[encrypt_table_raw1.go](../server/codec/encrypt_table_raw1.go)（前 16 行）、[encrypt_table_raw2.go](../server/codec/encrypt_table_raw2.go)（后 19 行）

#### 4.1.3 读写工具

- [game_read_tool.go](../server/codec/game_read_tool.go)：`GameReader` 提供 `ReadInt/ReadUInt/ReadShort/ReadLong/ReadUByte/ReadBoolean/ReadString/ReadString2/ReadString4/ReadBytes/Skip`
- [game_write_tool.go](../server/codec/game_write_tool.go)：`GameWriter` 提供对应写方法
- [reflect_codec.go](../server/codec/reflect_codec.go)：反射自动编解码，按 struct 字段顺序读写

### 4.2 消息抽象（server/msg/）

- `InMessage` / `OutMessage` 接口，只需实现 `Cmd() uint16`
- 可选 `CustomCodec` 接口（`WriteBody(w)` / `ReadBody(r)`）用于自定义复杂消息
- 未实现 `CustomCodec` 时，走反射自动读写
- `msg.WriteFrame(m, tableIndex, tickCount)` 把 OutMessage 打包成完整出站帧

### 4.3 网络层（server/network/）

#### 4.3.1 TCP Server（socket/server.go）

```
Accept loop（1 个 goroutine）
    │
    └── 每条新连接 → handleConn（独立 goroutine）
            ├── 构造 SocketCmdContext（含 firewall + session）
            ├── 注册到 broadcaster
            ├── sendLoop（独立 goroutine，消费 sendQ）
            └── readLoop（当前 goroutine）
                   ├── FrameReader.ReadFrame() 逐帧解析
                   ├── firewall.Check() 限流/WPE 检测
                   └── main_thread.Process(handler.Dispatch)
                         ↑
                    投递到主线程串行执行
```

- **每个连接 2 个 goroutine**：read + send
- **业务 handler 始终在主线程执行**，天然无 data race
- 连接断开时：关闭 conn → sendLoop 退出 → readLoop 退出 → 从 broadcaster 移除

#### 4.3.2 防火墙（firewall/firewall.go）

单连接维度：
- 最小包间隔检测：默认 20ms（每秒最多 ~50 包）
- 滑动窗口限流：5 秒内最多 300 包
- 日志限流：同类告警 5 秒最多一条

#### 4.3.3 Handler 派发（handler/handler_factory.go）

```go
// 业务侧用法（每个 cmd 一个 init）
func init() {
    handler.Register(0x0101, "Echo", EchoHandler)
}

func EchoHandler(ctx context.MyCmdContext, frame *codec.Frame, reader *codec.GameReader) error {
    text := reader.ReadString()
    ctx.Write(&EchoResp{Text: text})
    return nil
}
```

- `handler.Register` 在 `init()` 中调用，**插件式自注册**
- `main.go` 只需匿名 import handler 包，无需维护 cmd 路由表
- 重复注册覆盖旧值（保留最后注册的 Name）

#### 4.3.4 广播器（broadcaster/broadcaster.go）

- 双索引：`sessionId → ctx` + `userId → ctx`
- 接口：`Broadcast` / `BroadcastExcept` / `SendToUser` / `GetCmdCtxByUserId`
- 线程安全（内部 RWMutex）

### 4.4 并发模型（comm/）

| 组件 | 角色 | 关键参数 |
|---|---|---|
| `main_thread` | 单 goroutine 串行业务队列 | 队列容量 2048；满了丢弃并限速告警 |
| `async_op` | 2048 路分片 worker 池 | `bindId % 2048` 路由；同 bindId 串行 |

**典型流程**：
1. IO goroutine 读到帧 → `main_thread.Process(handler.Dispatch)`
2. handler 在主线程执行 → 需要 DB 聚合时投 `async_op.Process(bindId, work, callback)`
3. worker 执行完后，`callback` 回投到主线程

### 4.5 上下文与会话

- [context/cmd_context.go](../server/context/cmd_context.go)：`MyCmdContext` 接口，传输无关（`BindUserId/GetUserId/Write/SendError/Disconnect`）
- [session/game_session.go](../server/session/game_session.go)：`GameSession` 逻辑会话
  - 原子字段：`id`（玩家 id）、`isConnection`（在线状态）、`lock`（CAS 串行锁）
  - 业务扩展位：`Chara interface{}`（业务侧自行替换为具体类型）

### 4.6 数据层

#### 4.6.1 GORM 初始化（db/db.go）

- 双数据源：`GORM()`（游戏库） + `AuthGORM()`（账号库）
- 全局单例 + `sync.Once` 幂等初始化
- Dialector 工厂：`db/sqlite|mysql|postgres`，可扩展（`InitWithDialector`）
- 连接池自适应：SQLite 默认 5 连接，MySQL/PG 默认 32

#### 4.6.2 缓存抽象（cache/cache.go）

```go
type Cache interface {
    GetRaw(ctx, key) ([]byte, error)
    SetRaw(ctx, key, val, ttl) error
    Del(ctx, keys...) error
    Exists(ctx, key) (bool, error)
    Close() error
}
```

- JSON 序列化在抽象层（`cache.Get/Set`），实现只负责字节存取
- 内置实现：`RedisCache`、`MemoryCache`
- 业务层注入：`db.SetCache(impl)`

#### 4.6.3 DAO 层（dao/）

- 泛型 `BaseDAO[T]` 提供 6 个通用方法（`GetByID/Create/Update/Delete/ListAll/FindBy`）
- Cache-Aside 模式：读未命中回写、写后失效
- 缓存 key：`wgame:{table}:{pk}`，TTL 10 分钟
- 98 个具体 DAO 通过类型别名生成，例如：

```go
type CharactersDAO = BaseDAO[model.Characters]
func NewCharactersDAO() *CharactersDAO {
    return NewBaseDAO[model.Characters](db.GORM(), db.Cache(), "characters",
        func(t *model.Characters) int64 { return int64(t.ID) })
}
```

- 业务专属方法挂在子 struct 上（参考 [user_dao.go](../server/dao/user_dao.go) 的 `GetByAccountName/UpdateLevel`）

#### 4.6.4 模型层（model/）

- 98 个 GORM 模型，对应 wd-server-fl 全部可持久化实体
- [models.go](../server/model/models.go)：`AllModels()` 返回所有需要 AutoMigrate 的模型
- 每个模型实现 `TableName() string`，表名来自 SQL 建表语句

### 4.7 配置层（config/）

- YAML + 环境变量 + 命令行 flag 三层覆盖
- 优先级：**命令行 > 环境变量 > 配置文件 > 默认值**
- 双数据源支持：`game_db` + `auth_db` 独立配置

```yaml
server:
  addr: ":8800"
game_db:
  driver: "sqlite"        # sqlite | mysql | postgres
  dsn: "data/game.db"
cache:
  driver: "redis"         # redis | memory | none
  redis: { addr: "127.0.0.1:6379" }
```

## 5. 启动流程

```
1. flag.Parse()              解析命令行
2. config.Load(path)         加载 YAML/ENV/默认值
3. applyFlags()              命令行覆盖配置
4. db.Init(gameDB, nil)      初始化游戏库
5. db.InitAuth(authDB)       初始化认证库（可选）
6. buildCache() + SetCache   初始化缓存
7. db.AutoMigrate(AllModels) 自动建表
8. socket.NewServer().Start()启动 TCP（goroutine）
9. signal.Notify(SIGINT)     阻塞等待信号
10. srv.Stop() + db.Close()  优雅关闭
```

## 6. Java 概念映射

| Java wd-server-fl | Go wgame-server |
|---|---|
| `Netty ServerBootstrap` | [socket.Server](../server/network/socket/server.go) |
| `LengthFieldBasedFrameDecoder(10240, 8, 2, 0, 4)` | [codec.FrameReader](../server/codec/frame.go) |
| `GameReadTool` / `GameWriteTool` | [game_read_tool.go](../server/codec/game_read_tool.go) / [game_write_tool.go](../server/codec/game_write_tool.go) |
| `EncryptTable`（35 行替换表） | [encrypt_table.go](../server/codec/encrypt_table.go) |
| `BaseWrite`（10 字节头 + cmd + body） | [codec.EncodeFrame](../server/codec/frame.go) |
| `ServerHandler.channelRead` switch case | [handler.Dispatch](../server/network/handler/handler_factory.go) |
| `GameSession.lock` CAS | [GameSession.TryLock](../server/session/game_session.go) |
| `XXXService.method()` | `dao.NewXXXDAO().Method()` |
| `broadcaster.broadcast` | [broadcaster.Broadcast](../server/network/broadcaster/broadcaster.go) |

| hero_story.go_server 模式 | Go wgame-server |
|---|---|
| WebSocket 传输 | raw TCP（自定义协议） |
| `MyCmdContext` 抽象 | [context.MyCmdContext](../server/context/cmd_context.go) |
| `main_thread.Process` 串行 | [main_thread.Process](../comm/main_thread/process.go) |
| `async_op.Process` 分片 | [async_op.Process](../comm/async_op/process.go) |
| `broadcaster` 广播 | [broadcaster](../server/network/broadcaster/broadcaster.go) |

## 7. 解耦与可替换性

| 维度 | 抽象方式 | 可选实现 |
|---|---|---|
| 数据库 | `gorm.Dialector` | SQLite / MySQL / Postgres / 自定义 |
| 缓存 | `cache.Cache` 接口 | Redis / Memory / 自定义 |
| 传输 | `MyCmdContext` 接口 | TCP（当前）/ WebSocket / KCP |
| 配置 | viper | YAML / ENV / flag |
| 业务 handler | `init()` 自注册 | 插件式，无需改 main |

## 8. 验证与运行

### 8.1 编译验证

```bash
go vet ./...
go build ./...
```

### 8.2 启动（本地开发）

```bash
# SQLite + Memory（零外部依赖）
go run . -db-driver=sqlite -db-dsn=data/game.db -cache-driver=memory

# MySQL + Redis（生产）
go run . \
  -db-driver=mysql \
  -db-dsn="user:pass@tcp(127.0.0.1:3306)/game?charset=utf8mb4&parseTime=True&loc=Local" \
  -cache-driver=redis -redis-addr=127.0.0.1:6379
```

### 8.3 端到端测试

```bash
# 终端 1：启动服务
go run . -addr=:8800 -cache-driver=memory

# 终端 2：跑测试客户端
go run ./cmd/client -addr=127.0.0.1:8800 -text=hello
```

客户端会发送 cmd=0x0101 的 Echo 请求，服务端回 cmd=0x0001 的 EchoResp，客户端校验内容一致后输出 `RESULT: PASS`。

## 9. 后续路线图

详见 [docs/tasks/](./tasks/)：

1. **[DAO 层](./tasks/01_dao_layer.md)**：已完成（98 个 DAO）
2. **[AutoMigrate](./tasks/02_auto_migrate.md)**：已完成（98 个模型注册）
3. **[业务 handler 翻译](./tasks/03_handler_translation.md)**：待执行（100+ 个 cmd）
