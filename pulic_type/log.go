package pulic_type

import (
	"encoding/json"
	"fmt"
)

// 日志初始化配置文件数据类型
type LogConf struct {
	// mysql -> tcp(xxx.xxx.xxx.xxx:port)  redis -> xxx.xxx.xxx.xxx:port  file -> /var/log/xxxx
	Addr string `json:"addr" yaml:"addr"`
	// mysql -> 表名    redis -> list的key名   file -> 目录名
	Area string `json:"area" yaml:"area"`
	// mysql -> 数据库账户   redis -> redis账户名  file -> -
	Auth_id string `json:"auth_id" yaml:"auth_id"`
	// mysql -> 数据库账户密码  redis -> redis账户密码  file -> -
	Auth_Secret string `json:"auth_secret" yaml:"auth_secret"`
	// 服务的类型(redis,file)
	Type string `json:"type" yaml:"type"`
	// 服务是否为备用
	Spare bool `json:"spare" yaml:"spare"`
	// 服务的权重，权重越大，越优先启动
	Weight int64 `json:"weight" yaml:"weight"`
}

//日志模块接口定义
type Logger interface {
	Debug(...LogBase)
	Info(...LogBase)
	Print(...LogBase)
	Warn(...LogBase)
	Warning(...LogBase)
	Error(...LogBase)
	Fatal(...LogBase)
	Panic(...LogBase)
	Log(string, ...LogBase)
}

// 日志基础数据接口
type LogBase interface {
	// 获取当前日志的等级
	GetLevel() string
	// 设置当前日志的等级
	SetLevel(string)
	// 设置这条日志的记录时间
	SetTime(int64)
}

type DefalutLogger struct {
}

func (this *DefalutLogger) Log(level string, values ...LogBase) {
	for _, v := range values {
		fmtJSON(v)
	}
}

func (this *DefalutLogger) Debug(values ...LogBase) {
	for _, v := range values {
		fmtJSON(v)
	}
}

func (this *DefalutLogger) Info(values ...LogBase) {
	for _, v := range values {
		fmtJSON(v)
	}
}

func (this *DefalutLogger) Print(values ...LogBase) {
	for _, v := range values {
		fmtJSON(v)
	}
}

func (this *DefalutLogger) Warn(values ...LogBase) {
	for _, v := range values {
		fmtJSON(v)
	}
}
func (this *DefalutLogger) Warning(values ...LogBase) {
	for _, v := range values {
		fmtJSON(v)
	}
}

func (this *DefalutLogger) Error(values ...LogBase) {
	for _, v := range values {
		fmtJSON(v)
	}
}

func (this *DefalutLogger) Fatal(values ...LogBase) {
	for _, v := range values {
		fmtJSON(v)
	}
}

func (this *DefalutLogger) Panic(values ...LogBase) {
	for _, v := range values {
		fmtJSON(v)
	}
}

func fmtJSON(v LogBase) {
	data, err := json.Marshal(v)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
}
