// Package main 是 wgame-server 的入口。
//
// 启动流程：
//  1. 初始化日志；
//  2. 触发 demo handler 包的 init() 自注册；
//  3. 启动 TCP socket server。
//
// 整个进程只暴露一个 TCP 监听端口，使用自定义 10 字节头协议。
package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"wgame-server/comm/log"
	"wgame-server/comm/main_thread"
	"wgame-server/server/network/socket"

	// 匿名导入以触发 init() 自注册
	_ "wgame-server/server/demo/handlers"
)

func main() {
	addr := flag.String("addr", ":8800", "TCP listen address")
	flag.Parse()

	srv := socket.NewServer(socket.Config{Addr: *addr})

	// 连接建立/断开回调示例（运行在主线程）
	srv.OnConnect(func(_ *socket.SocketCmdContext) {
		// 示例：什么都不做；业务层可在此发欢迎消息 / 记录上线日志。
	})
	srv.OnDisconnect(func(c *socket.SocketCmdContext) {
		log.Info("[main] disconnect sid=%d uid=%d", c.GetSessionId(), c.GetUserId())
	})

	// 主线程保持空跑：所有业务在 main_thread.Process 中被 handler 触发。
	go func() {
		if err := srv.Start(); err != nil {
			log.Error("[main] server start failed: %v", err)
			os.Exit(1)
		}
	}()

	// 阻塞等待信号
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	log.Info("[main] recv signal %v, shutting down...", <-sig)
	srv.Stop()

	// 让 main_thread 排空（简化处理）
	_ = main_thread.Process
	log.Info("[main] bye")
}
