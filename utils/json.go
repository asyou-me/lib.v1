package utils

import (
	"encoding/json"
)

// json 序列化为字符串
func JsonToStr(v interface{}) (*string, error) {
	buf, err := json.Marshal(v)
	var str string
	if err != nil {
		return nil, err
	}
	str = *BytesToStr(buf)
	return &str, nil
}

// json字符串 反序列化为结构
func StrToJson(data *string, v interface{}) error {
	return json.Unmarshal([]byte(*data), v)
}
