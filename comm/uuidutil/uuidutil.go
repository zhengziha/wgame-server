// Package uuidutil 提供全局唯一 id（gid）的生成工具。
//
// wd-server-fl 客户端只认 gid，不认数据库自增主键 id。
// gid 的格式与 Java 端保持一致：
//
//	UUID.randomUUID().toString().replace("-", "")
//
// 即 32 位无横杠的十六进制字符串，例如 "550e8400e29b41d4a716446655440000"。
package uuidutil

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

// NewGid 生成一个 32 字符长的无横杠 UUID v4 字符串。
//
// 格式：xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx（32 个十六进制字符）
// 版本：UUID v4（第 13 位是 '4'，第 17 位是 8/9/a/b 之一）
//
// 与 Java 端 UUID.randomUUID().toString().replace("-", "") 等价。
// 使用 crypto/rand 保证随机性。
//
// 失败情况极其罕见（系统熵池耗尽），此时返回错误，调用方应决定是否重试。
func NewGid() (string, error) {
	var b [16]byte
	if _, err := rand.Read(b[:]); err != nil {
		return "", fmt.Errorf("uuidutil: read random bytes failed: %w", err)
	}
	// UUID v4 版本位与变体位设置
	b[6] = (b[6] & 0x0f) | 0x40 // version 4
	b[8] = (b[8] & 0x3f) | 0x80 // variant 10
	return hex.EncodeToString(b[:]), nil
}

// MustNewGid 生成一个 gid，失败时 panic。
// 仅用于程序初始化阶段或确信不会失败的场景；业务路径请用 NewGid。
func MustNewGid() string {
	gid, err := NewGid()
	if err != nil {
		panic(err)
	}
	return gid
}
