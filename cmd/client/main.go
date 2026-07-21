// Package main 是 wgame-server 的简易测试客户端，
// 用于验证自定义 socket 协议端到端可用性。
//
// 用法：
//
//	go run ./cmd/client -addr=127.0.0.1:8800 -text=hello
//
// 行为：
//  1. 建立 TCP 连接
//  2. 发送一个 EchoReq 帧（cmd=0x0101，body=1 字节长度前缀 + GBK 字符串）
//  3. 等待并解析服务端回包（cmd=0x0001 的 EchoMsg）
//  4. 与发送文本比对，输出 PASS / FAIL
package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"time"

	"wgame-server/server/codec"
)

// CmdEchoReq 与 server/demo/handlers 中定义的回显请求 cmd 保持一致
const CmdEchoReq uint16 = 0x0101

func main() {
	addr := flag.String("addr", "127.0.0.1:8800", "server address")
	text := flag.String("text", "hello", "text to echo")
	flag.Parse()

	// 1) 建立连接
	conn, err := net.DialTimeout("tcp", *addr, 5*time.Second)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[client] dial %s failed: %v\n", *addr, err)
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Printf("[client] connected to %s\n", *addr)

	// 2) 构造请求帧
	//    body 布局与 GameWriter.WriteString 一致：1 字节长度前缀 + GBK 字节
	w := codec.NewGameWriter(64)
	w.WriteString(*text)
	payload := w.Bytes()

	//    整帧：tableIndex=0 表示不加密（与服务端 msg.WriteFrame 默认行为一致）
	frameBytes, err := codec.EncodeFrame(0, 0, CmdEchoReq, payload)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[client] encode frame failed: %v\n", err)
		os.Exit(1)
	}

	if _, err := conn.Write(frameBytes); err != nil {
		fmt.Fprintf(os.Stderr, "[client] write frame failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("[client] sent cmd=0x%04X text=%q (%d bytes)\n",
		CmdEchoReq, *text, len(frameBytes))

	// 3) 读取响应帧
	//    设置读超时，避免服务端异常无响应时永久阻塞
	_ = conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	fr := codec.NewFrameReader(bufio.NewReader(conn))
	resp, err := fr.ReadFrame()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[client] read frame failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("[client] recv cmd=0x%04X tableIndex=%d tickCount=%d bodyLen=%d\n",
		resp.Cmd, resp.TableIndex, resp.TickCount, len(resp.Body))

	// 4) 解析 body（与 server/demo/msg.EchoMsg.WriteBody 对应：一个 WriteString）
	r := codec.NewGameReader(resp.Body)
	echoText, err := r.ReadString()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[client] parse body failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("[client] echo text=%q\n", echoText)

	// 5) 简单断言
	if echoText == *text {
		fmt.Println("[client] RESULT: PASS")
	} else {
		fmt.Println("[client] RESULT: MISMATCH")
		os.Exit(2)
	}
}
