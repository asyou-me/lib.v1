package utils

import (
	"github.com/asyoume/lib.v1/pulic_type"
)

// 解析公用配置文件
func ConfigInit(path string) (*pulic_type.ConfType, error) {
	var obj *pulic_type.ConfType = new(pulic_type.ConfType)

	err := JsonConf(path, obj)
	return obj, err
}

// yaml 格式配置文件解析
func YamlConf(path string, v interface{}) error {
	return FromYaml(path, v)
}

// json 格式配置文件解析
func JsonConf(path string, v interface{}) error {
	return FromJson(path, v)
}
