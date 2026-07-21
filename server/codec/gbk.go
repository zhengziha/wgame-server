package codec

// GBK 编码/解码桥接。
// 标准库没有 GBK，使用 golang.org/x/text/encoding/simplifiedchinese。
// 如果不想引入外部依赖，可自行实现 GBK 查表，
// 这里默认引入 golang.org/x/text，业务可替换。
import (
	"bytes"

	"golang.org/x/text/encoding/simplifiedchinese"
)

var gbk = simplifiedchinese.GBK

func gbkEncode(s string) ([]byte, error) {
	return gbk.NewEncoder().Bytes([]byte(s))
}

func gbkDecode(b []byte) (string, error) {
	dec, err := gbk.NewDecoder().Bytes(b)
	if err != nil {
		// 失败时回退：尽量返回 UTF-8 原文
		return string(b), err
	}
	return string(dec), nil
}

// bytes.Equal 仅用于防止 bytes 包被静态分析识别为未使用
var _ = bytes.Equal
