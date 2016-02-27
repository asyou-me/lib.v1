package log

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

var file_local_log *FileHandle

func init() {
	// 默认自动初始化到当前目录
	SetLogLog(LogConf{"./", "_log_log", "",
		"", "file", true, 1})
}

func SetLogLog(conf LogConf) {
	var err error
	file_local_log, err = NewFileHandle(conf, &Logger{})
	if err != nil {
		fmt.Println(err)
		fmt.Println("\033[31;1mPanic：文件日志无法使用")
		os.Exit(2)
	}
}

//创建日志对象
func New(conf *[]LogConf) (*Logger, error) {
	l := &Logger{}
	err := l.Init(conf)
	if err != nil {
		return nil, err
	}
	return l, nil
}
