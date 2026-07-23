package codec

import (
	"reflect"
	"testing"
)

// 测试用常量注册（init 函数自动注册）
func init() {
	RegisterBuildFieldKey("UserId", 1)
	RegisterBuildFieldKey("Level", 2)
	RegisterBuildFieldKey("Exp", 3)
	RegisterBuildFieldKey("Name", 4)
	RegisterBuildFieldKey("IsVip", 5)
	RegisterBuildFieldKey("Gold", 10)
}

// 使用常量名的简化格式测试
type SimpleMessage struct {
	UserId int32  `codec:"bf:UserId"` // 注册常量名
	Level  int32  `codec:"bf:Level"`
	Exp    int64  `codec:"bf:Exp"`
	Name   string `codec:"bf:Name"`
	IsVip  int8   `codec:"bf:IsVip"`
	Gold   int64  `codec:"bf:Gold"`
}

// 使用数字的兼容格式测试
type NumericMessage struct {
	UserId int32  `codec:"bf:1"` // 直接使用数字
	Level  int32  `codec:"bf:2"`
	Name   string `codec:"bf:4"`
}

// 完整格式测试结构体（向后兼容）
type FullMessage struct {
	UserId int32  `codec:"buildfield:int:UserId"`
	Level  int32  `codec:"buildfield:int:Level"`
	Name   string `codec:"buildfield:string:Name"`
}

// 混合格式：完整格式 + 数字 key
type FullNumericMessage struct {
	UserId int32  `codec:"buildfield:int:1"`
	Level  int32  `codec:"buildfield:int:2"`
	Name   string `codec:"buildfield:string:4"`
}

func TestBuildFieldConstantNameFormat(t *testing.T) {
	// 创建测试数据
	msg := &SimpleMessage{
		UserId: 1001,
		Level:  50,
		Exp:    999999,
		Name:   "测试玩家",
		IsVip:  1,
		Gold:   888888,
	}

	// 写入
	w := NewGameWriter(256)
	err := AutoWrite(w, msg)
	if err != nil {
		t.Fatalf("AutoWrite failed: %v", err)
	}

	data := w.Bytes()
	t.Logf("常量名格式写入 %d 字节", len(data))

	// 读取
	r := NewGameReader(data)
	var msg2 SimpleMessage
	err = AutoRead(r, &msg2)
	if err != nil {
		t.Fatalf("AutoRead failed: %v", err)
	}

	// 验证
	if msg2.UserId != msg.UserId {
		t.Errorf("UserId: expected %d, got %d", msg.UserId, msg2.UserId)
	}
	if msg2.Level != msg.Level {
		t.Errorf("Level: expected %d, got %d", msg.Level, msg2.Level)
	}
	if msg2.Exp != msg.Exp {
		t.Errorf("Exp: expected %d, got %d", msg.Exp, msg2.Exp)
	}
	if msg2.Name != msg.Name {
		t.Errorf("Name: expected %s, got %s", msg.Name, msg2.Name)
	}
	if msg2.IsVip != msg.IsVip {
		t.Errorf("IsVip: expected %d, got %d", msg.IsVip, msg2.IsVip)
	}
	if msg2.Gold != msg.Gold {
		t.Errorf("Gold: expected %d, got %d", msg.Gold, msg2.Gold)
	}
	t.Log("常量名格式读写测试通过!")
}

func TestBuildFieldNumericFormat(t *testing.T) {
	// 向后兼容测试：数字 key
	msg := &NumericMessage{
		UserId: 1001,
		Level:  50,
		Name:   "测试玩家",
	}

	w := NewGameWriter(256)
	err := AutoWrite(w, msg)
	if err != nil {
		t.Fatalf("AutoWrite failed: %v", err)
	}

	r := NewGameReader(w.Bytes())
	var msg2 NumericMessage
	err = AutoRead(r, &msg2)
	if err != nil {
		t.Fatalf("AutoRead failed: %v", err)
	}

	if msg2.UserId != msg.UserId || msg2.Name != msg.Name {
		t.Errorf("Numeric format mismatch: %+v vs %+v", msg, msg2)
	}
	t.Log("数字格式（向后兼容）读写测试通过!")
}

func TestBuildFieldFullFormat(t *testing.T) {
	// 完整格式 + 常量名
	msg := &FullMessage{
		UserId: 1001,
		Level:  50,
		Name:   "测试玩家",
	}

	w := NewGameWriter(256)
	err := AutoWrite(w, msg)
	if err != nil {
		t.Fatalf("AutoWrite failed: %v", err)
	}

	r := NewGameReader(w.Bytes())
	var msg2 FullMessage
	err = AutoRead(r, &msg2)
	if err != nil {
		t.Fatalf("AutoRead failed: %v", err)
	}

	if msg2.UserId != msg.UserId || msg2.Name != msg.Name {
		t.Errorf("Full format mismatch: %+v vs %+v", msg, msg2)
	}
	t.Log("完整格式（常量名）读写测试通过!")
}

func TestBuildFieldFullNumericFormat(t *testing.T) {
	// 完整格式 + 数字 key（向后兼容）
	msg := &FullNumericMessage{
		UserId: 1001,
		Level:  50,
		Name:   "测试玩家",
	}

	w := NewGameWriter(256)
	err := AutoWrite(w, msg)
	if err != nil {
		t.Fatalf("AutoWrite failed: %v", err)
	}

	r := NewGameReader(w.Bytes())
	var msg2 FullNumericMessage
	err = AutoRead(r, &msg2)
	if err != nil {
		t.Fatalf("AutoRead failed: %v", err)
	}

	if msg2.UserId != msg.UserId || msg2.Name != msg.Name {
		t.Errorf("Full numeric format mismatch: %+v vs %+v", msg, msg2)
	}
	t.Log("完整格式（数字）读写测试通过!")
}

func TestLookupBuildFieldKey(t *testing.T) {
	tests := []struct {
		name      string
		expectKey int16
		expectOK  bool
	}{
		// 已注册的常量
		{"UserId", 1, true},
		{"Level", 2, true},
		{"Name", 4, true},
		// 数字格式（兼容）
		{"263", 263, true},
		{"1", 1, true},
		{"-1", -1, true},
		// 不存在
		{"NonExistent", 0, false},
		{"abc", 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, ok := LookupBuildFieldKey(tt.name)
			if ok != tt.expectOK {
				t.Errorf("LookupBuildFieldKey(%q) ok = %v, want %v", tt.name, ok, tt.expectOK)
				return
			}
			if tt.expectOK && key != tt.expectKey {
				t.Errorf("LookupBuildFieldKey(%q) key = %d, want %d", tt.name, key, tt.expectKey)
			}
		})
	}
}

func TestInferCodecType(t *testing.T) {
	tests := []struct {
		kind     reflect.Kind
		expected string
	}{
		{reflect.Int8, "byte"},
		{reflect.Uint8, "ubyte"},
		{reflect.Int16, "short"},
		{reflect.Uint16, "ushort"},
		{reflect.Int32, "int"},
		{reflect.Uint32, "uint"},
		{reflect.Int64, "long"},
		{reflect.Uint64, "ulong"},
		{reflect.Float32, "float"},
		{reflect.Float64, "double"},
		{reflect.String, "string"},
		{reflect.Bool, "bool"},
	}

	for _, tt := range tests {
		t.Run(tt.kind.String(), func(t *testing.T) {
			result := inferCodecTypeFromGo(tt.kind)
			if result != tt.expected {
				t.Errorf("inferCodecTypeFromGo(%s) = %s, want %s", tt.kind, result, tt.expected)
			}
		})
	}
}

func TestParseBuildFieldTag(t *testing.T) {
	tests := []struct {
		tag         string
		expectType  string
		expectKey   int16
		expectError bool
	}{
		// 常量名格式
		{"bf:UserId", "", 1, false},
		{"bf:Name", "", 4, false},
		// 数字格式（兼容）
		{"bf:263", "", 263, false},
		{"bf:1", "", 1, false},
		// 完整格式 + 常量名
		{"buildfield:int:UserId", "int", 1, false},
		{"buildfield:string:Name", "string", 4, false},
		// 完整格式 + 数字（兼容）
		{"buildfield:int:263", "int", 263, false},
		{"buildfield:long:100", "long", 100, false},
		// 错误格式
		{"invalid", "", 0, true},
		{"bf:NonExistent", "", 0, true}, // 未注册的常量名
	}

	for _, tt := range tests {
		t.Run(tt.tag, func(t *testing.T) {
			fieldType, key, err := parseBuildFieldTag(tt.tag)
			if tt.expectError {
				if err == nil {
					t.Errorf("expected error for %s, got type=%s key=%d", tt.tag, fieldType, key)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if fieldType != tt.expectType {
				t.Errorf("type: expected %q, got %q", tt.expectType, fieldType)
			}
			if key != tt.expectKey {
				t.Errorf("key: expected %d, got %d", tt.expectKey, key)
			}
		})
	}
}
