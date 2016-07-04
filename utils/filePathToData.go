package utils

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

func FromYaml(path string, v interface{}) error {
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

func FromJson(path string, v interface{}) error {
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
