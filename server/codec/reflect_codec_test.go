package codec

import (
	"bytes"
	"testing"
)

// internalStruct 仅用于 AutoWrite/AutoRead 单元测试，覆盖各内置类型。
type internalStruct struct {
	Bool   bool
	UByte  uint8
	Short  int16
	UShort uint16
	Int    int32
	UInt   uint32
	Long   int64
	ULong  uint64 `codec:"ulong"`
	F32    float32
	F64    float64
	Str    string
	Str2   string `codec:"string2"`
	Str4   string `codec:"string4"`
	Bytes  []byte

	// Go 平台相关 int / uint，自动推断为 4 字节。
	PlainInt  int
	PlainUint uint

	// 被跳过的字段
	Skipped string `codec:"-"`
}

// TestAutoWriteMatchesManual 验证反射 AutoWrite 的每个字段
// 与对应手写 GameWriter 方法产生的字节完全一致。
func TestAutoWriteMatchesManual(t *testing.T) {
	m := &internalStruct{
		Bool:      true,
		UByte:     0xAB,
		Short:     -2,
		UShort:    0xCAFE,
		Int:       0x12345678,
		UInt:      0xDEADBEEF,
		Long:      0x0123456789ABCDEF,
		ULong:     0xFEDCBA9876543210,
		F32:       3.14,
		F64:       2.718281828459045,
		Str:       "hero",
		Str2:      "猎魔人",
		Str4:      "hello 世界",
		Bytes:     []byte{0x01, 0x02, 0x03},
		PlainInt:  100,
		PlainUint: 200,
		Skipped:   "should-not-be-written",
	}

	// 手写预期：字段顺序必须与 struct 定义一致
	want := NewGameWriter(64)
	want.WriteBoolean(m.Bool)
	want.WriteUByte(int(m.UByte))
	want.WriteShort(m.Short)
	want.WriteUShort(m.UShort)
	want.WriteInt(m.Int)
	want.WriteUInt(m.UInt)
	want.WriteLong(m.Long)
	want.WriteLong(int64(m.ULong)) // ulong
	want.WriteFloat(m.F32)
	want.WriteDouble(m.F64)
	want.WriteString(m.Str)
	want.WriteString2(m.Str2)
	want.WriteString4(m.Str4)
	want.WriteBytes(m.Bytes)
	want.WriteInt(int32(m.PlainInt))
	want.WriteUInt(uint32(m.PlainUint))

	got := NewGameWriter(64)
	if err := AutoWrite(got, m); err != nil {
		t.Fatalf("AutoWrite failed: %v", err)
	}

	if !bytes.Equal(want.Bytes(), got.Bytes()) {
		t.Fatalf("byte mismatch:\n want=% x\n got =% x", want.Bytes(), got.Bytes())
	}
}

// TestAutoReadRoundTrip 验证 AutoWrite -> AutoRead 的往返一致性。
func TestAutoReadRoundTrip(t *testing.T) {
	src := &internalStruct{
		Bool:      false,
		UByte:     0x12,
		Short:     -100,
		UShort:    1000,
		Int:       -777,
		UInt:      888,
		Long:      -9999999,
		ULong:     1234567890,
		F32:       -1.5,
		F64:       6.022e23,
		Str:       "abc",
		Str2:      "中文",
		Str4:      "longer string",
		Bytes:     []byte{0xFF, 0x00},
		PlainInt:  42,
		PlainUint: 24,
	}

	w := NewGameWriter(64)
	if err := AutoWrite(w, src); err != nil {
		t.Fatalf("AutoWrite failed: %v", err)
	}

	dst := &internalStruct{}
	r := NewGameReader(w.Bytes())
	if err := AutoRead(r, dst); err != nil {
		t.Fatalf("AutoRead failed: %v", err)
	}

	// 逐字段比对（Skipped 应保持零值）
	if dst.Bool != src.Bool || dst.UByte != src.UByte || dst.Short != src.Short ||
		dst.UShort != src.UShort || dst.Int != src.Int || dst.UInt != src.UInt ||
		dst.Long != src.Long || dst.ULong != src.ULong || dst.F32 != src.F32 ||
		dst.F64 != src.F64 || dst.Str != src.Str || dst.Str2 != src.Str2 ||
		dst.Str4 != src.Str4 || !bytes.Equal(dst.Bytes, src.Bytes) ||
		dst.PlainInt != src.PlainInt || dst.PlainUint != src.PlainUint {
		t.Fatalf("round-trip mismatch:\n src=%+v\ndst=%+v", src, dst)
	}
	if dst.Skipped != "" {
		t.Fatalf("Skipped field should be zero, got %q", dst.Skipped)
	}
}

// TestUnsupportedSliceDoesNotPanic 回归测试：非 []byte 切片在反射写入时
// 不应 panic，而是返回明确的错误。
// 历史问题：defaultCodecType 把所有 Slice 都当 bytes，
// 走到 v.Bytes() 会对非 []uint8 切片 panic。
func TestUnsupportedSliceDoesNotPanic(t *testing.T) {
	type bad struct {
		List []int32
	}
	w := NewGameWriter(16)
	err := AutoWrite(w, &bad{List: []int32{1, 2, 3}})
	if err == nil {
		t.Fatal("expected error for non-[]byte slice, got nil")
	}
}

// TestSupportedByteSliceStillWorks 回归测试：[]byte 仍按 bytes tag 工作。
func TestSupportedByteSliceStillWorks(t *testing.T) {
	type ok struct {
		B []byte
	}
	src := &ok{B: []byte{0xAA, 0xBB, 0xCC}}
	w := NewGameWriter(16)
	if err := AutoWrite(w, src); err != nil {
		t.Fatalf("AutoWrite []byte failed: %v", err)
	}

	dst := &ok{}
	r := NewGameReader(w.Bytes())
	if err := AutoRead(r, dst); err != nil {
		t.Fatalf("AutoRead []byte failed: %v", err)
	}
	if !bytes.Equal(src.B, dst.B) {
		t.Fatalf("[]byte round-trip mismatch: want=% x got=% x", src.B, dst.B)
	}
}

// TestAutoWriteErrors 验证错误路径：非 struct / nil 指针 / 不支持的类型。
func TestAutoWriteErrors(t *testing.T) {
	w := NewGameWriter(16)

	if err := AutoWrite(w, 42); err == nil {
		t.Fatal("expected error for non-struct")
	}
	if err := AutoWrite(w, (*internalStruct)(nil)); err == nil {
		t.Fatal("expected error for nil pointer")
	}

	type bad struct {
		X map[string]int
	}
	if err := AutoWrite(w, &bad{}); err == nil {
		t.Fatal("expected error for unsupported map field")
	}
}

// TestAutoReadErrors 验证读路径错误。
func TestAutoReadErrors(t *testing.T) {
	r := NewGameReader(nil)

	if err := AutoRead(r, 42); err == nil {
		t.Fatal("expected error for non-struct")
	}
	if err := AutoRead(r, (*internalStruct)(nil)); err == nil {
		t.Fatal("expected error for nil pointer")
	}
}
