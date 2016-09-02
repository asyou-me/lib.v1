package log_client

import (
	"fmt"
	"os"

	"github.com/asyou-me/lib.v1/pulic_type"
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

var _base_log *ConsoleHandle

func init() {
	// 默认自动初始化到当前目录
	SetBaseLog(pulic_type.LogConf{"", "", "",
		"", "console", true, 1})
}

func SetBaseLog(conf pulic_type.LogConf) {
	var err error
	_base_log, err = NewConsoleHandle(conf, &Logger{})
	if err != nil {
		fmt.Println(err)
		fmt.Println("\033[31;1mPanic：基础日志无法使用")
		os.Exit(2)
	}
}

// 创建日志对象
func New(conf *[]pulic_type.LogConf) (*Logger, error) {
	l := &Logger{}
	err := l.Init(conf)
	if err != nil {
		return nil, err
	}
	return l, nil
}
