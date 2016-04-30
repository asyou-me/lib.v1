package log_client

import (
	"fmt"
	"time"
)

// 创建一个终端日志处理对象
func NewConsoleHandle(conf LogConf, log *Logger) (*ConsoleHandle, error) {
	flog := ConsoleHandle{}
	return &flog, nil
}

// 终端日志处理对象
type ConsoleHandle struct {
}

// 文档日志对象健康检查
func (r *ConsoleHandle) CheckHealth() bool {
	return true
}

// 文档日志处理句柄
func (l *ConsoleHandle) WriteTo(msg LogBase) {
	var NowTime = time.Now().Unix()

	msg.SetTime(NowTime)
	msgbyte := append(jsonFormat(msg), '\n')

	fmt.Println(string(msgbyte))
}

// 文档日志处理句柄
func (l *ConsoleHandle) RecoveryTo(msg string) {
}
