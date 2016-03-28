package log_client

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// 格式日志对象到json
func jsonFormat(data LogBase) []byte {
	serialized, err := json.Marshal(data)
	if err != nil {
		return nil
	}
	return serialized
}

// 按权重排序配置文件 (降序排列)
func LogConfSort(array []LogConf) {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j].Weight < array[j+1].Weight {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
}

//转换成绝对路径并验证文件是否存在
func file_path_check(path string) (os.FileInfo, error) {
	path = strings.Replace(path, " ", "", -1)
	if string(path[0]) != "/" {
		curr_path, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			return nil, err
		}
		path = curr_path + "/" + path
	}

	fileinfo, isExist := Exist(path)
	if isExist != nil {
		return nil, isExist
	}
	return fileinfo, nil
}

// 检查文件或目录是否存在
// 如果由 filename 指定的文件或目录存在则返回 true，否则返回 false
func Exist(filename string) (os.FileInfo, error) {
	fileinfo, err := os.Stat(filename)
	if err == nil {
		return fileinfo, nil
	}
	return nil, err
}

//  切分ip和端口
func IpPort(filename string) (string, int, error) {
	strs := strings.Split(filename, ":")
	if len(strs) != 2 {
		return "", 0, errors.New("参数不符合xxx.xxx.xxx.xxx:port")
	}
	// 转换端口号为数字
	port, err := strconv.Atoi(strs[1])
	if err != nil {
		return "", 0, err
	}

	return strs[0], port, err
}
