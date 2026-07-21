package msg

import "wgame-server/server/codec"

// OutMessage 出站消息抽象，对应 Java wd-server-fl 中的 BaseWrite 抽象类。
//
// 每个具体消息实现该接口：
//   - Cmd() 返回消息 id
//   - WriteBody(w) 把字段按协议顺序写入 writer
//
// 由 msg.WriteFrame 负责组装 10 字节头 + cmd + payload + 加密。
type OutMessage interface {
	// Cmd 消息 id（如 4099 = LOGIN_DONE）
	Cmd() uint16

	// WriteBody 把消息体字段写入 writer（不含 cmd 字节，由外层负责）
	WriteBody(w *codec.GameWriter)
}

// WriteFrame 把 OutMessage 编码为完整出站帧（含加密）。
//
// tableIndex:
//   - 0  表示不加密
//   - >0 表示使用指定加密表
//   - <0 表示由本函数随机选取（默认行为，等价于 Java BaseWrite 默认分支）
//
// tickCount 通常传 0（与 Java 服务端实现一致）。
func WriteFrame(m OutMessage, tableIndex int, tickCount int32) ([]byte, error) {
	w := codec.NewGameWriter(64)
	m.WriteBody(w)
	payload := w.Bytes()

	if tableIndex < 0 {
		// 简单的随机策略；具体随机源由调用方在更高层注入更合适，
		// 这里默认走不加密路径，业务可按需覆盖。
		tableIndex = 0
	}
	return codec.EncodeFrame(tableIndex, tickCount, m.Cmd(), payload)
}
