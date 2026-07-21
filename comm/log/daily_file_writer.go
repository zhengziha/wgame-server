package log

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// dailyFileWriter 按天滚动日志文件
// 参考 hero_story.go_server/comm/log/daily_file_writer.go
type dailyFileWriter struct {
	mu sync.Mutex

	logDir     string
	filePrefix string

	currentDate string
	fileHandle  *os.File
}

// newDailyFileWriter 创建一个按天滚动的文件 writer
// logDir 为目录，filePrefix 为文件名前缀（最终文件名为 <prefix>_YYYYMMDD.log）
func newDailyFileWriter(logDir, filePrefix string) *dailyFileWriter {
	return &dailyFileWriter{
		logDir:     logDir,
		filePrefix: filePrefix,
	}
}

func (w *dailyFileWriter) Write(p []byte) (int, error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	if err := w.rotateIfNeeded(); err != nil {
		return 0, err
	}
	return w.fileHandle.Write(p)
}

func (w *dailyFileWriter) rotateIfNeeded() error {
	today := time.Now().Format("20060102")
	if today == w.currentDate && w.fileHandle != nil {
		return nil
	}

	if w.fileHandle != nil {
		_ = w.fileHandle.Close()
		w.fileHandle = nil
	}

	if err := os.MkdirAll(w.logDir, 0o755); err != nil {
		return err
	}

	filename := filepath.Join(w.logDir, fmt.Sprintf("%s_%s.log", w.filePrefix, today))
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	w.fileHandle = f
	w.currentDate = today
	return nil
}

func (w *dailyFileWriter) Close() error {
	w.mu.Lock()
	defer w.mu.Unlock()
	if w.fileHandle != nil {
		err := w.fileHandle.Close()
		w.fileHandle = nil
		return err
	}
	return nil
}

// 保证 *dailyFileWriter 实现 io.WriteCloser
var _ io.WriteCloser = (*dailyFileWriter)(nil)
