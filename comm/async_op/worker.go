package async_op

import "wgame-server/comm/main_thread"

// worker 是一个独立的 goroutine + 任务队列。
// 同一 worker 内的任务严格串行，跨 worker 并行执行。
type worker struct {
	id      int
	taskQ   chan taskItem
}

type taskItem struct {
	asyncOp      func()
	continueWith func()
}

func newWorker(id int) *worker {
	w := &worker{
		id:    id,
		taskQ: make(chan taskItem, 2048),
	}
	go w.run()
	return w
}

func (w *worker) enqueue(asyncOp func(), continueWith func()) {
	select {
	case w.taskQ <- taskItem{asyncOp: asyncOp, continueWith: continueWith}:
	default:
		// 队列满：此处选择阻塞调用方以保证不丢任务。
		// 若希望像 Java 版那样“丢弃并报警”，可替换为 panic / log。
		w.taskQ <- taskItem{asyncOp: asyncOp, continueWith: continueWith}
	}
}

func (w *worker) run() {
	for item := range w.taskQ {
		safeRun(item.asyncOp)
		if item.continueWith != nil {
			cw := item.continueWith
			// 把回调投回主线程串行执行
			main_thread.Process(cw)
		}
	}
}

func safeRun(f func()) {
	defer func() {
		if r := recover(); r != nil {
			_ = r
		}
	}()
	f()
}
