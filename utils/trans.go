package utils

import (
	"strconv"
)

func StringToInt64(value string, def int64) int64 {
	i, err := strconv.Atoi(value)
	if err != nil {
		return def
	}
	return int64(i)
}

func StringToInt(value string, def int) int {
	i, err := strconv.Atoi(value)
	if err != nil {
		return def
	}
	return i
}
