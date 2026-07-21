# 任务：翻译 wd-server-fl 业务 handler 到 Go

## 项目背景

### 目标 Go 项目（你写代码的地方）
- 根目录：/Users/zhengzihang/wgame-server
- Go module：wgame-server
- 已有框架组件：
  - 自定义 TCP socket 协议（10 字节头），server/network/socket/
  - handler 自注册机制：server/network/handler/handler_factory.go 的 `handler.Register(cmd, name, fn)`
  - handler 函数签名：`func(ctx context.MyCmdContext, frame *codec.Frame, reader *codec.GameReader) error`
  - context 接口：server/context/cmd_context.go 的 `MyCmdContext`（GetUserId/Write/SendError/Disconnect 等）
  - 出站消息抽象：server/msg/base_write.go 的 `OutMessage` 接口（`Cmd()` + `WriteBody(*codec.GameWriter)`）
  - 编解码工具：server/codec/game_read_tool.go、game_write_tool.go（GBK + 大端）
  - 98 个 GORM 模型：server/model/*.go
  - DAO 层：server/dao/（参考 user_dao.go）

### 参考 Java 源码项目（你翻译的来源）
- 根目录：/Users/zhengzihang/my-src/java- projects/wd-server-fl
- 关键目录：
  - src/main/java/com/fengshen/server/ （Netty 协议层，GameReadTool/GameWriteTool）
  - src/main/java/com/fengshen/core/ （业务核心：GameSession、Chara、GameCore 等）
  - src/main/java/com/fengshen/db/service/ （Java Service 层，业务逻辑多在此）
- 协议字节布局与 Go 端完全一致（已校准）：`magic(2) + tableIndex(1) + tickCount(4) + length(2) + cmd(2) + body`

## 已有 Go 端 handler 示例（必读，模仿其结构）
- /Users/zhengzihang/wgame-server/server/demo/handlers/handlers.go （Echo handler）
- /Users/zhengzihang/wgame-server/server/demo/handlers/user_handler.go （DAO 调用示例）
- /Users/zhengzihang/wgame-server/server/demo/msg/echo_msg.go （OutMessage 实现）

## 翻译范围与分批策略
wd-server-fl 有 100+ 个 cmd。请先：
1. Grep 搜索 Java 源码中所有的 cmd 派发入口（通常在 ServerHandler.java 或类似文件，搜 `case 0x` 或 switch case 数字）
2. 汇总一份 `cmd 编号 -> 业务名 -> Java 源文件位置` 的清单
3. 然后询问用户/或按登录流程优先级分批翻译：
   - 第一批：登录/账号（账号校验、角色加载、选角）
   - 第二批：基础信息（属性、背包、装备）
   - 第三批：社交（好友、帮派、聊天）
   - 第四批：战斗/养成
   - 第五批：其它

每批限制在 15-25 个 cmd，避免单批过大。

## 翻译规则

### 1. 目录结构
在 Go 端新建：
- `server/handler/<模块>/` 存放 handler（如 `server/handler/auth/`、`server/handler/chara/`）
- `server/msg/<模块>/` 存放 OutMessage 实现（如 `server/msg/auth/`）

每个 handler 文件对应一个 cmd，文件名用蛇形小写业务名（如 `login.go`）。

### 2. handler 注册
每个 handler 文件用 `init()` 自注册，模仿 demo/handlers：

```go
func init() {
    handler.Register(0x0100, "Login", LoginHandler)
}
```

cmd 编号必须与 Java 端完全一致。

### 3. handler 签名

```go
func XxxHandler(ctx myctx.MyCmdContext, frame *codec.Frame, reader *codec.GameReader) error
```

读取入参用 `reader.ReadXXX()`（与 Java GameReadTool 方法一一对应）。
返回响应用 `ctx.Write(msgObj)`，msgObj 实现 `msg.OutMessage`。

### 4. 业务逻辑翻译原则
- Java 的 `ServerHandler/channelRead` 中的 switch case 每个分支对应一个 Go handler
- Java 调用 `XXXService.method()` 的，Go 端调用对应的 `dao.NewXXXDAO().Method()`
- Java 用 `GameSession.getAttribute/setAttribute` 的，Go 端用 ctx 的 session 字段（参考 server/session/game_session.go）
- Java 的 `broadcaster.broadcast` 改用 server/network/broadcaster/broadcaster.go
- 业务逻辑保持等价，不要做"优化"或"重构"，只做直译
- 遇到 Java 中依赖未翻译的模块（如某些 runtime 对象 Chara/Pet），先在 Go 端定义等价 struct 占位，加 TODO 注释，不要硬编码绕过

### 5. 协议字段映射
- Java `GameReadTool.readInt()` -> Go `reader.ReadInt()`
- Java `GameReadTool.readString()` -> Go `reader.ReadString()`（GBK + 1 字节长度头）
- Java `GameReadTool.readString4()` -> Go `reader.ReadString4()`（4 字节长度头）
- Java GameWriteTool 同理对应 game_write_tool.go
- 所有数值用大端（两端已校准）

### 6. 错误处理
- 协议解析失败：`return err`（框架会记录日志）
- 业务错误：`ctx.SendError(errorCode, errorInfo)` 然后 `return nil`
- 参考 Java 的 ErrorCode 定义翻译一份 `server/errcode/errcode.go`

## 限制
- 不修改 server/codec/、server/network/、server/context/ 等框架层代码
- 不改 main.go（新 handler 通过 init() 自注册，无需改动 main）
- 每个 handler 文件不超过 200 行；超长说明业务太重，拆分成多个私有辅助函数

## 验证
每批完成后执行：

```bash
cd /Users/zhengzihang/wgame-server && go vet ./... && go build ./...
```

然后用 `cmd/client/main.go` 改造一个测试客户端，验证本批其中一个 cmd 能正常往返。

## 交付清单
第一批结束时汇报：
- Java 源码中 cmd 总数与分类清单
- 本批翻译的 cmd 列表（cmd hex -> Java 位置 -> Go 文件位置）
- 占位实现的依赖（列出 TODO）
- 遇到的协议歧义点（如有）
