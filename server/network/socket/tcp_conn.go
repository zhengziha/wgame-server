package socket

import (
	"bufio"
	"net"
	"sync"
)

// tcpConn 是对 net.Conn 的轻量封装。
// 把 *bufio.Reader / *bufio.Writer 与底层 conn 一起携带，
// 便于 readLoop 直接使用 codec.FrameReader。
//
// close 操作是幂等的。
type tcpConn struct {
	raw    net.Conn
	br     *bufio.Reader
	bw     *bufio.Writer
	id     string // 等价 Java Netty channel.id().asLongText()
	remoteIP string

	closeOnce sync.Once
}

func newTCPConn(c net.Conn) *tcpConn {
	host, _, _ := net.SplitHostPort(c.RemoteAddr().String())
	return &tcpConn{
		raw:      c,
		br:       bufio.NewReaderSize(c, 8192),
		bw:       bufio.NewWriterSize(c, 8192),
		id:       c.RemoteAddr().String(), // 简单起见用对端地址字符串
		remoteIP: host,
	}
}

// close 幂等关闭底层连接
func (t *tcpConn) close() {
	t.closeOnce.Do(func() {
		_ = t.raw.Close()
	})
}
