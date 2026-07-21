package codec

import "time"

// NowMilli 返回当前 unix 毫秒，供会话/连接层统一使用。
// 集中定义避免各包重复 time.Now().UnixMilli()，并方便后续做时钟替换（测试时钟注入）。
func NowMilli() int64 {
	return time.Now().UnixMilli()
}
