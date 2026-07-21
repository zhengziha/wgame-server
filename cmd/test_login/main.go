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

	fmt.Println("连接成功，开始测试登录流程...")

	testCmdAccount(conn, "18071859875", "test")
}

func testCmdAccount(conn net.Conn, account, password string) {
	fmt.Println("\n=== 测试 CMD_L_ACCOUNT (9040) ===")

	writer := codec.NewGameWriter(200)
	writer.WriteString("cmd_l_login")
	writer.WriteString(account)
	writer.WriteString(password)
	writer.WriteString("00:00:00:00:00:00")
	writer.WriteString("")
	writer.WriteString("")
	writer.WriteString("")
	writer.WriteString("1")
	writer.WriteUByte(0)
	writer.WriteString("")
	writer.WriteString("")
	writer.WriteString("")
	writer.WriteString("")
	writer.WriteString("")
	writer.WriteUByte(0)
	writer.WriteUByte(0)
	writer.WriteString("")

	payload := writer.Bytes()
	frameData, err := codec.EncodeFrame(0, 0, 9040, payload)
	if err != nil {
		fmt.Printf("构建帧失败: %v\n", err)
		return
	}

	fmt.Printf("发送数据: cmd=9040, len=%d\n", len(frameData))
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

		if cmd == 9041 {
			fmt.Println("--- MSG_AGENT_RESULT (认证成功) ---")
			reader := codec.NewGameReader(buf[12:n])

			result, _ := reader.ReadInt()
			id, _ := reader.ReadInt()
			privilege, _ := reader.ReadInt()
			ip, _ := reader.ReadString()
			port, _ := reader.ReadInt()
			serverName, _ := reader.ReadString()
			serverStatus, _ := reader.ReadInt()
			msg, _ := reader.ReadString()

			fmt.Printf("Result=%d ID=%d Privilege=%d\n", result, id, privilege)
			fmt.Printf("IP=%s Port=%d ServerName=%s\n", ip, port, serverName)
			fmt.Printf("ServerStatus=%d Msg=%s\n", serverStatus, msg)

			testCmdLogin(conn, account, "神经玩家")
		} else if cmd == 9042 {
			fmt.Println("--- MSG_AUTH (认证失败) ---")
			reader := codec.NewGameReader(buf[12:n])
			msg, _ := reader.ReadString()
			fmt.Printf("Msg=%s\n", msg)
		}
	}
}

func testCmdLogin(conn net.Conn, user, charName string) {
	fmt.Println("\n=== 测试 CMD_LOGIN (12290) ===")

	writer := codec.NewGameWriter(200)
	writer.WriteString(user)
	writer.WriteInt(12345)
	writer.WriteInt(67890)
	writer.WriteUByte(0)
	writer.WriteUByte(0)
	writer.WriteString("1.0.0")
	writer.WriteString("")
	writer.WriteShort(0)
	writer.WriteUByte(0)
	writer.WriteString("")
	writer.WriteString("")
	writer.WriteUByte(0)

	payload := writer.Bytes()
	frameData, err := codec.EncodeFrame(0, 0, 12290, payload)
	if err != nil {
		fmt.Printf("构建帧失败: %v\n", err)
		return
	}

	fmt.Printf("发送数据: cmd=12290, len=%d\n", len(frameData))
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

		if cmd == 61538 {
			fmt.Println("--- MSG_EXISTED_CHAR_LIST (角色列表) ---")
			reader := codec.NewGameReader(buf[12:n])

			accountOnline, _ := reader.ReadInt()
			charCount, _ := reader.ReadUShort()

			fmt.Printf("AccountOnline=%d CharCount=%d\n", accountOnline, charCount)

			for i := uint16(0); i < charCount; i++ {
				charID, _ := reader.ReadInt()
				name, _ := reader.ReadString()
				level, _ := reader.ReadInt()
				polar, _ := reader.ReadInt()
				sex, _ := reader.ReadInt()
				_, _ = reader.ReadInt() // onlineState
				_, _ = reader.ReadInt() // fashionIcon
				_, _ = reader.ReadInt() // upgradeLevel
				_, _ = reader.ReadInt() // petIcon
				_, _ = reader.ReadInt() // mountIcon
				_, _ = reader.ReadInt() // specialIcon
				_, _ = reader.ReadInt() // genchongIcon
				_, _ = reader.ReadInt() // upgradeType
				_, _ = reader.ReadInt() // nice
				_, _ = reader.ReadInt() // weeklyLoginDays
				_, _ = reader.ReadInt() // isFeisheng
				_, _ = reader.ReadInt() // tao
				gid, _ := reader.ReadString()
				mapID, _ := reader.ReadInt()
				mapName, _ := reader.ReadString()
				_, _ = reader.ReadInt() // line
				x, _ := reader.ReadInt()
				y, _ := reader.ReadInt()
				_, _ = reader.ReadString() // partyName
				_, _ = reader.ReadString() // family
				_, _ = reader.ReadString() // title

				fmt.Printf("角色%d: id=%d name=%s level=%d polar=%d sex=%d\n", i+1, charID, name, level, polar, sex)
				fmt.Printf("  gid=%s map=%d(%s) x=%d y=%d\n", gid, mapID, mapName, x, y)
			}

			testCmdLoadExistedChar(conn, charName)
		}
	}
}

func testCmdLoadExistedChar(conn net.Conn, charName string) {
	fmt.Println("\n=== 测试 CMD_LOAD_EXISTED_CHAR (4192) ===")

	writer := codec.NewGameWriter(100)
	writer.WriteString(charName)

	payload := writer.Bytes()
	frameData, err := codec.EncodeFrame(0, 0, 4192, payload)
	if err != nil {
		fmt.Printf("构建帧失败: %v\n", err)
		return
	}

	fmt.Printf("发送数据: cmd=4192, len=%d\n", len(frameData))
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

		if cmd == 13142 {
			fmt.Println("--- MSG_KICK_OFF ---")
			reader := codec.NewGameReader(buf[12:n])
			msg, _ := reader.ReadString()
			fmt.Printf("Msg=%s\n", msg)
		} else if cmd == 16416 {
			fmt.Println("--- MSG_MAP_INFO ---")
			reader := codec.NewGameReader(buf[12:n])
			mapID, _ := reader.ReadInt()
			mapName, _ := reader.ReadString()
			fmt.Printf("MapID=%d MapName=%s\n", mapID, mapName)
		} else if cmd == 16392 {
			fmt.Println("--- MSG_APPEAR ---")
			reader := codec.NewGameReader(buf[12:n])
			charID, _ := reader.ReadInt()
			name, _ := reader.ReadString()
			gid, _ := reader.ReadString()
			level, _ := reader.ReadInt()
			polar, _ := reader.ReadInt()
			sex, _ := reader.ReadInt()
			x, _ := reader.ReadInt()
			y, _ := reader.ReadInt()
			dir, _ := reader.ReadInt()
			waiguan, _ := reader.ReadInt()
			nice, _ := reader.ReadInt()
			fashionIcon, _ := reader.ReadInt()
			fmt.Printf("CharID=%d Name=%s Gid=%s Level=%d\n", charID, name, gid, level)
			fmt.Printf("Polar=%d Sex=%d X=%d Y=%d Dir=%d\n", polar, sex, x, y, dir)
			fmt.Printf("Waiguan=%d Nice=%d FashionIcon=%d\n", waiguan, nice, fashionIcon)
		}
	}

	n, err = conn.Read(buf)
	if err != nil {
		fmt.Printf("读取第二条消息失败: %v\n", err)
		return
	}

	fmt.Printf("\n收到第二条响应: len=%d\n", n)
	if n >= 12 {
		cmd := uint16(buf[10])<<8 | uint16(buf[11])
		fmt.Printf("响应CMD: %d\n", cmd)

		if cmd == 16416 {
			fmt.Println("--- MSG_MAP_INFO ---")
			reader := codec.NewGameReader(buf[12:n])
			mapID, _ := reader.ReadInt()
			mapName, _ := reader.ReadString()
			fmt.Printf("MapID=%d MapName=%s\n", mapID, mapName)
		} else if cmd == 16392 {
			fmt.Println("--- MSG_APPEAR ---")
			reader := codec.NewGameReader(buf[12:n])
			charID, _ := reader.ReadInt()
			name, _ := reader.ReadString()
			gid, _ := reader.ReadString()
			level, _ := reader.ReadInt()
			polar, _ := reader.ReadInt()
			sex, _ := reader.ReadInt()
			x, _ := reader.ReadInt()
			y, _ := reader.ReadInt()
			dir, _ := reader.ReadInt()
			waiguan, _ := reader.ReadInt()
			nice, _ := reader.ReadInt()
			fashionIcon, _ := reader.ReadInt()
			fmt.Printf("CharID=%d Name=%s Gid=%s Level=%d\n", charID, name, gid, level)
			fmt.Printf("Polar=%d Sex=%d X=%d Y=%d Dir=%d\n", polar, sex, x, y, dir)
			fmt.Printf("Waiguan=%d Nice=%d FashionIcon=%d\n", waiguan, nice, fashionIcon)

			testCmdEcho(conn)
		}
	}
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
		}
	}
}
