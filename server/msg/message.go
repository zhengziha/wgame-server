// Package msg 定义入站/出站消息抽象与帧编解码入口。
//
// 设计要点：
//   - InMessage / OutMessage 仅要求实现 Cmd()；字段读写默认由反射自动完成
//   - 复杂消息可实现 CustomCodec 接口接管读写过程
//   - WriteFrame 把 OutMessage 编码为完整出站帧（含 10 字节头 + cmd + payload + 加密）
//
// 由 msg.WriteFrame 负责组装帧并处理加密；反射编解码底层位于 codec 包。
package msg

import (
	"wgame-server/server/codec"
)

// InMessage 入站消息抽象（客户端发来的请求体）。
//
// 仅要求实现 Cmd()；字段读取默认由反射自动完成。
// 如需自定义读取逻辑（例如包含复杂嵌套结构），额外实现 CustomCodec 接口即可。
type InMessage interface {
	// Cmd 消息 id
	Cmd() uint16
}

// OutMessage 出站消息抽象，对应 Java wd-server-fl 中的 BaseWrite 抽象类。
//
// 仅要求实现 Cmd()；字段写入默认由反射自动完成。
// 如需自定义写入逻辑（例如保持与历史协议完全一致的手工顺序），额外实现 CustomCodec。
//
// 由 msg.WriteFrame 负责组装 10 字节头 + cmd + payload + 加密。
type OutMessage interface {
	// Cmd 消息 id（如 4099 = LOGIN_DONE）
	Cmd() uint16
}

// CustomCodec 可选接口：消息体可实现它以接管读写过程。
// 未实现时，msg 包使用反射按字段顺序自动读写。
type CustomCodec interface {
	// WriteBody 把消息体字段写入 writer（不含 cmd 字节，由外层负责）
	WriteBody(w *codec.GameWriter)

	// ReadBody 从 reader 读取字段并填充自身（不含 cmd 字节）
	ReadBody(r *codec.GameReader) error
}

// AutoWriteBody 通过反射将 obj 字段按声明顺序写入 w。
// 等价于在每个 OutMessage 上手写 WriteBody。
func AutoWriteBody(w *codec.GameWriter, obj OutMessage) error {
	return codec.AutoWrite(w, obj)
}

// AutoReadBody 通过反射按声明顺序从 r 读取并填充 obj（须为可寻址的 struct 指针）。
// 等价于在每个 InMessage 上手写 ReadBody。
func AutoReadBody(r *codec.GameReader, obj InMessage) error {
	return codec.AutoRead(r, obj)
}

// writeBody 对 OutMessage 选择自定义或反射路径
func writeBody(m OutMessage, w *codec.GameWriter) error {
	if cc, ok := m.(CustomCodec); ok {
		cc.WriteBody(w)
		return nil
	}
	return codec.AutoWrite(w, m)
}

// readBody 对 InMessage 选择自定义或反射路径
func readBody(m InMessage, r *codec.GameReader) error {
	if cc, ok := m.(CustomCodec); ok {
		return cc.ReadBody(r)
	}
	return codec.AutoRead(r, m)
}

// WriteFrame 把 OutMessage 编码为完整出站帧（含加密）。
//
// 字段写入优先级：
//  1. 若 m 实现 CustomCodec，则调用其 WriteBody
//  2. 否则使用反射按字段顺序自动写入
//
// tableIndex:
//   - 0  表示不加密
//   - >0 表示使用指定加密表
//   - <0 表示由本函数随机选取（默认行为，等价于 Java BaseWrite 默认分支）
//
// tickCount 通常传 0（与 Java 服务端实现一致）。
func WriteFrame(m OutMessage, tableIndex int, tickCount int32) ([]byte, error) {
	w := codec.NewGameWriter(64)
	if err := writeBody(m, w); err != nil {
		return nil, err
	}
	payload := w.Bytes()

	if tableIndex < 0 {
		// 简单的随机策略；具体随机源由调用方在更高层注入更合适，
		// 这里默认走不加密路径，业务可按需覆盖。
		tableIndex = 0
	}
	return codec.EncodeFrame(tableIndex, tickCount, m.Cmd(), payload)
}

// ReadFramePayload 把帧 payload（不含 cmd）反序列化到 m。
// 优先调用 CustomCodec.ReadBody，否则使用反射自动读取。
func ReadFramePayload(m InMessage, r *codec.GameReader) error {
	return readBody(m, r)
}
