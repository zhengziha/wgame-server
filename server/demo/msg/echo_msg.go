package demo

import (
	"wgame-server/server/codec"
)

// EchoMsg 演示用的回显消息。
// 协议：cmd=0x0001（ECHO_DONE），body = 原样回写的字符串（GBK + 4 字节长度头）。
const CmdEcho uint16 = 0x0001

type EchoMsg struct {
	Text string
}

func (m *EchoMsg) Cmd() uint16 { return CmdEcho }

func (m *EchoMsg) WriteBody(w *codec.GameWriter) {
	w.WriteString(m.Text)
}
