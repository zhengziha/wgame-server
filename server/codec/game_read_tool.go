package codec

import (
	"encoding/binary"
	"errors"
)

// 大端字节读写工具，对应 Java wd-server-fl 中的 GameReadTool。
// 字符串默认使用 GBK 编码（与原项目一致）。

// GameReader 从给定的字节切片中按大端序读取基本类型。
type GameReader struct {
	buf  []byte
	pos  int
}

// ErrShortBuffer 表示可读字节不足以完成读取
var ErrShortBuffer = errors.New("game reader: not enough bytes")

// NewGameReader 创建一个新的读取器
func NewGameReader(buf []byte) *GameReader {
	return &GameReader{buf: buf}
}

// Remaining 返回剩余可读字节数
func (r *GameReader) Remaining() int {
	return len(r.buf) - r.pos
}

// ReadInt 读取 4 字节大端有符号 int32
func (r *GameReader) ReadInt() (int32, error) {
	if r.Remaining() < 4 {
		return 0, ErrShortBuffer
	}
	v := int32(binary.BigEndian.Uint32(r.buf[r.pos:]))
	r.pos += 4
	return v, nil
}

// ReadUInt 读取 4 字节大端无符号 uint32
func (r *GameReader) ReadUInt() (uint32, error) {
	if r.Remaining() < 4 {
		return 0, ErrShortBuffer
	}
	v := binary.BigEndian.Uint32(r.buf[r.pos:])
	r.pos += 4
	return v, nil
}

// ReadShort 读取 2 字节大端有符号 int16
func (r *GameReader) ReadShort() (int16, error) {
	if r.Remaining() < 2 {
		return 0, ErrShortBuffer
	}
	v := int16(binary.BigEndian.Uint16(r.buf[r.pos:]))
	r.pos += 2
	return v, nil
}

// ReadUShort 读取 2 字节大端无符号 uint16
func (r *GameReader) ReadUShort() (uint16, error) {
	if r.Remaining() < 2 {
		return 0, ErrShortBuffer
	}
	v := binary.BigEndian.Uint16(r.buf[r.pos:])
	r.pos += 2
	return v, nil
}

// ReadLong 读取 8 字节大端有符号 int64
func (r *GameReader) ReadLong() (int64, error) {
	if r.Remaining() < 8 {
		return 0, ErrShortBuffer
	}
	v := int64(binary.BigEndian.Uint64(r.buf[r.pos:]))
	r.pos += 8
	return v, nil
}

// ReadUByte 读取 1 字节（返回 int 形式，与 Java readUnsignedByte 对应）
func (r *GameReader) ReadUByte() (int, error) {
	if r.Remaining() < 1 {
		return 0, ErrShortBuffer
	}
	v := int(r.buf[r.pos])
	r.pos++
	return v, nil
}

// ReadByte 读取 1 字节有符号（返回 int8，与 Java GameReadTool.readByte 对应）
func (r *GameReader) ReadByte() (int8, error) {
	if r.Remaining() < 1 {
		return 0, ErrShortBuffer
	}
	v := int8(r.buf[r.pos])
	r.pos++
	return v, nil
}

// ReadBoolean 读取 1 字节布尔（0/1）
func (r *GameReader) ReadBoolean() (bool, error) {
	b, err := r.ReadUByte()
	if err != nil {
		return false, err
	}
	return b != 0, nil
}

// ReadString 读取 1 字节长度前缀的 GBK 字符串
func (r *GameReader) ReadString() (string, error) {
	n, err := r.ReadUByte()
	if err != nil {
		return "", err
	}
	return r.readGBK(n)
}

// ReadString2 读取 2 字节长度前缀的 GBK 字符串
func (r *GameReader) ReadString2() (string, error) {
	n, err := r.ReadUShort()
	if err != nil {
		return "", err
	}
	return r.readGBK(int(n))
}

// ReadString4 读取 4 字节长度前缀的 GBK 字符串
func (r *GameReader) ReadString4() (string, error) {
	n, err := r.ReadUInt()
	if err != nil {
		return "", err
	}
	return r.readGBK(int(n))
}

// ReadBytes 读取 2 字节长度前缀 + 原始字节切片
func (r *GameReader) ReadBytes() ([]byte, error) {
	n, err := r.ReadUShort()
	if err != nil {
		return nil, err
	}
	if r.Remaining() < int(n) {
		return nil, ErrShortBuffer
	}
	out := make([]byte, n)
	copy(out, r.buf[r.pos:r.pos+int(n)])
	r.pos += int(n)
	return out, nil
}

func (r *GameReader) readGBK(n int) (string, error) {
	if n < 0 || r.Remaining() < n {
		return "", ErrShortBuffer
	}
	s, err := gbkDecode(r.buf[r.pos : r.pos+n])
	r.pos += n
	return s, err
}

// Skip 跳过 n 个字节
func (r *GameReader) Skip(n int) error {
	if n < 0 || r.Remaining() < n {
		return ErrShortBuffer
	}
	r.pos += n
	return nil
}
