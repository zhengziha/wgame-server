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
// 类型缺省推断（与现有 Java 协议一致，Level 用 Go 的 int 也写成 4 字节）：
//
//	bool        → bool
//	int8/uint8  → ubyte
//	int16       → short
//	uint16      → ushort
//	int32       → int
//	uint32      → uint
//	int64/uint64→ long
//	int         → int    （平台无关，固定 4 字节，匹配 Java int）
//	uint        → uint   （固定 4 字节）
//	float32     → float
//	float64     → double
//	string      → string
//	[]byte      → bytes
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
)

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
func fieldCodecType(f reflect.StructField, k reflect.Kind) (string, error) {
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
	return typ, nil
}

func writeField(w *GameWriter, f reflect.StructField, v reflect.Value) error {
	typ, err := fieldCodecType(f, v.Kind())
	if err != nil {
		return err
	}
	if typ == "" {
		return nil
	}
	switch typ {
	case tagBool:
		w.WriteBoolean(v.Bool())
	case tagUByte:
		w.WriteUByte(int(v.Uint()))
	case tagShort:
		w.WriteShort(int16(v.Int()))
	case tagUShort:
		w.WriteUShort(uint16(v.Uint()))
	case tagInt:
		w.WriteInt(int32(v.Int()))
	case tagUInt:
		w.WriteUInt(uint32(v.Uint()))
	case tagLong:
		w.WriteLong(v.Int())
	case tagULong:
		w.WriteLong(int64(v.Uint()))
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
	typ, err := fieldCodecType(f, v.Kind())
	if err != nil {
		return err
	}
	if typ == "" {
		return nil
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
		v.SetUint(uint64(n))
	case tagShort:
		n, err := r.ReadShort()
		if err != nil {
			return err
		}
		v.SetInt(int64(n))
	case tagUShort:
		n, err := r.ReadUShort()
		if err != nil {
			return err
		}
		v.SetUint(uint64(n))
	case tagInt:
		n, err := r.ReadInt()
		if err != nil {
			return err
		}
		v.SetInt(int64(n))
	case tagUInt:
		n, err := r.ReadUInt()
		if err != nil {
			return err
		}
		v.SetUint(uint64(n))
	case tagLong:
		n, err := r.ReadLong()
		if err != nil {
			return err
		}
		v.SetInt(n)
	case tagULong:
		n, err := r.ReadLong()
		if err != nil {
			return err
		}
		v.SetUint(uint64(n))
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
		// 仅默认 []byte / []uint8；其他切片需显式 tag（当前不支持）
		if t.Elem().Kind() == reflect.Uint8 {
			return tagBytes
		}
		return ""
	}
	return ""
}
