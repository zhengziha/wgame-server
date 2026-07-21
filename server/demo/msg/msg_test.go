package msg_test

import (
	"bytes"
	"testing"

	"wgame-server/server/codec"
	demomsg "wgame-server/server/demo/msg"
)

// TestUserQueryRespAutoWriteMatchesManual 验证 UserQueryResp 的反射自动写入
// 与原始手写 WriteBody 字节级等价，避免协议回归。
//
// 原手写顺序：
//
//	w.WriteBoolean(m.Found)
//	w.WriteLong(m.ID)
//	w.WriteString(m.Account)
//	w.WriteString(m.Nickname)
//	w.WriteInt(int32(m.Level))
func TestUserQueryRespAutoWriteMatchesManual(t *testing.T) {
	m := &demomsg.UserQueryResp{
		Found:    true,
		ID:       0x0123456789ABCDEF,
		Account:  "hero",
		Nickname: "猎魔人",
		Level:    42,
	}

	want := codec.NewGameWriter(64)
	want.WriteBoolean(m.Found)
	want.WriteLong(m.ID)
	want.WriteString(m.Account)
	want.WriteString(m.Nickname)
	want.WriteInt(int32(m.Level))

	got := codec.NewGameWriter(64)
	if err := codec.AutoWrite(got, m); err != nil {
		t.Fatalf("AutoWrite failed: %v", err)
	}

	if !bytes.Equal(want.Bytes(), got.Bytes()) {
		t.Fatalf("byte mismatch:\n want=% x\n got =% x", want.Bytes(), got.Bytes())
	}
}

// TestUserQueryRespRoundTrip 验证 UserQueryResp 的反射读写往返一致。
func TestUserQueryRespRoundTrip(t *testing.T) {
	src := &demomsg.UserQueryResp{
		Found:    false,
		ID:       -12345,
		Account:  "abc",
		Nickname: "中文昵称",
		Level:    7,
	}

	w := codec.NewGameWriter(64)
	if err := codec.AutoWrite(w, src); err != nil {
		t.Fatalf("AutoWrite failed: %v", err)
	}

	dst := &demomsg.UserQueryResp{}
	r := codec.NewGameReader(w.Bytes())
	if err := codec.AutoRead(r, dst); err != nil {
		t.Fatalf("AutoRead failed: %v", err)
	}

	if *src != *dst {
		t.Fatalf("round-trip mismatch: src=%+v dst=%+v", src, dst)
	}
}

// TestEchoReqAutoRead 验证 EchoReq 反射读出与手写 ReadString 等价。
func TestEchoReqAutoRead(t *testing.T) {
	const text = "hello 世界"

	// 按原手写方式编码
	w := codec.NewGameWriter(64)
	w.WriteString(text)

	req := &demomsg.EchoReq{}
	r := codec.NewGameReader(w.Bytes())
	if err := codec.AutoRead(r, req); err != nil {
		t.Fatalf("AutoRead failed: %v", err)
	}
	if req.Text != text {
		t.Fatalf("text mismatch: want=%q got=%q", text, req.Text)
	}
}

// TestEchoMsgAutoWriteMatchesManual 验证 EchoMsg 反射写出
// 与原手写 w.WriteString(m.Text) 字节级等价。
func TestEchoMsgAutoWriteMatchesManual(t *testing.T) {
	m := &demomsg.EchoMsg{Text: "ping 你好"}

	want := codec.NewGameWriter(16)
	want.WriteString(m.Text)

	got := codec.NewGameWriter(16)
	if err := codec.AutoWrite(got, m); err != nil {
		t.Fatalf("AutoWrite failed: %v", err)
	}

	if !bytes.Equal(want.Bytes(), got.Bytes()) {
		t.Fatalf("byte mismatch:\n want=% x\n got =% x", want.Bytes(), got.Bytes())
	}
}
