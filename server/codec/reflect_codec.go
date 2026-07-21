// Package codec 反射驱动的自动编解码工具。
//
// 设计目标：
//   - 让消息体 struct 不必再手写 WriteBody / ReadBody，由反射按字段顺序读写
//   - 字段编码方式由 `codec:"<type>"` tag 指定；缺省时按 Go 类型推断
//   - 与 GameWriter / GameReader 既有方法严格一一对应，行为与手写代码等价
//
// 支持的 tag type（不区分大小写）：
//
//	bool     1 字节布尔         WriteBoolean / ReadBoolean
//	ubyte    1 字节无符号       WriteUByte   / ReadUByte
//	short    2 字节有符号       WriteShort   / ReadShort
//	ushort   2 字节无符号       WriteUShort  / ReadUShort
//	int      4 字节有符号       WriteInt     / ReadInt
//	uint     4 字节无符号       WriteUInt    / ReadUInt
//	long     8 字节有符号       WriteLong    / ReadLong
//	ulong    8 字节无符号       WriteLong    / ReadLong（按位写入）
//	float    4 字节 IEEE754     WriteFloat   / ReadFloat
//	double   8 字节 IEEE754     WriteDouble  / ReadDouble
//	string   1 字节长度前缀     WriteString  / ReadString
//	string2  2 字节长度前缀     WriteString2 / ReadString2
//	string4  4 字节长度前缀     WriteString4 / ReadString4
//	bytes    2 字节长度前缀     WriteBytes   / ReadBytes
//
// 对象数组（参考 wd-server-fl MSG_FIXED_TEAM_DATA 的 List<Member> 模式）：
//
//	list             长度前缀 short（默认，2 字节有符号）
//	list:byte        长度前缀 1 字节（ubyte）
//	list:short       长度前缀 2 字节有符号
//	list:ushort      长度前缀 2 字节无符号
//	list:int         长度前缀 4 字节有符号
//	list:uint        长度前缀 4 字节无符号
//
// 仅支持元素类型为 struct 或指向 struct 的指针：
//
//	[]Member        → 按字段顺序展开每个元素
//	[]*Member       → 同上，元素可为 nil 时应在 tag 中说明（当前实现遇 nil 报错）
//
// 类型缺省推断（与现有 Java 协议一致，Level 用 Go 的 int 也写成 4 字节）：
//
//	bool         → bool
//	int8/uint8   → ubyte
//	int16        → short
//	uint16       → ushort
//	int32        → int
//	uint32       → uint
//	int64/uint64 → long
//	int          → int    （平台无关，固定 4 字节，匹配 Java int）
//	uint         → uint   （固定 4 字节）
//	float32      → float
//	float64      → double
//	string       → string
//	[]byte       → bytes
//	[]Struct     → list:short   （对象数组，长度前缀默认 short）
//	[]*Struct    → list:short
package codec

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"strings"
)

// codecTag 支持的编码类型常量（小写形式）
const (
	tagBool    = "bool"
	tagUByte   = "ubyte"
	tagShort   = "short"
	tagUShort  = "ushort"
	tagInt     = "int"
	tagUInt    = "uint"
	tagLong    = "long"
	tagULong   = "ulong"
	tagFloat   = "float"
	tagDouble  = "double"
	tagString  = "string"
	tagString2 = "string2"
	tagString4 = "string4"
	tagBytes   = "bytes"
	tagList    = "list" // 对象数组，可带 "list:<lentype>"
)

// defaultListLenType 对象数组默认长度前缀类型。
// 与 wd-server-fl GameWriteTool.writeShort(size) 行为一致。
const defaultListLenType = tagShort

// AutoWrite 通过反射将 obj 的导出字段按声明顺序写入 w。
// obj 可以是 struct 或指向 struct 的指针；遇到非结构体返回错误。
// 跳过未导出字段与带 `codec:"-"` 的字段。
func AutoWrite(w *GameWriter, obj interface{}) error {
	v := reflect.ValueOf(obj)
	for v.Kind() == reflect.Pointer {
		if v.IsNil() {
			return errors.New("codec.AutoWrite: nil pointer")
		}
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return fmt.Errorf("codec.AutoWrite: expected struct, got %s", v.Kind())
	}
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if !f.IsExported() {
			continue
		}
		if err := writeField(w, f, v.Field(i)); err != nil {
			return err
		}
	}
	return nil
}

// AutoRead 通过反射按声明顺序从 r 读取并填充 obj（必须为可寻址的 struct 或其指针）。
func AutoRead(r *GameReader, obj interface{}) error {
	v := reflect.ValueOf(obj)
	for v.Kind() == reflect.Pointer {
		if v.IsNil() {
			return errors.New("codec.AutoRead: nil pointer")
		}
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return fmt.Errorf("codec.AutoRead: expected struct, got %s", v.Kind())
	}
	if !v.CanAddr() {
		return errors.New("codec.AutoRead: value not addressable; pass a pointer to struct")
	}
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if !f.IsExported() {
			continue
		}
		if err := readField(r, f, v.Field(i)); err != nil {
			return err
		}
	}
	return nil
}

// fieldCodecType 解析字段 tag；返回空串表示该字段应被跳过（codec:"-").
//
// 返回值是规范的 tag 字符串（如 "int", "string2", "list:short"）。
// 对于 list，返回值包含完整 "list:<lentype>" 形式（默认 short）。
func fieldCodecType(f reflect.StructField) (string, error) {
	raw := f.Tag.Get("codec")
	if raw == "-" {
		return "", nil
	}
	typ := strings.ToLower(strings.TrimSpace(raw))
	if typ == "" {
		typ = defaultCodecType(f.Type)
		if typ == "" {
			return "", fmt.Errorf("codec: field %s has unsupported type %s; set an explicit `codec:\"...\"` tag", f.Name, f.Type)
		}
	}
	// "list" 单独作为 tag 时补齐默认长度类型，统一后续解析
	if typ == tagList {
		typ = tagList + ":" + defaultListLenType
	}
	return typ, nil
}

// isListTag 判断 tag 是否为对象数组（"list" 或 "list:..."）
func isListTag(typ string) bool {
	return typ == tagList || strings.HasPrefix(typ, tagList+":")
}

// parseListLenType 从 "list:<lentype>" 中提取 lentype；无后缀返回默认 short。
func parseListLenType(typ string) string {
	if idx := strings.Index(typ, ":"); idx >= 0 {
		return typ[idx+1:]
	}
	return defaultListLenType
}

// isUintKind 判断反射 Kind 是否属于无符号整数系列。
// 用于正确选择 v.Int() / v.Uint() 访问器，避免在 int8 等有符号字段上误用 Uint。
func isUintKind(k reflect.Kind) bool {
	return reflect.Uint <= k && k <= reflect.Uint64
}

// intVal 统一读取数值字段的 int64 值（同时支持有/无符号类型）。
func intVal(v reflect.Value) int64 {
	if isUintKind(v.Kind()) {
		return int64(v.Uint())
	}
	return v.Int()
}

// setIntVal 统一把 int64 写入数值字段（按目标 Kind 选 SetInt / SetUint）。
func setIntVal(v reflect.Value, n int64) {
	if isUintKind(v.Kind()) {
		v.SetUint(uint64(n))
	} else {
		v.SetInt(n)
	}
}

func writeField(w *GameWriter, f reflect.StructField, v reflect.Value) error {
	typ, err := fieldCodecType(f)
	if err != nil {
		return err
	}
	if typ == "" {
		return nil
	}
	if isListTag(typ) {
		return writeList(w, f, v, parseListLenType(typ))
	}
	switch typ {
	case tagBool:
		w.WriteBoolean(v.Bool())
	case tagUByte:
		w.WriteUByte(int(intVal(v)))
	case tagShort:
		w.WriteShort(int16(intVal(v)))
	case tagUShort:
		w.WriteUShort(uint16(intVal(v)))
	case tagInt:
		w.WriteInt(int32(intVal(v)))
	case tagUInt:
		w.WriteUInt(uint32(intVal(v)))
	case tagLong:
		w.WriteLong(intVal(v))
	case tagULong:
		w.WriteLong(int64(uint64(intVal(v))))
	case tagFloat:
		w.WriteFloat(float32(v.Float()))
	case tagDouble:
		w.WriteDouble(v.Float())
	case tagString:
		w.WriteString(v.String())
	case tagString2:
		w.WriteString2(v.String())
	case tagString4:
		w.WriteString4(v.String())
	case tagBytes:
		w.WriteBytes(v.Bytes())
	default:
		return fmt.Errorf("codec: unknown codec tag %q on field %s", typ, f.Name)
	}
	return nil
}

func readField(r *GameReader, f reflect.StructField, v reflect.Value) error {
	typ, err := fieldCodecType(f)
	if err != nil {
		return err
	}
	if typ == "" {
		return nil
	}
	if isListTag(typ) {
		return readList(r, f, v, parseListLenType(typ))
	}
	switch typ {
	case tagBool:
		b, err := r.ReadBoolean()
		if err != nil {
			return err
		}
		v.SetBool(b)
	case tagUByte:
		n, err := r.ReadUByte()
		if err != nil {
			return err
		}
		setIntVal(v, int64(n))
	case tagShort:
		n, err := r.ReadShort()
		if err != nil {
			return err
		}
		setIntVal(v, int64(n))
	case tagUShort:
		n, err := r.ReadUShort()
		if err != nil {
			return err
		}
		setIntVal(v, int64(n))
	case tagInt:
		n, err := r.ReadInt()
		if err != nil {
			return err
		}
		setIntVal(v, int64(n))
	case tagUInt:
		n, err := r.ReadUInt()
		if err != nil {
			return err
		}
		setIntVal(v, int64(n))
	case tagLong:
		n, err := r.ReadLong()
		if err != nil {
			return err
		}
		setIntVal(v, n)
	case tagULong:
		n, err := r.ReadLong()
		if err != nil {
			return err
		}
		setIntVal(v, n)
	case tagFloat:
		n, err := r.ReadUInt()
		if err != nil {
			return err
		}
		v.SetFloat(float64(math.Float32frombits(n)))
	case tagDouble:
		n, err := r.ReadLong()
		if err != nil {
			return err
		}
		v.SetFloat(math.Float64frombits(uint64(n)))
	case tagString:
		s, err := r.ReadString()
		if err != nil {
			return err
		}
		v.SetString(s)
	case tagString2:
		s, err := r.ReadString2()
		if err != nil {
			return err
		}
		v.SetString(s)
	case tagString4:
		s, err := r.ReadString4()
		if err != nil {
			return err
		}
		v.SetString(s)
	case tagBytes:
		b, err := r.ReadBytes()
		if err != nil {
			return err
		}
		v.SetBytes(b)
	default:
		return fmt.Errorf("codec: unknown codec tag %q on field %s", typ, f.Name)
	}
	return nil
}

// ---- 对象数组 ----

// listElemType 返回 slice 元素解引用指针后的 struct 类型；
// 若元素不是 struct 或 *struct 则返回错误。
func listElemType(f reflect.StructField) (reflect.Type, bool, error) {
	if f.Type.Kind() != reflect.Slice {
		return nil, false, fmt.Errorf("codec: list field %s is not a slice (%s)", f.Name, f.Type)
	}
	elem := f.Type.Elem()
	isPtr := elem.Kind() == reflect.Pointer
	realElem := elem
	if isPtr {
		realElem = elem.Elem()
	}
	if realElem.Kind() != reflect.Struct {
		return nil, false, fmt.Errorf("codec: list field %s element must be struct or *struct, got %s", f.Name, realElem.Kind())
	}
	return realElem, isPtr, nil
}

func writeList(w *GameWriter, f reflect.StructField, v reflect.Value, lenType string) error {
	if _, _, err := listElemType(f); err != nil {
		return err
	}
	n := v.Len()
	if err := writeListLen(w, n, lenType, f.Name); err != nil {
		return err
	}
	for i := 0; i < n; i++ {
		elem := v.Index(i)
		for elem.Kind() == reflect.Pointer {
			if elem.IsNil() {
				return fmt.Errorf("codec: list field %s element %d is nil pointer", f.Name, i)
			}
			elem = elem.Elem()
		}
		// 传指针给 AutoWrite，便于内部统一处理（AutoWrite 同时接受 struct / *struct）
		if err := AutoWrite(w, elem.Addr().Interface()); err != nil {
			return err
		}
	}
	return nil
}

func readList(r *GameReader, f reflect.StructField, v reflect.Value, lenType string) error {
	realElem, isPtr, err := listElemType(f)
	if err != nil {
		return err
	}
	n, err := readListLen(r, lenType, f.Name)
	if err != nil {
		return err
	}
	if n < 0 {
		return fmt.Errorf("codec: list field %s got negative length %d", f.Name, n)
	}

	sliceType := reflect.SliceOf(f.Type.Elem())
	out := reflect.MakeSlice(sliceType, n, n)
	for i := 0; i < n; i++ {
		elem := out.Index(i)
		if isPtr {
			elem.Set(reflect.New(realElem)) // 分配 *Struct
			elem = elem.Elem()              // 切到内部 struct
		}
		// elem 此时为 addressable 的 struct；传指针给 AutoRead
		if err := AutoRead(r, elem.Addr().Interface()); err != nil {
			return err
		}
	}
	v.Set(out)
	return nil
}

// writeListLen 按指定长度前缀类型写入元素个数。
func writeListLen(w *GameWriter, n int, lenType, fieldName string) error {
	switch lenType {
	case tagUByte, "byte":
		if n > 0xFF {
			return fmt.Errorf("codec: list field %s length %d exceeds ubyte (max 255)", fieldName, n)
		}
		w.WriteUByte(n)
	case tagShort:
		if n > 0x7FFF {
			return fmt.Errorf("codec: list field %s length %d exceeds short (max 32767)", fieldName, n)
		}
		w.WriteShort(int16(n))
	case tagUShort:
		if n > 0xFFFF {
			return fmt.Errorf("codec: list field %s length %d exceeds ushort (max 65535)", fieldName, n)
		}
		w.WriteUShort(uint16(n))
	case tagInt:
		w.WriteInt(int32(n))
	case tagUInt:
		w.WriteUInt(uint32(n))
	default:
		return fmt.Errorf("codec: list field %s unsupported length prefix type %q (use byte/short/ushort/int/uint)",
			fieldName, lenType)
	}
	return nil
}

// readListLen 按指定长度前缀类型读取元素个数。
func readListLen(r *GameReader, lenType, fieldName string) (int, error) {
	switch lenType {
	case tagUByte, "byte":
		n, err := r.ReadUByte()
		return n, err
	case tagShort:
		n, err := r.ReadShort()
		return int(n), err
	case tagUShort:
		n, err := r.ReadUShort()
		return int(n), err
	case tagInt:
		n, err := r.ReadInt()
		return int(n), err
	case tagUInt:
		n, err := r.ReadUInt()
		return int(n), err
	default:
		return 0, fmt.Errorf("codec: list field %s unsupported length prefix type %q", fieldName, lenType)
	}
}

// defaultCodecType 根据 Go 类型推断默认编码；返回空串表示无法推断。
func defaultCodecType(t reflect.Type) string {
	switch t.Kind() {
	case reflect.Bool:
		return tagBool
	case reflect.Int8, reflect.Uint8:
		return tagUByte
	case reflect.Int16:
		return tagShort
	case reflect.Uint16:
		return tagUShort
	case reflect.Int32, reflect.Int:
		return tagInt
	case reflect.Uint32, reflect.Uint:
		return tagUInt
	case reflect.Int64, reflect.Uint64:
		return tagLong
	case reflect.Float32:
		return tagFloat
	case reflect.Float64:
		return tagDouble
	case reflect.String:
		return tagString
	case reflect.Slice:
		// []byte / []uint8 → bytes
		if t.Elem().Kind() == reflect.Uint8 {
			return tagBytes
		}
		// []struct / []*struct → list:short（与 Java GameWriteTool.writeShort(size) 默认一致）
		if elem := t.Elem(); elem.Kind() == reflect.Struct ||
			(elem.Kind() == reflect.Pointer && elem.Elem().Kind() == reflect.Struct) {
			return tagList + ":" + defaultListLenType
		}
		return ""
	}
	return ""
}
