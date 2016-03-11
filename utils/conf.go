package utils

import (
	"encoding/json"
	"errors"
	"github.com/asyoume/lib/pulic_type"
	"io/ioutil"
	"os"
)

func ConfigInit(path string) (*pulic_type.ConfType, error) {
	fi, err := os.Open(path)
	var obj pulic_type.ConfType = pulic_type.ConfType{}

	if err != nil {
		return &obj, errors.New("配置文件" + path + "不存在")
	}

	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	if err != nil {
		return &obj, err
	}

	if err = json.Unmarshal(fd, &obj); err != nil {
		if err != nil {
			return &obj, err
		}
	}
	return &obj, nil
}
