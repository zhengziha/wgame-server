package log

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
)

// 全局日志对象，参考 hero_story.go_server/comm/log/log.go
var (
	once    sync.Once
	infoLgr *log.Logger
	errLgr  *log.Logger
)

// Init 初始化日志系统。logFilePath 形如 /var/log/wgame/biz_server.log，
// 函数内部基于该路径拆分目录与文件前缀，并按天滚动。
func Init(logFilePath string) {
	once.Do(func() {
		dir := filepath.Dir(logFilePath)
		prefix := filepath.Base(logFilePath)
		if ext := filepath.Ext(prefix); ext != "" {
			prefix = prefix[:len(prefix)-len(ext)]
		}
		w := newDailyFileWriter(dir, prefix)
		infoLgr = log.New(w, "[INFO] ", log.Ldate|log.Lmicroseconds|log.Lshortfile)
		errLgr = log.New(w, "[ERROR] ", log.Ldate|log.Lmicroseconds|log.Lshortfile)
	})
}

// Info 输出 info 级别日志，format 用法与 fmt.Printf 一致
func Info(format string, args ...interface{}) {
	if infoLgr == nil {
		infoLgr = log.New(os.Stdout, "[INFO] ", log.Ldate|log.Lmicroseconds|log.Lshortfile)
	}
	infoLgr.Output(3, fmt.Sprintf(format, args...))
}

// Error 输出 error 级别日志
func Error(format string, args ...interface{}) {
	if errLgr == nil {
		errLgr = log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Lmicroseconds|log.Lshortfile)
	}
	errLgr.Output(3, fmt.Sprintf(format, args...))
}
