package log_client

import (
	"fmt"
	"time"

	"github.com/asyou-me/lib.v1/pulic_type"
)

// 创建一个文档日志处理对象
func NewConsoleHandle(conf pulic_type.LogConf, log *Logger) (*ConsoleHandle, error) {
	clog := ConsoleHandle{
		log: log,
	}
	return &clog, nil
}

// 日志处理对象
type ConsoleHandle struct {
	log *Logger
}

// 日志对象健康检查
func (r *ConsoleHandle) CheckHealth() bool {
	return true
}

// 日志处理句柄
func (l *ConsoleHandle) WriteTo(msg pulic_type.LogBase) {
	var NowTime = time.Now().Unix()

	msg.SetTime(NowTime)
	fmt.Println(msg)

	msg = nil
}

// 文档日志处理句柄
func (l *ConsoleHandle) RecoveryTo(msg string) {
	fmt.Println(msg)
	msg = ""
}
