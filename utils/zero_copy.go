package utils

import (
	"unsafe"
)

// 字符串转换成字节
func StrToBytes(s *string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

// 字节转换成字符串
func BytesToStr(b []byte) *string {
	return (*string)(unsafe.Pointer(&b))
}
