package log_client

import (
	"github.com/asyoume/lib/pulic_type"
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

// 日志初始化配置文件数据类型
type LogConf struct {
	// mysql -> tcp(xxx.xxx.xxx.xxx:port)  redis -> xxx.xxx.xxx.xxx:port  file -> /var/log/xxxx
	Addr string `json:"addr" yaml:"addr"`
	// mysql -> 表名    redis -> list的key名   file -> 目录名
	Area string `json:"area" yaml:"area"`
	// mysql -> 数据库账户   redis -> redis账户名  file -> -
	Auth_id string `json:"auth_id" yaml:"auth_id"`
	// mysql -> 数据库账户密码  redis -> redis账户密码  file -> -
	Auth_Secret string `json:"auth_secret" yaml:"auth_secret"`
	// 服务的类型(redis,file)
	Type string `json:"type" yaml:"type"`
	// 服务是否为备用
	Spare bool `json:"spare" yaml:"spare"`
	// 服务的权重，权重越大，越优先启动
	Weight int64 `json:"weight" yaml:"weight"`
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

// 日志基础数据接口
type LogBase interface {
	// 获取当前日志的等级
	GetLevel() string
	// 设置当前日志的等级
	SetLevel(string)
	// 设置这条日志的记录时间
	SetTime(int64)
}
