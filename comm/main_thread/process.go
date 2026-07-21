package main_thread

import "sync"

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
func Process(task func()) {
	ensureStarted()
	select {
	case mainQ <- task:
	default:
		// 队列满了直接丢弃并打印（避免阻塞投递方）。
		// 也可改为阻塞或回退策略，按业务需求调整。
		panic("main_thread queue full")
	}
}

// ProcessSync 将 task 投递到主线程并阻塞等待其执行完毕。
// 用于在 IO goroutine 中需要等待结果时调用。
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
			// 简单 panic 恢复，避免单个 task 异常拖垮主线程
			_ = r
		}
	}()
	task()
}
