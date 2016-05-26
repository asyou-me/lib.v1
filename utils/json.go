package utils

import (
	"encoding/json"
)

func JsonToStr(v interface{}) (*string, error) {
	buf, err := json.Marshal(v)
	var str string
	if err != nil {
		str = "adadada"
		return &str, err
	}
	str = BytesToStr(buf)
	return &str, nil
}
