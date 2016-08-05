package utils

import "encoding/json"

// json 序列化为字符串
func JsonToStr(v interface{}) (*string, error) {
	buf, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return BytesToStr(buf), nil
}

// json 序列化为字符串
func StrToJson(data *string, v interface{}) error {
	return json.Unmarshal(StrToBytes(data), v)
}
