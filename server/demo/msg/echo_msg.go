package msg

// EchoMsg 演示用的回显消息。
// 协议：cmd=0x0001（ECHO_DONE），body = 原样回写的字符串（GBK + 1 字节长度头）。
const CmdEcho uint16 = 0x0001

type EchoMsg struct {
	Text string
}

func (m *EchoMsg) Cmd() uint16 { return CmdEcho }

// EchoReq 客户端发来的回显请求，cmd=0x0101
// 入站 body: WriteString(text)
const CmdEchoReq uint16 = 0x0101

type EchoReq struct {
	Text string
}

func (m *EchoReq) Cmd() uint16 { return CmdEchoReq }
