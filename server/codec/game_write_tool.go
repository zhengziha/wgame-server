package codec

import (
	"encoding/binary"
	"math"
)

// 大端字节写入工具，对应 Java wd-server-fl 中的 GameWriteTool。
// 字符串默认使用 GBK 编码。

// GameWriter 向内部字节缓冲按大端序写入基本类型。
type GameWriter struct {
	buf []byte
}

// NewGameWriter 创建一个新的 writer，预留 cap 字节以减少扩容
func NewGameWriter(cap int) *GameWriter {
	return &GameWriter{buf: make([]byte, 0, cap)}
}

// Bytes 返回已写入的字节切片（指向内部缓冲，调用方不应修改）
func (w *GameWriter) Bytes() []byte {
	return w.buf
}

// Len 返回已写入的字节数
func (w *GameWriter) Len() int {
	return len(w.buf)
}

// WriteUByte 写入 1 字节无符号（参数以 int 传入，仅取低 8 位）
func (w *GameWriter) WriteUByte(b int) {
	w.buf = append(w.buf, byte(b&0xFF))
}

// WriteByte 写入 1 字节有符号（对应 Java GameWriteTool.writeByte）
func (w *GameWriter) WriteByte(b int8) {
	w.buf = append(w.buf, byte(b))
}

// WriteBoolean 写入 1 字节布尔
func (w *GameWriter) WriteBoolean(v bool) {
	if v {
		w.buf = append(w.buf, 1)
	} else {
		w.buf = append(w.buf, 0)
	}
}

// WriteShort 写入 2 字节大端 int16
func (w *GameWriter) WriteShort(v int16) {
	var tmp [2]byte
	binary.BigEndian.PutUint16(tmp[:], uint16(v))
	w.buf = append(w.buf, tmp[:]...)
}

// WriteUShort 写入 2 字节大端 uint16
func (w *GameWriter) WriteUShort(v uint16) {
	var tmp [2]byte
	binary.BigEndian.PutUint16(tmp[:], v)
	w.buf = append(w.buf, tmp[:]...)
}

// WriteInt 写入 4 字节大端 int32
func (w *GameWriter) WriteInt(v int32) {
	var tmp [4]byte
	binary.BigEndian.PutUint32(tmp[:], uint32(v))
	w.buf = append(w.buf, tmp[:]...)
}

// WriteUInt 写入 4 字节大端 uint32
func (w *GameWriter) WriteUInt(v uint32) {
	var tmp [4]byte
	binary.BigEndian.PutUint32(tmp[:], v)
	w.buf = append(w.buf, tmp[:]...)
}

// WriteLong 写入 8 字节大端 int64
func (w *GameWriter) WriteLong(v int64) {
	var tmp [8]byte
	binary.BigEndian.PutUint64(tmp[:], uint64(v))
	w.buf = append(w.buf, tmp[:]...)
}

// WriteFloat 写入 4 字节大端 float32
func (w *GameWriter) WriteFloat(v float32) {
	w.WriteUInt(math.Float32bits(v))
}

// WriteDouble 写入 8 字节大端 float64
func (w *GameWriter) WriteDouble(v float64) {
	w.WriteLong(int64(math.Float64bits(v)))
}

// WriteString 写入 1 字节长度前缀 + GBK 字节
// 注意：GBK 长度超过 255 会触发 panic（与 Java 实现一致）
func (w *GameWriter) WriteString(s string) {
	b, err := gbkEncode(s)
	if err != nil {
		b = []byte(s)
	}
	if len(b) > 0xFF {
		// 与 Java GameWriteTool.writeString 行为一致：直接截断
		b = b[:0xFF]
	}
	w.buf = append(w.buf, byte(len(b)))
	w.buf = append(w.buf, b...)
}

// WriteString2 写入 2 字节长度前缀 + GBK 字节
func (w *GameWriter) WriteString2(s string) {
	b, err := gbkEncode(s)
	if err != nil {
		b = []byte(s)
	}
	w.WriteUShort(uint16(len(b)))
	w.buf = append(w.buf, b...)
}

// WriteString4 写入 4 字节长度前缀 + GBK 字节
func (w *GameWriter) WriteString4(s string) {
	b, err := gbkEncode(s)
	if err != nil {
		b = []byte(s)
	}
	w.WriteUInt(uint32(len(b)))
	w.buf = append(w.buf, b...)
}

// WriteBytes 写入 2 字节长度前缀 + 原始字节
func (w *GameWriter) WriteBytes(b []byte) {
	w.WriteUShort(uint16(len(b)))
	w.buf = append(w.buf, b...)
}

// WriteRaw 写入原始字节（无长度前缀）
func (w *GameWriter) WriteRaw(b []byte) {
	w.buf = append(w.buf, b...)
}

// WriteZero 写入 n 个 0 字节
func (w *GameWriter) WriteZero(n int) {
	for i := 0; i < n; i++ {
		w.buf = append(w.buf, 0)
	}
}
