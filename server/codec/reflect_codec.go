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
//	byte     1 字节有符号       WriteByte    / ReadByte
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
//	int8         → byte
//	uint8        → ubyte
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
	tagBool        = "bool"
	tagByte        = "byte"        // 有符号 1 字节，对应 Java writeByte/readByte
	tagUByte       = "ubyte"       // 无符号 1 字节，对应 Java writeUnsignedByte
	tagShort       = "short"       // 有符号 2 字节
	tagUShort      = "ushort"      // 无符号 2 字节
	tagInt         = "int"         // 有符号 4 字节
	tagUInt        = "uint"        // 无符号 4 字节
	tagLong        = "long"        // 有符号 8 字节
	tagULong       = "ulong"       // 无符号 8 字节
	tagFloat       = "float"       // 4 字节 IEEE754
	tagDouble      = "double"      // 8 字节 IEEE754
	tagString      = "string"      // 1 字节长度前缀 + GBK
	tagString2     = "string2"     // 2 字节长度前缀 + GBK
	tagString4     = "string4"     // 4 字节长度前缀 + GBK
	tagBytes       = "bytes"       // 2 字节长度前缀 + 原始字节
	tagList        = "list"        // 对象数组，可带 "list:<lentype>"
	tagBuildFields = "buildfields" // BuildField 容器标记（struct 字段，内部全是 bf 字段）
	tagBF          = "bf"          // 单个 BuildField 字段，格式: bf:<key> 或 bf:<type>:<key>
	tagStruct      = "struct"      // 嵌套结构体，递归调用 AutoWrite/AutoRead
)

// defaultListLenType 对象数组默认长度前缀类型。
// 与 wd-server-fl GameWriteTool.writeShort(size) 行为一致。
const defaultListLenType = tagShort

// AutoWrite 通过反射将 obj 的导出字段按声明顺序写入 w。
//
// 编码规则（严格按字段定义顺序，不做分类/重排）：
//   - 普通 int/string/byte 等字段：按 tag 直接写
//   - list 字段：写长度前缀 + 依次写元素
//   - bf:<key> / bf:<type>:<key> 单字段：按 key/type/value 写
//   - buildfields 容器字段（struct 标记 codec:"buildfields"）：先写 short(字段数)，
//     再按顺序展开内部字段（约定内部只含 bf 字段，数量编译期固定，无需运行时统计）
//   - 普通 struct 字段（无 codec 标签）：递归调用 AutoWrite
//
// 设计说明：BuildField 必须封装在独立子结构体中（如 VoExistedCharInfo），
// 用 codec:"buildfields" 显式标记。这样外层结构体按声明顺序处理时，
// 遇到该字段就知道要写数量前缀并展开，字节顺序与 Java 端完全一致。
//
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

	// 严格按字段定义顺序写入各字段
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if !f.IsExported() {
			continue
		}
		typ, err := fieldCodecType(f)
		if err != nil {
			return err
		}
		if typ == "" {
			continue // codec:"-" 表示跳过该字段
		}
		if err := writeField(w, f, v.Field(i)); err != nil {
			return err
		}
	}

	return nil
}

// countExportedCodecFields 统计结构体 t 中参与编解码的字段数量
// （导出且非 codec:"-"）。用于 buildfield 容器写入数量前缀。
func countExportedCodecFields(t reflect.Type) int {
	count := 0
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if !f.IsExported() {
			continue
		}
		typ, err := fieldCodecType(f)
		if err != nil || typ == "" {
			continue
		}
		count++
	}
	return count
}

// AutoRead 通过反射按声明顺序从 r 读取并填充 obj（必须为可寻址的 struct 或其指针）。
//
// 解码规则（严格按字段定义顺序，与 AutoWrite 对应）：
//   - 普通 int/string/byte 等字段：按 tag 直接读
//   - list 字段：读长度前缀 + 依次读元素
//   - bf:<key> / bf:<type>:<key> 单字段：按 key/type/value 读
//   - buildfields 容器字段（struct 标记 codec:"buildfields"）：先读 short(字段数) 校验，
//     再按顺序读内部字段
//   - 普通 struct 字段（无 codec 标签）：递归调用 AutoRead
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

	// 只有传入指针（如 &msg），解引用后的 v 才是可寻址的，才能修改字段值
	if !v.CanAddr() {
		return errors.New("codec.AutoRead: value not addressable; pass a pointer to struct")
	}

	t := v.Type()

	// 严格按字段定义顺序读取各字段
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if !f.IsExported() {
			continue
		}
		typ, err := fieldCodecType(f)
		if err != nil {
			return err
		}
		if typ == "" {
			continue // codec:"-" 跳过
		}
		if err := readField(r, f, v.Field(i)); err != nil {
			return err
		}
	}

	return nil
}

// fieldCodecType 解析字段的 codec tag，确定编码类型。
// 返回空串表示该字段应被跳过（codec:"-"）。
//
// 【反射知识点 - StructField】
// - f.Tag: 获取字段的 tag（如 `codec:"int" json:"name"`）
// - f.Tag.Get("codec"): 获取 codec tag 的值
// - f.Type: 获取字段的 Go 类型
//
// 返回值是规范的 tag 字符串（如 "int", "string2", "list:short"）。
func fieldCodecType(f reflect.StructField) (string, error) {
	// 获取 codec tag 的原始值
	raw := f.Tag.Get("codec")
	if raw == "-" {
		return "", nil // "-" 表示跳过该字段
	}

	// 去除空格并转为小写（tag 不区分大小写）
	trimmed := strings.TrimSpace(raw)
	lower := strings.ToLower(trimmed)

	// 检查是否为 BuildField tag
	// 两种形式：
	//   - "buildfields"     -> BuildField 容器标记（字段本身是 struct，内部全是 bf 字段）
	//   - "bf:<key>" 或 "bf:<type>:<key>" -> 单个 BuildField 字段
	if lower == tagBuildFields {
		return tagBuildFields, nil // 容器标记
	}
	if strings.HasPrefix(lower, tagBF+":") {
		// 保留 trimmed 的大小写，但把 type 部分转小写，保留 key 大小写
		rest := trimmed[len(tagBF)+1:]
		parts := strings.SplitN(rest, ":", 2)
		if len(parts) == 2 {
			// bf:<type>:<key> -> type 小写，key 保留
			return tagBF + ":" + strings.ToLower(parts[0]) + ":" + parts[1], nil
		}
		// bf:<key> -> 保留 key 大小写（常量名大小写敏感）
		return tagBF + ":" + rest, nil
	}

	// 普通 tag 类型
	typ := lower
	if typ == "" {
		// 如果没有指定 tag，根据 Go 类型推断编码类型
		typ = defaultCodecType(f.Type)
		if typ == "" {
			return "", fmt.Errorf("codec: field %s has unsupported type %s; set an explicit `codec:\"...\"` tag", f.Name, f.Type)
		}
	}

	// "list" 单独作为 tag 时补齐默认长度类型
	if typ == tagList {
		typ = tagList + ":" + defaultListLenType
	}
	return typ, nil
}

// isListTag 判断 tag 是否为对象数组（"list" 或 "list:..."）
func isListTag(typ string) bool {
	return typ == tagList || strings.HasPrefix(typ, tagList+":")
}

// isBuildFieldTag 判断 tag 是否为单个 BuildField 字段
// 支持格式: "bf:<key>" 或 "bf:<type>:<key>"
func isBuildFieldTag(typ string) bool {
	return strings.HasPrefix(typ, tagBF+":")
}

// parseBuildFieldTag 解析单个 BuildField 字段的 tag。
// 支持格式：
//   - "bf:<常量名>" -> 查映射表得 key，由调用方根据 Go 类型推断 codec 类型
//   - "bf:<数字>"   -> 直接用数字作为 key
//   - "bf:<type>:<常量名>" -> 显式指定 codec 类型 + 查映射表得 key
//   - "bf:<type>:<数字>"   -> 显式指定 codec 类型 + 直接用数字作为 key
//
// 返回 (fieldType, key, error)，fieldType 为空表示需要调用方推断。
func parseBuildFieldTag(typ string) (string, int16, error) {
	rest := strings.TrimPrefix(typ, tagBF+":")
	parts := strings.SplitN(rest, ":", 2)
	if len(parts) == 2 {
		// bf:<type>:<key>
		fieldType := parts[0]
		key, ok := LookupBuildFieldKey(parts[1])
		if !ok {
			return "", 0, fmt.Errorf("codec: unknown bf key %q in tag %q (not a registered constant or valid number)", parts[1], typ)
		}
		return fieldType, key, nil
	}
	// bf:<key>
	key, ok := LookupBuildFieldKey(parts[0])
	if !ok {
		return "", 0, fmt.Errorf("codec: unknown bf key %q in tag %q (not a registered constant or valid number)", parts[0], typ)
	}
	return "", key, nil // 空类型标记，表示需要从 Go 类型推断
}

// inferCodecTypeFromGo 根据 Go 类型推断 codec tag 类型
// Java 对比：类似于根据 Java 类型自动选择写入方法（writeInt/writeLong/writeString 等）
func inferCodecTypeFromGo(kind reflect.Kind) string {
	switch kind {
	case reflect.Bool:
		return tagBool
	case reflect.Int8:
		return tagByte
	case reflect.Uint8:
		return tagUByte
	case reflect.Int16:
		return tagShort
	case reflect.Uint16:
		return tagUShort
	case reflect.Int32:
		return tagInt
	case reflect.Uint32:
		return tagUInt
	case reflect.Int64:
		return tagLong
	case reflect.Uint64:
		return tagULong
	case reflect.Float32:
		return tagFloat
	case reflect.Float64:
		return tagDouble
	case reflect.String:
		return tagString
	default:
		return ""
	}
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

// writeField 根据字段的 tag 类型写入单个字段。
//
// 【反射知识点 - 读写字段值】
// - v.Bool(): 获取 bool 字段的值
// - v.Int(): 获取有符号整数字段的值
// - v.Uint(): 获取无符号整数字段的值
// - v.Float(): 获取浮点数字段的值
// - v.String(): 获取字符串字段的值
// - v.Bytes(): 获取 []byte 字段的值
//
// 注意：调用这些方法时，必须确保字段类型匹配，否则会 panic
func writeField(w *GameWriter, f reflect.StructField, v reflect.Value) error {
	// 解析字段的编码类型
	typ, err := fieldCodecType(f)
	if err != nil {
		return err
	}
	if typ == "" {
		return nil // 空类型表示跳过
	}

	// 处理特殊类型：列表、BuildField 容器、BuildField 单字段、嵌套结构体
	if isListTag(typ) {
		lenType := parseListLenType(typ)
		return writeList(w, f, v, lenType)
	}
	if typ == tagBuildFields {
		// BuildField 容器：写 short(字段数) + 展开内部 bf 字段
		return writeBuildFieldContainer(w, f, v)
	}
	if isBuildFieldTag(typ) {
		return writeBuildField(w, typ, f, v)
	}
	if typ == tagStruct {
		// 普通嵌套结构体：递归调用 AutoWrite 处理其内部字段
		return AutoWrite(w, v.Interface())
	}

	// 处理普通类型：根据 tag 类型选择写入方法
	switch typ {
	case tagBool:
		w.WriteBoolean(v.Bool()) // v.Bool() 获取 bool 值
	case tagByte:
		w.WriteByte(int8(intVal(v))) // intVal 统一处理有/无符号整数
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
		w.WriteFloat(float32(v.Float())) // v.Float() 获取浮点值
	case tagDouble:
		w.WriteDouble(v.Float())
	case tagString:
		w.WriteString(v.String()) // v.String() 获取字符串值
	case tagString2:
		w.WriteString2(v.String())
	case tagString4:
		w.WriteString4(v.String())
	case tagBytes:
		w.WriteBytes(v.Bytes()) // v.Bytes() 获取 []byte 值
	default:
		return fmt.Errorf("codec: unknown codec tag %q on field %s", typ, f.Name)
	}
	return nil
}

// writeBuildFieldContainer 写入 BuildField 容器（字段标记 codec:"buildfields"）。
//
// 对应 Java 端 writeShort(count) + 连续 BuildFieldsNew 写入。
// 容器字段必须是 struct（或指向 struct 的指针），内部约定只含 bf:<key> 或 bf:<type>:<key> 字段。
// 字段数量由结构体定义决定（编译期固定），直接统计导出且非跳过的字段数，无需运行时判断。
func writeBuildFieldContainer(w *GameWriter, f reflect.StructField, v reflect.Value) error {
	// 解引用指针
	for v.Kind() == reflect.Pointer {
		if v.IsNil() {
			return fmt.Errorf("codec: buildfield container %s is nil pointer", f.Name)
		}
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return fmt.Errorf("codec: buildfield container %s must be struct, got %s", f.Name, v.Kind())
	}

	t := v.Type()
	// 先写字段数量
	w.WriteShort(int16(countExportedCodecFields(t)))
	// 按顺序展开内部字段
	for i := 0; i < t.NumField(); i++ {
		inner := t.Field(i)
		if !inner.IsExported() {
			continue
		}
		innerTyp, err := fieldCodecType(inner)
		if err != nil {
			return err
		}
		if innerTyp == "" {
			continue
		}
		if err := writeField(w, inner, v.Field(i)); err != nil {
			return err
		}
	}
	return nil
}

// writeBuildField 写入 BuildFieldsNew 格式的字段
// Java格式：先写2字节key(short)，再写1字节类型标记(3=int,4=string)，最后写值
// 支持两种 tag 格式：
//   - "buildfield:<type>:<key>" 显式指定类型
//   - "bf:<key>" 从 Go 类型自动推断
func writeBuildField(w *GameWriter, tag string, f reflect.StructField, v reflect.Value) error {
	fieldType, key, err := parseBuildFieldTag(tag)
	if err != nil {
		return err
	}
	// 简化格式: bf:<key> -> 从 Go 类型自动推断
	if fieldType == "" {
		fieldType = inferCodecTypeFromGo(v.Kind())
		if fieldType == "" {
			return fmt.Errorf("codec: cannot infer codec type for Go kind %s on field %s", v.Kind(), f.Name)
		}
	}
	// 先写2字节key
	w.WriteShort(key)
	// 再写1字节类型标记
	switch fieldType {
	case tagInt, tagLong, tagShort, tagByte, tagUByte, tagUInt, tagULong, tagUShort:
		w.WriteByte(3) // 3=整数类型
	case tagString, tagString2:
		w.WriteByte(4) // 4=字符串类型
	default:
		return fmt.Errorf("codec: unsupported buildfield type %q on field %s", fieldType, f.Name)
	}
	// 最后写实际值
	switch fieldType {
	case tagBool:
		w.WriteBoolean(v.Bool())
	case tagByte:
		w.WriteByte(int8(intVal(v)))
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
	default:
		return fmt.Errorf("codec: unsupported buildfield type %q on field %s", fieldType, f.Name)
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
		lenType := parseListLenType(typ)
		return readList(r, f, v, lenType)
	}
	if typ == tagBuildFields {
		// BuildField 容器：读 short(字段数) 校验 + 读内部 bf 字段
		return readBuildFieldContainer(r, f, v)
	}
	if isBuildFieldTag(typ) {
		return readBuildField(r, typ, f, v)
	}
	if typ == tagStruct {
		// 嵌套结构体：递归调用 AutoRead 处理其内部字段
		// 需要 v 可寻址，以便 AutoRead 能修改内部字段值
		if !v.CanAddr() {
			return fmt.Errorf("codec: struct field %s not addressable", f.Name)
		}
		return AutoRead(r, v.Addr().Interface())
	}
	switch typ {
	case tagBool:
		b, err := r.ReadBoolean()
		if err != nil {
			return err
		}
		v.SetBool(b)
	case tagByte:
		n, err := r.ReadByte()
		if err != nil {
			return err
		}
		setIntVal(v, int64(n))
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

// readBuildFieldContainer 读取 BuildField 容器（字段标记 codec:"buildfields"）。
// 与 writeBuildFieldContainer 对应：先读 short(字段数) 并校验，再按顺序读内部字段。
func readBuildFieldContainer(r *GameReader, f reflect.StructField, v reflect.Value) error {
	// 解引用指针并确保可寻址
	for v.Kind() == reflect.Pointer {
		if v.IsNil() {
			v.Set(reflect.New(v.Type().Elem()))
		}
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return fmt.Errorf("codec: buildfield container %s must be struct, got %s", f.Name, v.Kind())
	}
	if !v.CanAddr() {
		return fmt.Errorf("codec: buildfield container %s not addressable", f.Name)
	}

	t := v.Type()
	// 读取并校验字段数量
	expect := countExportedCodecFields(t)
	n, err := r.ReadShort()
	if err != nil {
		return err
	}
	if int(n) != expect {
		return fmt.Errorf("codec: buildfield container %s count mismatch: expected %d, got %d", f.Name, expect, n)
	}
	// 按顺序读内部字段
	for i := 0; i < t.NumField(); i++ {
		inner := t.Field(i)
		if !inner.IsExported() {
			continue
		}
		innerTyp, err := fieldCodecType(inner)
		if err != nil {
			return err
		}
		if innerTyp == "" {
			continue
		}
		if err := readField(r, inner, v.Field(i)); err != nil {
			return err
		}
	}
	return nil
}

// readBuildField 读取 BuildFieldsNew 格式的字段
// Java格式：先读2字节key(short)，再读1字节类型标记(3=int,4=string)，最后读值
// 支持两种 tag 格式：
//   - "buildfield:<type>:<key>" 显式指定类型
//   - "bf:<key>" 从 Go 类型自动推断
func readBuildField(r *GameReader, tag string, f reflect.StructField, v reflect.Value) error {
	fieldType, _, err := parseBuildFieldTag(tag)
	if err != nil {
		return err
	}
	// 简化格式: bf:<key> -> 从 Go 类型自动推断
	if fieldType == "" {
		fieldType = inferCodecTypeFromGo(v.Kind())
		if fieldType == "" {
			return fmt.Errorf("codec: cannot infer codec type for Go kind %s on field %s", v.Kind(), f.Name)
		}
	}
	// 先读2字节key（跳过）
	_, err = r.ReadShort()
	if err != nil {
		return err
	}
	// 再读1字节类型标记（跳过，用于验证）
	_, err = r.ReadByte()
	if err != nil {
		return err
	}
	// 最后读实际值
	switch fieldType {
	case tagBool:
		b, err := r.ReadBoolean()
		if err != nil {
			return err
		}
		v.SetBool(b)
	case tagByte:
		n, err := r.ReadByte()
		if err != nil {
			return err
		}
		setIntVal(v, int64(n))
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
	default:
		return fmt.Errorf("codec: unsupported buildfield type %q on field %s", fieldType, f.Name)
	}
	return nil
}

// ---- 对象数组 ----

// listElemType 获取 slice 元素的实际类型（解引用指针后的 struct 类型）。
//
// 【反射知识点 - 类型操作】
// - f.Type: 获取字段的类型（如 []Member 或 []*Member）
// - f.Type.Kind(): 获取类型的种类（Slice, Pointer, Struct 等）
// - f.Type.Elem(): 获取 slice/pointer 的元素类型
// - elem.Elem(): 解引用指针，获取指向的类型
//
// 返回值：
//   - realElem: 元素的实际 struct 类型
//   - isPtr: 元素是否为指针类型（[]*Member 时为 true）
func listElemType(f reflect.StructField) (reflect.Type, bool, error) {
	// 检查字段是否为 slice 类型
	if f.Type.Kind() != reflect.Slice {
		return nil, false, fmt.Errorf("codec: list field %s is not a slice (%s)", f.Name, f.Type)
	}

	// 获取 slice 的元素类型
	elem := f.Type.Elem() // []Member → Member, []*Member → *Member

	// 检查元素是否为指针类型
	isPtr := elem.Kind() == reflect.Pointer
	realElem := elem
	if isPtr {
		// 如果是指针，解引用获取实际的 struct 类型
		realElem = elem.Elem() // *Member → Member
	}

	// 验证元素是否为 struct 类型
	if realElem.Kind() != reflect.Struct {
		return nil, false, fmt.Errorf("codec: list field %s element must be struct or *struct, got %s", f.Name, realElem.Kind())
	}
	return realElem, isPtr, nil
}

// writeList 写入对象数组（slice）。
//
// 【反射知识点 - slice 操作】
// - v.Len(): 获取 slice 的长度
// - v.Index(i): 获取 slice 的第 i 个元素
// - elem.Addr(): 获取元素的地址（返回 *Value）
// - elem.Addr().Interface(): 将 *Value 转为 interface{}
func writeList(w *GameWriter, f reflect.StructField, v reflect.Value, lenType string) error {
	// 验证元素类型
	if _, _, err := listElemType(f); err != nil {
		return err
	}

	// 获取 slice 长度
	n := v.Len()

	// 写入长度前缀
	if err := writeListLen(w, n, lenType, f.Name); err != nil {
		return err
	}

	// 遍历 slice 的每个元素
	for i := 0; i < n; i++ {
		elem := v.Index(i) // 获取第 i 个元素

		// 解引用指针（如果元素是指针类型）
		for elem.Kind() == reflect.Pointer {
			if elem.IsNil() {
				return fmt.Errorf("codec: list field %s element %d is nil pointer", f.Name, i)
			}
			elem = elem.Elem()
		}

		// 递归调用 AutoWrite 写入每个元素
		// elem.Addr() 获取元素的地址，转为 *Struct 类型
		// .Interface() 将 reflect.Value 转为 interface{}，以便传入 AutoWrite
		if err := AutoWrite(w, elem.Addr().Interface()); err != nil {
			return err
		}
	}
	return nil
}

// readList 读取对象数组（slice）。
//
// 【反射知识点 - 创建和设置 slice】
// - reflect.SliceOf(elemType): 创建 slice 类型
// - reflect.MakeSlice(type, len, cap): 创建新的 slice
// - reflect.New(type): 创建新的指针（分配内存）
// - elem.Set(value): 设置字段值
// - v.Set(slice): 用新创建的 slice 替换原字段
func readList(r *GameReader, f reflect.StructField, v reflect.Value, lenType string) error {
	// 获取元素类型信息
	realElem, isPtr, err := listElemType(f)
	if err != nil {
		return err
	}

	// 读取 slice 长度
	n, err := readListLen(r, lenType, f.Name)
	if err != nil {
		return err
	}
	if n < 0 {
		return fmt.Errorf("codec: list field %s got negative length %d", f.Name, n)
	}

	// ========== 创建新的 slice ==========
	// reflect.SliceOf(f.Type.Elem()) 创建 slice 类型
	// 例如：f.Type.Elem() 是 Member，则创建 []Member 类型
	sliceType := reflect.SliceOf(f.Type.Elem())

	// reflect.MakeSlice 创建指定长度和容量的 slice
	// 相当于 make([]Member, n, n)
	out := reflect.MakeSlice(sliceType, n, n)

	// 遍历并读取每个元素
	for i := 0; i < n; i++ {
		elem := out.Index(i) // 获取第 i 个元素

		if isPtr {
			// 如果元素是指针类型，需要分配内存
			// reflect.New(realElem) 创建一个 *Member 指针
			elem.Set(reflect.New(realElem))
			// 然后切到内部的 struct
			elem = elem.Elem()
		}

		// 递归调用 AutoRead 读取每个元素
		// elem.Addr() 获取元素的地址（必须是可寻址的）
		if err := AutoRead(r, elem.Addr().Interface()); err != nil {
			return err
		}
	}

	// 用新创建的 slice 替换原字段
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
	case reflect.Int8:
		return tagByte
	case reflect.Uint8:
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
	case reflect.Struct:
		// 嵌套结构体：递归调用 AutoWrite/AutoRead 处理其内部字段
		return tagStruct
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
