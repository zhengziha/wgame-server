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

// buildFieldStruct 用于测试 buildfield 标签
type buildFieldStruct struct {
	Fix17         int16  `codec:"short"`                 // 固定值17
	IntField      int32  `codec:"buildfield:int:263"`    // int类型buildfield
	StringField   string `codec:"buildfield:string:428"` // string类型buildfield
	LastLoginTime int32  `codec:"int"`                   // 普通int字段
}

// TestBuildFieldWrite 验证 buildfield 写入格式是否正确
// Java格式：先写2字节key(short)，再写1字节类型标记(3=int,4=string)，最后写值
func TestBuildFieldWrite(t *testing.T) {
	m := &buildFieldStruct{
		Fix17:         17,
		IntField:      12345,
		StringField:   "test_string",
		LastLoginTime: 67890,
	}

	w := NewGameWriter(64)
	if err := AutoWrite(w, m); err != nil {
		t.Fatalf("AutoWrite failed: %v", err)
	}

	got := w.Bytes()

	// 预期字节序列：
	// Fix17: 17 as short = [0, 17]
	// IntField: key=263 as short = [1, 7], type=3 (int) = [3], value=12345 as int = [0, 0, 48, 57]
	// StringField: key=428 as short = [1, 172], type=4 (string) = [4], value="test_string" = len=11 + "test_string"
	// LastLoginTime: 67890 as int = [0, 0, 167, 42]

	// 检查Fix17
	expectedFix17 := []byte{0, 17}
	if !bytes.Equal(got[0:2], expectedFix17) {
		t.Errorf("Fix17 mismatch: want % x, got % x", expectedFix17, got[0:2])
	}

	// 检查IntField的key
	expectedKey263 := []byte{1, 7} // 263 = 0x0107
	if !bytes.Equal(got[2:4], expectedKey263) {
		t.Errorf("IntField key mismatch: want % x, got % x", expectedKey263, got[2:4])
	}

	// 检查IntField的type标记
	if got[4] != 3 {
		t.Errorf("IntField type mismatch: want 3, got %d", got[4])
	}

	// 检查IntField的值
	expectedIntValue := []byte{0, 0, 48, 57} // 12345 = 0x00003039
	if !bytes.Equal(got[5:9], expectedIntValue) {
		t.Errorf("IntField value mismatch: want % x, got % x", expectedIntValue, got[5:9])
	}

	// 检查StringField的key
	expectedKey428 := []byte{1, 172} // 428 = 0x01AC
	if !bytes.Equal(got[9:11], expectedKey428) {
		t.Errorf("StringField key mismatch: want % x, got % x", expectedKey428, got[9:11])
	}

	// 检查StringField的type标记
	if got[11] != 4 {
		t.Errorf("StringField type mismatch: want 4, got %d", got[11])
	}

	// 检查StringField的值
	if got[12] != 11 { // "test_string" 的长度
		t.Errorf("StringField length mismatch: want 11, got %d", got[12])
	}
	if string(got[13:24]) != "test_string" {
		t.Errorf("StringField value mismatch: want 'test_string', got '%s'", string(got[13:24]))
	}

	// 检查LastLoginTime
	expectedInt67890 := []byte{0, 1, 9, 50} // 67890 = 0x00010932
	if !bytes.Equal(got[24:28], expectedInt67890) {
		t.Errorf("LastLoginTime mismatch: want % x, got % x", expectedInt67890, got[24:28])
	}

	t.Logf("buildfield write test passed, total bytes: %d", len(got))
}

// TestBuildFieldRead 验证 buildfield 读取的往返一致性
func TestBuildFieldRead(t *testing.T) {
	src := &buildFieldStruct{
		Fix17:         17,
		IntField:      54321,
		StringField:   "hello",
		LastLoginTime: 12345,
	}

	w := NewGameWriter(64)
	if err := AutoWrite(w, src); err != nil {
		t.Fatalf("AutoWrite failed: %v", err)
	}

	dst := &buildFieldStruct{}
	r := NewGameReader(w.Bytes())
	if err := AutoRead(r, dst); err != nil {
		t.Fatalf("AutoRead failed: %v", err)
	}

	if dst.Fix17 != src.Fix17 {
		t.Errorf("Fix17 mismatch: want %d, got %d", src.Fix17, dst.Fix17)
	}
	if dst.IntField != src.IntField {
		t.Errorf("IntField mismatch: want %d, got %d", src.IntField, dst.IntField)
	}
	if dst.StringField != src.StringField {
		t.Errorf("StringField mismatch: want '%s', got '%s'", src.StringField, dst.StringField)
	}
	if dst.LastLoginTime != src.LastLoginTime {
		t.Errorf("LastLoginTime mismatch: want %d, got %d", src.LastLoginTime, dst.LastLoginTime)
	}

	t.Logf("buildfield read/write round-trip test passed")
}
