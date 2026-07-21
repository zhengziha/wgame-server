package main_thread

import (
	"sync"

	"wgame-server/comm/log"
)

// 单 goroutine 串行执行业务逻辑，参考 hero_story.go_server/comm/main_thread/process.go。
// 这是同步骨架：所有 handler、异步回调、广播都进入此队列执行，
// 因此业务代码无需再加锁即可安全访问共享数据。
var (
	mainQ   = make(chan func(), 2048)
	started = false
	mu      sync.Mutex
)

// Process 将 task 投递到主线程队列异步执行。
// 首次调用时会懒加载启动消费 goroutine。
//
// 队列满时不会 panic（避免拖垮整个进程），改为：
//   - 记一条 Error 日志（带速率限制，避免日志风暴）
//   - 丢弃当前 task
//
// 队列满通常意味着业务 handler 平均耗时 > 帧到达间隔，
// 应通过指标观察并优化 handler，而不是崩溃进程。
func Process(task func()) {
	ensureStarted()
	select {
	case mainQ <- task:
	default:
		onDropped(task)
	}
}

// ProcessSync 将 task 投递到主线程并阻塞等待其执行完毕。
// 用于在 IO goroutine 中需要等待结果时调用。
//
// 注意：若在 main_thread 自身 goroutine 内调用 ProcessSync 会死锁。
func ProcessSync(task func()) {
	done := make(chan struct{})
	Process(func() {
		defer close(done)
		task()
	})
	<-done
}

func ensureStarted() {
	if started {
		return
	}
	mu.Lock()
	defer mu.Unlock()
	if started {
		return
	}
	started = true
	go execute()
}

func execute() {
	for {
		select {
		case task, ok := <-mainQ:
			if !ok {
				return
			}
			safeInvoke(task)
		}
	}
}

func safeInvoke(task func()) {
	defer func() {
		if r := recover(); r != nil {
			log.Error("[main_thread] task panic recovered: %v", r)
		}
	}()
	task()
}

// ---- 队列满丢弃的日志限流 ----

var (
	dropMu      sync.Mutex
	dropLastLog int64 // unix milli；为简化不引入 time，改用 atomic
	dropCount   int64 // 自上次日志以来的丢弃次数
)

func onDropped(_ func()) {
	dropMu.Lock()
	dropCount++
	now := nowMilli()
	// 每 5 秒最多打一条日志，聚合统计
	if now-dropLastLog >= 5000 || dropLastLog == 0 {
		n := dropCount
		dropCount = 0
		dropLastLog = now
		dropMu.Unlock()
		log.Error("[main_thread] queue full, dropped %d tasks in last window (queue size=2048)", n)
		return
	}
	dropMu.Unlock()
}
