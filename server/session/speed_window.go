package session

import (
	"sync"
	"time"
)

const (
	ECHO_INTERVAL_LOWER_MS = 30
)

// SpeedWindow 环形缓冲区，用于检测加速器
// 参考 Java wd-server-fl 中的 SpeedWindow
//
// Java 对比：
//   - sync.Mutex 类似于 Java 的 synchronized 或 ReentrantLock
//   - defer w.mu.Unlock() 自动解锁，相当于 try-finally { lock.lock(); try { ... } finally { lock.unlock() } }
type SpeedWindow struct {
	buffer   []int64
	head     int
	size     int
	capacity int
	mu       sync.Mutex // Go 互斥锁：同一时刻只允许一个 goroutine 访问
}

// NewSpeedWindow 创建一个新的滑动窗口
func NewSpeedWindow(capacity int) *SpeedWindow {
	return &SpeedWindow{
		buffer:   make([]int64, capacity),
		capacity: capacity,
	}
}

// Add 添加一个时间戳（unix milli）
func (w *SpeedWindow) Add(timestamp int64) {
	w.mu.Lock()         // 获取锁：类似 Java 的 lock.lock()
	defer w.mu.Unlock() // defer 语句：函数返回时自动执行，确保锁一定被释放

	w.buffer[w.head] = timestamp
	w.head = (w.head + 1) % w.capacity
	if w.size < w.capacity {
		w.size++
	}
}

// Size 返回窗口中元素数量
func (w *SpeedWindow) Size() int {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.size
}

// CountIntervalsBetween 统计时间间隔在 [lower, upper) 范围内的次数
func (w *SpeedWindow) CountIntervalsBetween(lower, upper int) int {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.size < 2 {
		return 0
	}

	count := 0
	// 最旧元素的位置
	oldest := (w.head - w.size + w.capacity) % w.capacity
	for i := 1; i < w.size; i++ {
		prevIdx := (oldest + i - 1) % w.capacity
		currIdx := (oldest + i) % w.capacity
		interval := w.buffer[currIdx] - w.buffer[prevIdx]
		if interval >= int64(lower) && interval < int64(upper) {
			count++
		}
	}

	return count
}

// Reset 重置窗口
func (w *SpeedWindow) Reset() {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.head = 0
	w.size = 0
}

// NowMilli 返回当前时间的毫秒数
func NowMilli() int64 {
	return time.Now().UnixMilli()
}
