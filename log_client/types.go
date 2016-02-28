package log_client

import ()

// 定义日志等级常量
const (
	FatalLevel = iota
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
)

// 日志后端对象接口定义
type LogInterface interface {
	Do(LogBase)
	Check() bool
	Recovery(string)
}

// 日志初始化配置文件
type LogConf struct {
	Addr        string `json:"addr"`
	Area        string `json:"area"`
	Auth_id     string `json:"auth_id"`
	Auth_Secret string `json:"auth_secret"`
	Type        string `json:"type"`
	Spare       bool   `json:"spare"`
	Weight      int64  `json:"weight"`
}

// 日志基础数据类型
type LogBase interface {
	GetLevel() string
	SetLevel(string)
	SetTime(int64)
}

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
