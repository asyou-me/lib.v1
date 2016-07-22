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

func GetString(value interface{}) string {
	if value == nil {
		return ""
	}
	data, ok := value.(string)
	if ok {
		return data
	}
	return ""
}

func GetInt64(value interface{}) int64 {
	if value == nil {
		return 0
	}
	data, ok := value.(int64)
	if ok {
		return data
	}
	return 0
}
