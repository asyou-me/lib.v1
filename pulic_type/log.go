package pulic_type

import (
	"fmt"
)

//日志模块接口定义
type Logger interface {
	Debug(...LogBase)
	Info(...LogBase)
	Print(...LogBase)
	Warn(...LogBase)
	Warning(...LogBase)
	Error(...LogBase)
	Fatal(...LogBase)
	Panic(...LogBase)
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

type DefalutLogger struct {
}

func (this *DefalutLogger) Debug(values ...LogBase) {
	fmt.Println(values)
}
func (this *DefalutLogger) Info(values ...LogBase) {
	fmt.Println(values)
}
func (this *DefalutLogger) Print(values ...LogBase) {
	fmt.Println(values)
}
func (this *DefalutLogger) Warn(values ...LogBase) {
	fmt.Println(values)
}
func (this *DefalutLogger) Warning(values ...LogBase) {
	fmt.Println(values)
}
func (this *DefalutLogger) Error(values ...LogBase) {
	fmt.Println(values)
}
func (this *DefalutLogger) Fatal(values ...LogBase) {
	fmt.Println(values)
}
func (this *DefalutLogger) Panic(values ...LogBase) {
	fmt.Println(values)
}
