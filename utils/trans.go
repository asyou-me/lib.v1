package utils

import (
	"strconv"
)

// 字符串转换为 int64
func StringToInt64(value string, def int64) int64 {
	i, err := strconv.Atoi(value)
	if err != nil {
		return def
	}
	return int64(i)
}

// 字符串转换成 int
func StringToInt(value string, def int) int {
	i, err := strconv.Atoi(value)
	if err != nil {
		return def
	}
	return i
}
