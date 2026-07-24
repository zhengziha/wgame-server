package codec

import (
	"bytes"
	"testing"
)

// teamMember 模拟 wd-server-fl Vo_FIXED_TEAM_DATA.Member 的字段顺序与类型，
// 用于验证对象数组反射编解码与 Java 手写 GameWriteTool.writeXxx 字节级等价。
//
// Java 原顺序：
//
//	writeString(gid)
//	writeString(name)
//	writeShort(level)
//	writeInt(icon)
//	writeInt(tao)
//	writeInt(lastLogoutTime)
//	writeInt(joinTime)
type teamMember struct {
	Gid            string
	Name           string
	Level          int16
	Icon           int32
	Tao            int32
	LastLogoutTime int32
	JoinTime       int32
}

// teamData 模拟 Vo_FIXED_TEAM_DATA 外层：
//
//	writeString(name)
//	writeByte(level)        // 注意：外层 level 是 byte，内层是 short
//	writeInt(intimacy)
//	writeInt(maxIntimacy)
//	writeShort(members.size()) + 循环展开每个 member
type teamData struct {
	Name        string
	Level       int8 `codec:"ubyte"` // 对应 Java writeByte（无符号 1 字节）
	Intimacy    int32
	MaxIntimacy int32
	Members     []teamMember // 默认 list:short，与 Java writeShort(size) 一致
}

// ptrTeamData 演示 []*Struct 元素的读写。
type ptrTeamData struct {
	Members []*teamMember
}

// TestTeamDataAutoWriteMatchesManual 验证反射自动写出与手写 Java 协议字节级等价。
func TestTeamDataAutoWriteMatchesManual(t *testing.T) {
	src := &teamData{
		Name:        "战神殿",
		Level:       99,
		Intimacy:    12345,
		MaxIntimacy: 99999,
		Members: []teamMember{
			{
				Gid:            "g-001",
				Name:           "龙骑士",
				Level:          88,
				Icon:           1001,
				Tao:            7777,
				LastLogoutTime: 1700000000,
				JoinTime:       1699000000,
			},
			{
				Gid:            "g-002",
				Name:           "魔法师",
				Level:          76,
				Icon:           1002,
				Tao:            8888,
				LastLogoutTime: 1700000001,
				JoinTime:       1699000001,
			},
		},
	}

	// 手写预期：完全对照 Java GameWriteTool 的调用顺序
	want := NewGameWriter(128)
	want.WriteString(src.Name)
	want.WriteUByte(int(src.Level))
	want.WriteInt(src.Intimacy)
	want.WriteInt(src.MaxIntimacy)
	want.WriteShort(int16(len(src.Members)))
	for _, m := range src.Members {
		want.WriteString(m.Gid)
		want.WriteString(m.Name)
		want.WriteShort(m.Level)
		want.WriteInt(m.Icon)
		want.WriteInt(m.Tao)
		want.WriteInt(m.LastLogoutTime)
		want.WriteInt(m.JoinTime)
	}

	got := NewGameWriter(128)
	if err := AutoWrite(got, src); err != nil {
		t.Fatalf("AutoWrite failed: %v", err)
	}

	if !bytes.Equal(want.Bytes(), got.Bytes()) {
		t.Fatalf("byte mismatch:\n want=% x\n got =% x", want.Bytes(), got.Bytes())
	}
}

// TestTeamDataRoundTrip 验证反射读写往返一致。
func TestTeamDataRoundTrip(t *testing.T) {
	src := &teamData{
		Name:        "联萌",
		Level:       1,
		Intimacy:    0,
		MaxIntimacy: 100,
		Members: []teamMember{
			{Gid: "a", Name: "甲", Level: 10, Icon: 1, Tao: 2, LastLogoutTime: 3, JoinTime: 4},
			{Gid: "b", Name: "乙", Level: 20, Icon: 5, Tao: 6, LastLogoutTime: 7, JoinTime: 8},
			{Gid: "c", Name: "丙 中文", Level: 30, Icon: 9, Tao: 10, LastLogoutTime: 11, JoinTime: 12},
		},
	}

	w := NewGameWriter(128)
	if err := AutoWrite(w, src); err != nil {
		t.Fatalf("AutoWrite failed: %v", err)
	}

	dst := &teamData{}
	r := NewGameReader(w.Bytes())
	if err := AutoRead(r, dst); err != nil {
		t.Fatalf("AutoRead failed: %v", err)
	}

	if dst.Name != src.Name || dst.Level != src.Level ||
		dst.Intimacy != src.Intimacy || dst.MaxIntimacy != src.MaxIntimacy {
		t.Fatalf("scalar mismatch: src=%+v dst=%+v", src, dst)
	}
	if len(dst.Members) != len(src.Members) {
		t.Fatalf("members length mismatch: want=%d got=%d", len(src.Members), len(dst.Members))
	}
	for i := range src.Members {
		if dst.Members[i] != src.Members[i] {
			t.Fatalf("member[%d] mismatch:\n want=%+v\n got =%+v", i, src.Members[i], dst.Members[i])
		}
	}
}

// TestPtrSliceRoundTrip 验证 []*Struct 元素的读写。
func TestPtrSliceRoundTrip(t *testing.T) {
	src := &ptrTeamData{
		Members: []*teamMember{
			{Gid: "p1", Name: "玩家一", Level: 11, Icon: 21, Tao: 31, LastLogoutTime: 41, JoinTime: 51},
			{Gid: "p2", Name: "玩家二", Level: 12, Icon: 22, Tao: 32, LastLogoutTime: 42, JoinTime: 52},
		},
	}

	w := NewGameWriter(128)
	if err := AutoWrite(w, src); err != nil {
		t.Fatalf("AutoWrite failed: %v", err)
	}

	dst := &ptrTeamData{}
	r := NewGameReader(w.Bytes())
	if err := AutoRead(r, dst); err != nil {
		t.Fatalf("AutoRead failed: %v", err)
	}

	if len(dst.Members) != len(src.Members) {
		t.Fatalf("members length mismatch: want=%d got=%d", len(src.Members), len(dst.Members))
	}
	for i := range src.Members {
		if dst.Members[i] == nil {
			t.Fatalf("dst.Members[%d] is nil", i)
		}
		if *dst.Members[i] != *src.Members[i] {
			t.Fatalf("member[%d] mismatch:\n want=%+v\n got =%+v", i, *src.Members[i], *dst.Members[i])
		}
	}
}

// TestListLenTypes 验证不同长度前缀类型（byte/short/int）都能正确读写。
func TestListLenTypes(t *testing.T) {
	type withByteLen struct {
		Items []teamMember `codec:"list:byte"`
	}
	type withShortLen struct {
		Items []teamMember `codec:"list:short"`
	}
	type withIntLen struct {
		Items []teamMember `codec:"list:int"`
	}
	type withUShortLen struct {
		Items []teamMember `codec:"list:ushort"`
	}

	mk := func() []teamMember {
		return []teamMember{
			{Gid: "x", Name: "y", Level: 1, Icon: 2, Tao: 3, LastLogoutTime: 4, JoinTime: 5},
		}
	}

	// byte 长度
	t.Run("byte", func(t *testing.T) {
		src := &withByteLen{Items: mk()}
		w := NewGameWriter(64)
		if err := AutoWrite(w, src); err != nil {
			t.Fatalf("AutoWrite: %v", err)
		}
		// 校验首字节即长度
		if int(w.Bytes()[0]) != len(src.Items) {
			t.Fatalf("byte-len prefix mismatch: got=%d want=%d", w.Bytes()[0], len(src.Items))
		}
		dst := &withByteLen{}
		if err := AutoRead(NewGameReader(w.Bytes()), dst); err != nil {
			t.Fatalf("AutoRead: %v", err)
		}
		if len(dst.Items) != 1 || dst.Items[0] != src.Items[0] {
			t.Fatalf("round-trip mismatch: got=%+v", dst.Items)
		}
	})

	// short 长度（默认）：前 2 字节为大端 int16
	t.Run("short", func(t *testing.T) {
		src := &withShortLen{Items: mk()}
		w := NewGameWriter(64)
		if err := AutoWrite(w, src); err != nil {
			t.Fatalf("AutoWrite: %v", err)
		}
		// short 长度占 2 字节大端
		if int(w.Bytes()[0])<<8|int(w.Bytes()[1]) != len(src.Items) {
			t.Fatalf("short-len prefix mismatch: % x", w.Bytes()[:2])
		}
		dst := &withShortLen{}
		if err := AutoRead(NewGameReader(w.Bytes()), dst); err != nil {
			t.Fatalf("AutoRead: %v", err)
		}
		if len(dst.Items) != 1 || dst.Items[0] != src.Items[0] {
			t.Fatalf("round-trip mismatch: got=%+v", dst.Items)
		}
	})

	// ushort 长度
	t.Run("ushort", func(t *testing.T) {
		src := &withUShortLen{Items: mk()}
		w := NewGameWriter(64)
		if err := AutoWrite(w, src); err != nil {
			t.Fatalf("AutoWrite: %v", err)
		}
		dst := &withUShortLen{}
		if err := AutoRead(NewGameReader(w.Bytes()), dst); err != nil {
			t.Fatalf("AutoRead: %v", err)
		}
		if len(dst.Items) != 1 || dst.Items[0] != src.Items[0] {
			t.Fatalf("round-trip mismatch: got=%+v", dst.Items)
		}
	})

	// int 长度（4 字节）
	t.Run("int", func(t *testing.T) {
		src := &withIntLen{Items: mk()}
		w := NewGameWriter(64)
		if err := AutoWrite(w, src); err != nil {
			t.Fatalf("AutoWrite: %v", err)
		}
		// int 长度占 4 字节大端
		gotLen := int(w.Bytes()[0])<<24 | int(w.Bytes()[1])<<16 | int(w.Bytes()[2])<<8 | int(w.Bytes()[3])
		if gotLen != len(src.Items) {
			t.Fatalf("int-len prefix mismatch: % x", w.Bytes()[:4])
		}
		dst := &withIntLen{}
		if err := AutoRead(NewGameReader(w.Bytes()), dst); err != nil {
			t.Fatalf("AutoRead: %v", err)
		}
		if len(dst.Items) != 1 || dst.Items[0] != src.Items[0] {
			t.Fatalf("round-trip mismatch: got=%+v", dst.Items)
		}
	})
}

// TestListEmptySlice 验证空切片写出长度=0、读回也是空切片。
func TestListEmptySlice(t *testing.T) {
	src := &teamData{Name: "solo", Members: nil}
	w := NewGameWriter(32)
	if err := AutoWrite(w, src); err != nil {
		t.Fatalf("AutoWrite: %v", err)
	}

	dst := &teamData{}
	if err := AutoRead(NewGameReader(w.Bytes()), dst); err != nil {
		t.Fatalf("AutoRead: %v", err)
	}
	if len(dst.Members) != 0 {
		t.Fatalf("expected empty members, got %d", len(dst.Members))
	}
}

// TestListUnsupportedElementType 非 struct 切片需报错（当前不支持基本类型数组）。
func TestListUnsupportedElementType(t *testing.T) {
	type bad struct {
		Nums []int32 `codec:"list"`
	}
	w := NewGameWriter(16)
	err := AutoWrite(w, &bad{Nums: []int32{1, 2, 3}})
	if err == nil {
		t.Fatal("expected error for non-struct slice with list tag")
	}
}

// TestListLenOverflow byte 长度前缀溢出时应报错而非截断。
func TestListLenOverflow(t *testing.T) {
	type withByteLen struct {
		Items []teamMember `codec:"list:byte"`
	}
	// 构造 256 个元素，超过 byte 上限
	src := &withByteLen{Items: make([]teamMember, 256)}
	w := NewGameWriter(16)
	err := AutoWrite(w, src)
	if err == nil {
		t.Fatal("expected overflow error for byte-length list with 256 items")
	}
}

// TestBuildFieldAutoCount 验证 buildfield 容器自动写入字段数量前缀。
// 新设计：BuildField 字段必须封装在独立结构体中，用 codec:"buildfields" 标记，
// 序列化时自动写 short(字段数) + 展开 bf 字段。
func TestBuildFieldAutoCount(t *testing.T) {
	// BuildField 容器：只含 bf 字段
	type VoExistedCharInfo struct {
		LeftTimeToDelete int32  `codec:"bf:LeftTimeToDelete"`
		Level            int32  `codec:"bf:Level"`
		Name             string `codec:"bf:Name"`
		Gid              string `codec:"bf:Gid"`
	}

	// 外层：buildfield 容器 + 普通字段
	type VoExistedChar struct {
		Info          VoExistedCharInfo `codec:"buildfields"` // BuildField 容器
		LastLoginTime int32             `codec:"int"`         // 普通字段
		LoginMac      string            `codec:"string"`      // 普通字段
	}

	src := VoExistedChar{
		Info: VoExistedCharInfo{
			LeftTimeToDelete: 100,
			Level:            30,
			Name:             "测试角色",
			Gid:              "test_gid_001",
		},
		LastLoginTime: 1700000000,
		LoginMac:      "00:11:22:33:44:55",
	}

	// 使用自动序列化
	w := NewGameWriter(128)
	if err := AutoWrite(w, &src); err != nil {
		t.Fatalf("AutoWrite failed: %v", err)
	}

	// 验证字节序列：先读 BuildField 数量（应为 4）
	r := NewGameReader(w.Bytes())
	bfCount, err := r.ReadShort()
	if err != nil {
		t.Fatalf("ReadShort failed: %v", err)
	}
	if bfCount != 4 {
		t.Fatalf("expected 4 buildfield fields, got %d", bfCount)
	}

	// 验证往返读写字段正确
	dst := VoExistedChar{}
	if err := AutoRead(NewGameReader(w.Bytes()), &dst); err != nil {
		t.Fatalf("AutoRead failed: %v", err)
	}

	if dst != src {
		t.Fatalf("mismatch: want=%+v got=%+v", src, dst)
	}
}

// TestBuildFieldAutoCountZero 验证没有 BuildField 字段时不写入数量前缀。
func TestBuildFieldAutoCountZero(t *testing.T) {
	type SimpleMsg struct {
		IntVal    int32  `codec:"int"`
		StringVal string `codec:"string"`
	}

	src := SimpleMsg{
		IntVal:    42,
		StringVal: "hello",
	}

	w := NewGameWriter(64)
	if err := AutoWrite(w, &src); err != nil {
		t.Fatalf("AutoWrite failed: %v", err)
	}

	// 验证没有 BuildField 数量前缀，直接是普通字段
	r := NewGameReader(w.Bytes())
	intVal, err := r.ReadInt()
	if err != nil {
		t.Fatalf("ReadInt failed: %v", err)
	}
	if intVal != 42 {
		t.Fatalf("expected 42, got %d", intVal)
	}

	dst := SimpleMsg{}
	if err := AutoRead(NewGameReader(w.Bytes()), &dst); err != nil {
		t.Fatalf("AutoRead failed: %v", err)
	}

	if dst != src {
		t.Fatalf("mismatch: want=%+v got=%+v", src, dst)
	}
}
