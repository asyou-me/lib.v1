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

// 日志初始化配置文件数据类型
type LogConf struct {
	// mysql -> tcp(xxx.xxx.xxx.xxx:port)  redis -> xxx.xxx.xxx.xxx:port  file -> /var/log/xxxx
	Addr string `json:"addr"`
	// mysql -> 表名    redis -> list的key名   file -> 目录名
	Area string `json:"area"`
	// mysql -> 数据库账户   redis -> redis账户名  file -> -
	Auth_id string `json:"auth_id"`
	// mysql -> 数据库账户密码  redis -> redis账户密码  file -> -
	Auth_Secret string `json:"auth_secret"`
	// 服务的类型(redis,file)
	Type        string `json:"type"`
	// 服务是否为备用
	Spare       bool   `json:"spare"`
	// 服务的权重，权重越大，越优先启动
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
