package socket

import (
	"errors"
	"io"
	"net"
	"sync"
	"sync/atomic"

	"wgame-server/comm/async_op"
	"wgame-server/comm/log"
	"wgame-server/comm/main_thread"
	"wgame-server/server/codec"
	"wgame-server/server/network/broadcaster"
	"wgame-server/server/network/firewall"
	"wgame-server/server/network/handler"
)

// Config 服务启动配置
type Config struct {
	Addr string // 监听地址，例如 ":8800"
}

// Server TCP socket 服务器主体。
//
// 与 hero_story.go_server/biz_server/network/server.go 的核心差异：
//   - 传输层从 WebSocket 换成 raw TCP（使用自定义 10 字节头协议）
//   - 不再做 RFC6455 分帧，直接走 codec.FrameReader
type Server struct {
	cfg Config

	ln net.Listener

	sessionIdSeq int32 // 连接 id 自增

	closeOnce sync.Once
	closed    atomic.Bool

	onConnect    func(*SocketCmdContext) // 可选钩子
	onDisconnect func(*SocketCmdContext) // 可选钩子
}

// NewServer 构造一个 Server 实例（未启动）
func NewServer(cfg Config) *Server {
	return &Server{cfg: cfg}
}

// OnConnect 注册一个连接建立回调（可选）。
// 该回调运行在主线程，可安全访问 broadcaster / 数据。
func (s *Server) OnConnect(fn func(*SocketCmdContext)) {
	s.onConnect = fn
}

// OnDisconnect 注册一个连接断开回调（可选）。
// 该回调运行在主线程。
func (s *Server) OnDisconnect(fn func(*SocketCmdContext)) {
	s.onDisconnect = fn
}

// Start 阻塞接受连接。调用方应放在独立 goroutine 中执行。
func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.cfg.Addr)
	if err != nil {
		return err
	}
	s.ln = ln
	log.Info("[socket] server listening on %s", s.cfg.Addr)

	for {
		c, err := ln.Accept()
		if err != nil {
			if s.closed.Load() {
				return nil
			}
			log.Error("[socket] accept err: %v", err)
			continue
		}
		go s.handleConn(c)
	}
}

// Stop 关闭监听器（已建立的连接不会被立即关闭）
func (s *Server) Stop() {
	s.closeOnce.Do(func() {
		s.closed.Store(true)
		if s.ln != nil {
			_ = s.ln.Close()
		}
	})
}

// handleConn 处理一条新连接：构造 ctx、启动 read/send loop、等待退出。
func (s *Server) handleConn(c net.Conn) {
	tcp := newTCPConn(c)
	sid := atomic.AddInt32(&s.sessionIdSeq, 1)
	fw := firewall.NewFirewall()
	ctx := NewSocketCmdContext(tcp, sid, fw)

	// 注册到 broadcaster
	broadcaster.AddCmdCtx(ctx)
	if s.onConnect != nil {
		cb := s.onConnect
		main_thread.Process(func() { cb(ctx) })
	}

	// sendLoop 在单独 goroutine 运行
	sendDone := make(chan struct{})
	go s.sendLoop(ctx, sendDone)

	// readLoop 在当前 goroutine 运行；
	// 退出后由 Disconnect 关闭 conn + sendQ，使 sendLoop 一同退出。
	s.readLoop(ctx)
	ctx.Disconnect()
	<-sendDone

	// 主线程内统一处理 onDisconnect + 移除 broadcaster
	main_thread.Process(func() {
		if s.onDisconnect != nil {
			s.onDisconnect(ctx)
		}
		broadcaster.RemoveCmdCtxBySessionId(sid)
	})
}

// readLoop 持续读帧并派发到主线程。
func (s *Server) readLoop(ctx *SocketCmdContext) {
	fr := codec.NewFrameReader(ctx.conn.br)
	for {
		frame, err := fr.ReadFrame()
		if err != nil {
			if !errors.Is(err, io.EOF) && !isClosedErr(err) {
				log.Error("[socket] read err sid=%d: %v", ctx.sessionId, err)
			}
			return
		}

		// 防火墙
		if !ctx.firewall.Check(ctx, frame) {
			ctx.Disconnect()
			return
		}

		// 派发策略：
		// - 已登录（userId != 0）走 async_op 按 userId 分片：跨用户并行，同用户串行
		// - 未登录（登录/注册等流程）走 main_thread 串行：避免伪造 userId 攻击分片
		cmd := frame.Cmd
		cp := frame
		uid := ctx.GetUserId()
		if uid != 0 {
			async_op.Process(int(uid), func() {
				ok, hErr := handler.Dispatch(ctx, cp)
				if !ok {
					log.Info("[socket] unknown cmd=%d sid=%d", cmd, ctx.sessionId)
					return
				}
				if hErr != nil {
					log.Error("[socket] handler err cmd=%d sid=%d: %v", cmd, ctx.sessionId, hErr)
				}
			}, nil)
		} else {
			main_thread.Process(func() {
				ok, hErr := handler.Dispatch(ctx, cp)
				if !ok {
					log.Info("[socket] unknown cmd=%d sid=%d", cmd, ctx.sessionId)
					return
				}
				if hErr != nil {
					log.Error("[socket] handler err cmd=%d sid=%d: %v", cmd, ctx.sessionId, hErr)
				}
			})
		}
	}
}

// sendLoop 持续把 sendQ 中的字节切片 flush 到底层连接。
// 关闭 sendQ 或底层连接断开都会退出。
func (s *Server) sendLoop(ctx *SocketCmdContext, done chan struct{}) {
	defer close(done)
	bw := ctx.conn.bw
	for {
		select {
		case data, ok := <-ctx.sendQ:
			if !ok {
				_ = bw.Flush()
				return
			}
			if _, err := bw.Write(data); err != nil {
				return
			}
			if err := bw.Flush(); err != nil {
				return
			}
		}
	}
}

// isClosedErr 判断错误是否是"连接已关闭"类，避免日志噪音
func isClosedErr(err error) bool {
	if err == nil {
		return false
	}
	if errors.Is(err, net.ErrClosed) {
		return true
	}
	if errors.Is(err, io.EOF) {
		return true
	}
	return false
}
