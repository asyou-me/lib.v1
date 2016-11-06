package errors

import "runtime"

// Resiger 注册错误类型
func Resiger(code int, err *ErrStruct) {
	defaultCodes.Resiger(code, err)
}

// New 创建一个错误容器，不记录错误路径
func New(code int, values ...string) *ErrStruct {
	return defaultCodes.New(code, values...)
}

// NewWithPath 创建一个错误容器，记录错误路径
func NewWithPath(code int, values ...string) *ErrStruct {
	pc, file, line, _ := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	e := defaultCodes.New(code, values...)
	e.Log = &LogStruct{
		Local: file + ":" + f.Name(),
		Line:  line,
	}
	return e
}
