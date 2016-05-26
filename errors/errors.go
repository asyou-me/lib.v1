package errors

import (
	"runtime"
)

func Resiger(code int, err *ErrStruct) {
	defaultCodes.Resiger(code, err)
}

func New(code int, values ...string) *ErrStruct {
	return defaultCodes.New(code, "", values...)
}

func NewWithPath(code int, values ...string) *ErrStruct {
	pc, file, _, _ := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	return defaultCodes.New(code, file+":"+f.Name(), values...)
}
