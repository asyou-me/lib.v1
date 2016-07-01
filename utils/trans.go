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

// 字符串转换为 uint64
func StringToUint64(value string, def uint64) uint64 {
	i, err := strconv.ParseUint(value, 10, 0)
	if err != nil {
		return def
	}
	return i
}

// 字符串转换成 int
func StringToInt(value string, def int) int {
	i, err := strconv.Atoi(value)
	if err != nil {
		return def
	}
	return i
}
