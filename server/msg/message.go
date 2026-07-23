// Package msg 定义入站/出站消息抽象与帧编解码入口。
//
// 设计要点：
//   - InMessage / OutMessage 仅要求实现 Cmd()；字段读写完全由反射自动完成
//   - 消息结构体字段按 Java 端写入顺序定义，字段类型用 `codec` tag 指定编码格式
//   - 复杂消息（如含 list 结构）使用 `codec:"list:..."` tag 声明
//   - WriteFrame 把 OutMessage 编码为完整出站帧（含 10 字节头 + cmd + payload + 加密）
//
// 反射编解码底层位于 codec 包，支持的 tag 类型：
//
//	bool, ubyte, short, ushort, int, uint, long, ulong,
//	float, double, string, string2, string4, bytes,
//	list:byte, list:short, list:ushort, list:int, list:uint
package msg

import (
	"wgame-server/server/codec"
)

// InMessage 入站消息抽象（客户端发来的请求体）。
// 仅要求实现 Cmd()；字段读取完全由反射自动完成。
type InMessage interface {
	Cmd() uint16
}

// OutMessage 出站消息抽象，对应 Java wd-server-fl 中的 BaseWrite 抽象类。
// 仅要求实现 Cmd()；字段写入完全由反射自动完成。
type OutMessage interface {
	Cmd() uint16
}

// WriteFrame 把 OutMessage 编码为完整出站帧（含加密）。
// 字段写入完全通过反射按结构体字段声明顺序自动完成。
//
// tableIndex:
//   - 0  表示不加密
//   - >0 表示使用指定加密表
//
// tickCount 通常传 0（与 Java 服务端实现一致）。
func WriteFrame(m OutMessage, tableIndex int, tickCount int32) ([]byte, error) {
	w := codec.NewGameWriter(64)
	if err := codec.AutoWrite(w, m); err != nil {
		return nil, err
	}
	payload := w.Bytes()

	if tableIndex < 0 {
		tableIndex = 0
	}
	return codec.EncodeFrame(tableIndex, tickCount, m.Cmd(), payload)
}

// ReadFramePayload 把帧 payload（不含 cmd）反序列化到 m。
// 字段读取完全通过反射按结构体字段声明顺序自动完成。
func ReadFramePayload(m InMessage, r *codec.GameReader) error {
	return codec.AutoRead(r, m)
}
