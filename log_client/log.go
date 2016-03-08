package log_client

import (
	"fmt"
	"os"
)

var (
	ReconnectInterval = 5
	FileTimeFormat    = "2006-01-02"
)

/*
测试环境
C P U:i5 6600
内 存:8G DDR4 2144
硬 盘:ssd 400mb/s

使用file后端  6w/s写入速度  内存占用170M  cpu 8%
使用redis后端  6w/s写入速度  内存占用170M  cpu 28%
*/

var _base_log *FileHandle

func init() {
	// 默认自动初始化到当前目录
	SetBaseLog(LogConf{"./", "_base_log", "",
		"", "file", true, 1})
}

func SetBaseLog(conf LogConf) {
	var err error
	_base_log, err = NewFileHandle(conf, &Logger{})
	if err != nil {
		fmt.Println(err)
		fmt.Println("\033[31;1mPanic：文件日志无法使用")
		os.Exit(2)
	}
}

// 创建日志对象
func New(conf *[]LogConf) (*Logger, error) {
	l := &Logger{}
	err := l.Init(conf)
	if err != nil {
		return nil, err
	}
	return l, nil
}

// 日志后端接口,所有日志的对象都需要实现这个接口
type LogInterface interface {
	// 将日志写入服务,写入的是一个对象
	WriteTo(LogBase)
	// 此方法用于检查日志服务的可用性
	CheckHealth() bool
	// 此方法用于将一条将被恢复的消息重新写入日志服务
	RecoveryTo(string)
}
