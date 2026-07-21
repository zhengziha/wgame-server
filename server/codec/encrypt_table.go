package codec

// 替换式加密表，对应 Java wd-server-fl 中的 v6/EncryptTable.java。
// encryptTable[i] 是一个 256 项的全排列（每项表示一个字节值的替换后结果）。
//
// 索引规则：
//   - header 中的 tableIndex 字段是 1-based（1..35）。
//   - tableIndex == 0 表示"不加密"。
//   - 查表使用 encryptTable[tableIndex - 1][原字节] = 加密后字节。
//
// 解密时使用反表：encryptInv[tableIndex - 1][加密后字节] = 原字节。
//
// 原始数据分两部分存放在 encrypt_table_raw1.go / encrypt_table_raw2.go。
// Java 源文件共 35 行（原项目注释中提及的 36 是历史笔误，实际是 35）。
var encryptTable [35][256]byte

// encryptInv 反表：encryptInv[i][encryptTable[i][b]] = b
var encryptInv [35][256]byte

// TableSize 返回可用表的数量（即最大 tableIndex）。
const TableSize = 35

// EncryptBody 对 body（包含 cmd 与 payload）使用指定 tableIndex 做替换加密。
// tableIndex 范围 1..35；超出范围则不做任何处理（视作不加密）。
func EncryptBody(tableIndex int, body []byte) {
	if tableIndex < 1 || tableIndex > TableSize {
		return
	}
	row := &encryptTable[tableIndex-1]
	for i := range body {
		body[i] = row[body[i]]
	}
}

// DecryptBody 对 body 做反向解密。tableIndex 范围 1..35。
func DecryptBody(tableIndex int, body []byte) {
	if tableIndex < 1 || tableIndex > TableSize {
		return
	}
	row := &encryptInv[tableIndex-1]
	for i := range body {
		body[i] = row[body[i]]
	}
}

func init() {
	// 合并两部分原始数据：part1 提供 0..15（16 行），part2 提供 16..34（19 行）。
	for i := 0; i < 16; i++ {
		for b := 0; b < 256; b++ {
			encryptTable[i][b] = byte(encryptRawPart1[i][b])
		}
	}
	for i := 0; i < 19; i++ {
		for b := 0; b < 256; b++ {
			encryptTable[i+16][b] = byte(encryptRawPart2[i][b])
		}
	}
	// 构造反表
	for i := 0; i < TableSize; i++ {
		for b := 0; b < 256; b++ {
			encryptInv[i][encryptTable[i][b]] = byte(b)
		}
	}
}
