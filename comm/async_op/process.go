package async_op

import (
	"hash/crc32"
	"sync"
)

// 异步操作 worker 池，参考 hero_story.go_server/comm/async_op/worker.go。
// 共 workerCount 个 worker，每个 worker 拥有独立的任务队列；
// 通过 bindId % workerCount 选取 worker，保证同一 bindId 的任务串行执行。

const workerCount = 2048

var (
	workerArray [workerCount]*worker
	workerMu    sync.Mutex
)

// Process 把异步操作投递给由 bindId 决定的 worker。
// asyncOp 在 worker goroutine 中执行；
// continueWith（非 nil 时）会被回投到主线程串行执行。
func Process(bindId int, asyncOp func(), continueWith func()) {
	w := getWorker(bindId)
	w.enqueue(asyncOp, continueWith)
}

func getWorker(bindId int) *worker {
	idx := bindId % workerCount
	if idx < 0 {
		idx = -idx
	}
	w := workerArray[idx]
	if w != nil {
		return w
	}
	workerMu.Lock()
	defer workerMu.Unlock()
	if workerArray[idx] == nil {
		workerArray[idx] = newWorker(idx)
	}
	return workerArray[idx]
}

// StrToBindId 把字符串（例如用户名）转换为稳定的 worker 分片 id。
// 使用 CRC32，保证同一字符串始终落到同一个 worker。
func StrToBindId(s string) int {
	h := int(crc32.ChecksumIEEE([]byte(s)))
	if h < 0 {
		h = -h
	}
	return h
}
