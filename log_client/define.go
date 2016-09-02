package log_client

import (
	"github.com/asyou-me/lib.v1/pulic_type"
)

// 定义日志等级常量
const (
	FatalLevel = iota
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
)

// 日志等级
type Level uint8

// 获取日志等级的文字描述
func (level Level) ToString() string {
	switch level {
	case DebugLevel:
		return "debug"
	case InfoLevel:
		return "info"
	case WarnLevel:
		return "warning"
	case ErrorLevel:
		return "error"
	case FatalLevel:
		return "fatal"
	}
	return "unknown"
}

// 日志后端接口,所有日志的对象都需要实现这个接口
type LogInterface interface {
	// 将日志写入服务,写入的是一个对象
	WriteTo(pulic_type.LogBase)
	// 此方法用于检查日志服务的可用性
	CheckHealth() bool
	// 此方法用于将一条将被恢复的消息重新写入日志服务
	RecoveryTo(string)
}
