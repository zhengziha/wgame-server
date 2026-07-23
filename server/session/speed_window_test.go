package session

import (
	"sync"
	"testing"
)

// TestSpeedWindowAddAndSize 测试添加元素和大小计算
func TestSpeedWindowAddAndSize(t *testing.T) {
	w := NewSpeedWindow(5)

	if w.Size() != 0 {
		t.Errorf("Initial size: got %d, want 0", w.Size())
	}

	w.Add(100)
	if w.Size() != 1 {
		t.Errorf("Size after 1 add: got %d, want 1", w.Size())
	}

	w.Add(200)
	w.Add(300)
	if w.Size() != 3 {
		t.Errorf("Size after 3 adds: got %d, want 3", w.Size())
	}

	// 填满缓冲区
	w.Add(400)
	w.Add(500)
	if w.Size() != 5 {
		t.Errorf("Size after 5 adds: got %d, want 5", w.Size())
	}

	// 超过容量，应该覆盖旧数据
	w.Add(600)
	if w.Size() != 5 {
		t.Errorf("Size after 6 adds: got %d, want 5", w.Size())
	}
}

// TestSpeedWindowCountIntervalsBetween 测试时间间隔统计
func TestSpeedWindowCountIntervalsBetween(t *testing.T) {
	w := NewSpeedWindow(5)

	// 少于2个元素，返回0
	if w.CountIntervalsBetween(0, 1000) != 0 {
		t.Error("CountIntervalsBetween with <2 elements should return 0")
	}

	// 添加时间戳，间隔为100ms
	w.Add(100)
	w.Add(200) // 间隔100ms
	w.Add(300) // 间隔100ms
	w.Add(400) // 间隔100ms
	w.Add(500) // 间隔100ms

	// 统计 [90, 110) 范围内的间隔
	count := w.CountIntervalsBetween(90, 110)
	if count != 4 {
		t.Errorf("CountIntervalsBetween(90,110): got %d, want 4", count)
	}

	// 统计 [0, 90) 范围内的间隔
	count = w.CountIntervalsBetween(0, 90)
	if count != 0 {
		t.Errorf("CountIntervalsBetween(0,90): got %d, want 0", count)
	}
}

// TestSpeedWindowCountIntervalsBetweenMixed 测试混合间隔场景
func TestSpeedWindowCountIntervalsBetweenMixed(t *testing.T) {
	w := NewSpeedWindow(5)

	// 添加时间戳，间隔混合
	w.Add(100)
	w.Add(150) // 间隔50ms
	w.Add(250) // 间隔100ms
	w.Add(320) // 间隔70ms
	w.Add(450) // 间隔130ms

	// 统计 [60, 120) 范围内的间隔（应该是100ms和70ms）
	count := w.CountIntervalsBetween(60, 120)
	if count != 2 {
		t.Errorf("CountIntervalsBetween(60,120): got %d, want 2", count)
	}
}

// TestSpeedWindowReset 测试重置功能
func TestSpeedWindowReset(t *testing.T) {
	w := NewSpeedWindow(5)

	w.Add(100)
	w.Add(200)
	w.Add(300)

	if w.Size() != 3 {
		t.Errorf("Size before reset: got %d, want 3", w.Size())
	}

	w.Reset()

	if w.Size() != 0 {
		t.Errorf("Size after reset: got %d, want 0", w.Size())
	}

	if w.CountIntervalsBetween(0, 1000) != 0 {
		t.Error("CountIntervalsBetween after reset should return 0")
	}
}

// TestSpeedWindowCircularBuffer 测试环形缓冲区覆盖逻辑
func TestSpeedWindowCircularBuffer(t *testing.T) {
	w := NewSpeedWindow(3)

	// 添加3个元素
	w.Add(100)
	w.Add(200)
	w.Add(300)

	// 统计间隔 [90, 110)
	count := w.CountIntervalsBetween(90, 110)
	if count != 2 {
		t.Errorf("Count before overwrite: got %d, want 2", count)
	}

	// 添加第4个元素，覆盖第一个
	w.Add(400)

	// 新的间隔: 200->300 (100ms), 300->400 (100ms)
	count = w.CountIntervalsBetween(90, 110)
	if count != 2 {
		t.Errorf("Count after 1 overwrite: got %d, want 2", count)
	}

	// 添加第5个元素，覆盖第二个
	w.Add(500)

	// 新的间隔: 300->400 (100ms), 400->500 (100ms)
	count = w.CountIntervalsBetween(90, 110)
	if count != 2 {
		t.Errorf("Count after 2 overwrites: got %d, want 2", count)
	}
}

// TestSpeedWindowAcceleratorDetection 测试加速器检测场景
func TestSpeedWindowAcceleratorDetection(t *testing.T) {
	w := NewSpeedWindow(5)

	// 模拟正常心跳（间隔约1000ms）
	w.Add(1000)
	w.Add(2000) // 1000ms
	w.Add(3000) // 1000ms
	w.Add(4000) // 1000ms
	w.Add(5000) // 1000ms

	// 统计异常短间隔（<100ms）- 正常情况下应该是0
	shortCount := w.CountIntervalsBetween(0, 100)
	if shortCount != 0 {
		t.Errorf("Normal intervals short count: got %d, want 0", shortCount)
	}

	// 模拟加速器（间隔只有50ms）
	w.Reset()
	w.Add(1000)
	w.Add(1050) // 50ms
	w.Add(1100) // 50ms
	w.Add(1150) // 50ms
	w.Add(1200) // 50ms

	// 统计异常短间隔（<100ms）- 应该检测到4次
	shortCount = w.CountIntervalsBetween(0, 100)
	if shortCount != 4 {
		t.Errorf("Accelerated intervals short count: got %d, want 4", shortCount)
	}

	// 模拟混合场景：一半正常，一半加速
	w.Reset()
	w.Add(1000)
	w.Add(1500) // 500ms (正常)
	w.Add(1530) // 30ms (加速)
	w.Add(2030) // 500ms (正常)
	w.Add(2060) // 30ms (加速)

	// 统计加速间隔（<100ms）- 应该检测到2次
	shortCount = w.CountIntervalsBetween(0, 100)
	if shortCount != 2 {
		t.Errorf("Mixed intervals short count: got %d, want 2", shortCount)
	}
}

// TestSpeedWindowConcurrent 测试并发安全性
func TestSpeedWindowConcurrent(t *testing.T) {
	w := NewSpeedWindow(10)
	var wg sync.WaitGroup
	numGoroutines := 100

	// 并发添加
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func(i int) {
			defer wg.Done()
			w.Add(int64(i * 100))
		}(i)
	}
	wg.Wait()

	// 并发读取
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			_ = w.Size()
			_ = w.CountIntervalsBetween(0, 1000)
		}()
	}
	wg.Wait()

	// 验证最终状态
	if w.Size() != 10 {
		t.Errorf("Final size: got %d, want 10", w.Size())
	}
}
