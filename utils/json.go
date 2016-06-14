package utils

import (
	"encoding/json"
)

func JsonToStr(v interface{}) (*string, error) {
	buf, err := json.Marshal(v)
	var str string
	if err != nil {
		return nil, err
	}
	str = *BytesToStr(buf)
	return &str, nil
}
