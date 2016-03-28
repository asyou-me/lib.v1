package utils

import (
	"os"
	"path/filepath"
)

// 将相对路径转换成绝对路径
func CompletePath(conf_path *string) {
	if len((*conf_path)) > 0 && string((*conf_path)[0]) != "/" {
		var curr_path string = ""
		curr_path, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			panic(err)
		}
		*conf_path = curr_path + "/" + *conf_path
	}
}
