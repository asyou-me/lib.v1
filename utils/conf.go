package utils

import (
	"encoding/json"
	"errors"
	"github.com/asyoume/lib/pulic_type"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// 解析公用配置文件
func ConfigInit(path string) (*pulic_type.ConfType, error) {
	var obj *pulic_type.ConfType = new(pulic_type.ConfType)

	err := JsonConf(path, obj)
	return obj, err
}

func YamlConf(path string, v interface{}) error {
	fi, err := os.Open(path)

	if err != nil {
		return errors.New("传入的配置文件路径: " + path + " 不存在")
	}

	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	if err != nil {
		return err
	}

	if err = yaml.Unmarshal(fd, v); err != nil {
		if err != nil {
			return err
		}
	}
	return nil
}

func JsonConf(path string, v interface{}) error {
	fi, err := os.Open(path)

	if err != nil {
		return errors.New("传入的配置文件路径: " + path + " 不存在")
	}

	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(fd, v); err != nil {
		if err != nil {
			return err
		}
	}
	return nil
}
