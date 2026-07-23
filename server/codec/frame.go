package codec

import (
	"encoding/binary"
	"errors"
	"io"
)

// 帧/协议头定义，严格对应 Java wd-server-fl 的 ServerChannelInitializer + BaseWrite 协议。
//
// 完整帧字节布局（共 10 字节头 + 变长 body）：
//
//	偏移  长度  字段         说明
//	0     2    magic       uint16 BE = 19802 (0x4D6A)
//	2     2    tableIndex  uint16 BE (1..35 = 加密表索引；0 = 不加密)
//	4     4    tickCount   int32  BE (C->S 客户端 tick；S->C 服务端写 0)
//	8     2    length      uint16 BE = 2 + len(payload)，即从 cmd 开始到末尾的字节数
//	10    2    cmd         uint16 BE
//	12    ...  payload     由具体消息决定
//
// 加密范围：bytes[10 : 10+length]（包含 cmd 与 payload），头部 10 字节明文。
// 服务端发送时会做替换加密；客户端到服务端的 cmd/payload 当前为明文（与 Java 实现一致）。
const (
	// Magic 协议魔数
	Magic uint16 = 19802
	// HeaderSize 协议头大小
	HeaderSize = 10
	// MaxFrameLength 单帧最大字节数（与 Java LengthFieldBasedFrameDecoder 一致）
	MaxFrameLength = 10240
	// LengthFieldOffset length 字段在头中的偏移
	LengthFieldOffset = 8
	// LengthFieldLength length 字段本身的长度
	LengthFieldLength = 2
)

// ErrFrameTooLong 单帧超出 MaxFrameLength
var ErrFrameTooLong = errors.New("codec: frame too long")

// ErrBadMagic 魔数错误
var ErrBadMagic = errors.New("codec: bad magic")

// Frame 表示解出来的一个完整帧。
// 解码后 Body 不做解密（与原 Java ServerHandler 行为一致：入站不 decrypt）。
type Frame struct {
	TableIndex uint16
	TickCount  int32
	Cmd        uint16
	Body       []byte // 不含 cmd，仅 payload
}

// NO_ENC_CMDs Java端BaseWrite中定义的不加密CMD列表
// 对应Java: public static final String NO_ENC = "45143#45555#13143";
var NO_ENC_CMDs = map[uint16]bool{
	45143: true, // MSG_L_WAIT_IN_LINE
	45555: true, // MSG_L_START_LOGIN
	13143: true, // MSG_L_AGENT_RESULT
}

// EncodeFrame 组装一个完整的出站帧字节切片（含加密）。
//
// 参数：
//
//	tableIndex 0 表示不加密；否则范围 1..35
//	tickCount  写入头部的 tickCount 字段（C->S 含义；S->C 通常传 0）
//	cmd        消息 id
//	payload    消息体字节（不含 cmd）
func EncodeFrame(tableIndex int, tickCount int32, cmd uint16, payload []byte) ([]byte, error) {
	bodyLen := 2 + len(payload) // cmd + payload
	if HeaderSize+bodyLen > MaxFrameLength {
		return nil, ErrFrameTooLong
	}

	buf := make([]byte, 0, HeaderSize+bodyLen)
	var hdr [HeaderSize]byte

	if NO_ENC_CMDs[cmd] {
		binary.BigEndian.PutUint16(hdr[0:2], Magic)
		binary.BigEndian.PutUint16(hdr[2:4], 0)          // tableIndex=0
		binary.BigEndian.PutUint32(hdr[4:8], uint32(0)) // tickCount=0
	} else {
		binary.BigEndian.PutUint16(hdr[0:2], Magic)
		binary.BigEndian.PutUint16(hdr[2:4], uint16(tableIndex))
		binary.BigEndian.PutUint32(hdr[4:8], uint32(tickCount))
	}
	binary.BigEndian.PutUint16(hdr[8:10], uint16(bodyLen))
	buf = append(buf, hdr[:]...)

	// body 区：cmd + payload
	var cmdBytes [2]byte
	binary.BigEndian.PutUint16(cmdBytes[:], cmd)
	buf = append(buf, cmdBytes[:]...)
	buf = append(buf, payload...)

	// 加密 body 区（含 cmd）
	if !NO_ENC_CMDs[cmd] && tableIndex >= 1 && tableIndex <= TableSize {
		EncryptBody(tableIndex, buf[HeaderSize:])
	}
	return buf, nil
}

// FrameReader 实现按帧从底层 io.Reader 中读取完整帧。
// 等价于 Java Netty 的 LengthFieldBasedFrameDecoder(10240, 8, 2, 0, 4)
// 但本实现不做 initialBytesToStrip，而是返回完整帧（含头部），
// 由调用方使用 DecodeFrame 解析头部字段。
type FrameReader struct {
	r       io.Reader
	hdrBuf  [HeaderSize]byte
	bodyBuf []byte
}

// NewFrameReader 创建一个帧读取器
func NewFrameReader(r io.Reader) *FrameReader {
	return &FrameReader{r: r}
}

// ReadFrame 阻塞读取一个完整帧；返回的 Frame 中 Body 仅包含 payload（不含 cmd）。
// 当底层连接关闭时返回 io.EOF。
func (fr *FrameReader) ReadFrame() (*Frame, error) {
	// 读 10 字节头
	if _, err := io.ReadFull(fr.r, fr.hdrBuf[:]); err != nil {
		return nil, err
	}
	magic := binary.BigEndian.Uint16(fr.hdrBuf[0:2])
	if magic != Magic {
		return nil, ErrBadMagic
	}
	tableIndex := binary.BigEndian.Uint16(fr.hdrBuf[2:4])
	tickCount := int32(binary.BigEndian.Uint32(fr.hdrBuf[4:8]))
	bodyLen := int(binary.BigEndian.Uint16(fr.hdrBuf[8:10]))
	if bodyLen < 2 {
		return nil, errors.New("codec: body length < 2")
	}
	if HeaderSize+bodyLen > MaxFrameLength {
		return nil, ErrFrameTooLong
	}
	if cap(fr.bodyBuf) < bodyLen {
		fr.bodyBuf = make([]byte, bodyLen)
	} else {
		fr.bodyBuf = fr.bodyBuf[:bodyLen]
	}
	if _, err := io.ReadFull(fr.r, fr.bodyBuf); err != nil {
		return nil, err
	}
	cmd := binary.BigEndian.Uint16(fr.bodyBuf[0:2])
	// 复制 payload，避免持有底层缓冲
	payload := make([]byte, bodyLen-2)
	copy(payload, fr.bodyBuf[2:bodyLen])
	return &Frame{
		TableIndex: tableIndex,
		TickCount:  tickCount,
		Cmd:        cmd,
		Body:       payload,
	}, nil
}
