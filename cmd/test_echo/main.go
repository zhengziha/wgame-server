package main

import (
	"fmt"
	"net"

	"wgame-server/server/codec"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8800")
	if err != nil {
		fmt.Printf("连接失败: %v\n", err)
		return
	}
	defer conn.Close()

	fmt.Println("连接成功，测试心跳检测...")

	testCmdEcho(conn)
}

func testCmdEcho(conn net.Conn) {
	fmt.Println("\n=== 测试 CMD_ECHO (4274) ===")

	currentTime := int32(1000000)
	peerTime := int32(500000)

	writer := codec.NewGameWriter(50)
	writer.WriteInt(currentTime)
	writer.WriteInt(peerTime)

	payload := writer.Bytes()
	frameData, err := codec.EncodeFrame(0, 0, 4274, payload)
	if err != nil {
		fmt.Printf("构建帧失败: %v\n", err)
		return
	}

	fmt.Printf("发送数据: cmd=4274, len=%d, currentTime=%d, peerTime=%d\n", len(frameData), currentTime, peerTime)
	conn.Write(frameData)

	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Printf("读取失败: %v\n", err)
		return
	}

	fmt.Printf("收到响应: len=%d\n", n)

	if n >= 12 {
		cmd := uint16(buf[10])<<8 | uint16(buf[11])
		fmt.Printf("响应CMD: %d\n", cmd)

		if cmd == 4275 {
			fmt.Println("--- MSG_REPLY_ECHO ---")
			reader := codec.NewGameReader(buf[12:n])
			a, _ := reader.ReadInt()
			fmt.Printf("A=%d (服务器时间，含随机延迟)\n", a)
			fmt.Printf("延迟范围: %d - %d ms\n", a-peerTime-1500, a-peerTime-100)
		} else if cmd == 13142 {
			fmt.Println("--- MSG_KICK_OFF ---")
			reader := codec.NewGameReader(buf[12:n])
			msg, _ := reader.ReadString()
			fmt.Printf("Msg=%s\n", msg)
		}
	}
}
