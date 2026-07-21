package main_thread

import "time"

// nowMilli 返回当前 unix 毫秒，专供 drop 日志限流使用。
// 抽成独立函数便于测试时替换时钟。
func nowMilli() int64 {
	return time.Now().UnixMilli()
}
