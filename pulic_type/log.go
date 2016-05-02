package pulic_type

import (
	"github.com/asyoume/lib/log_client"
)

//日志模块接口定义
type Loger interface {
	Debug(...log_client.LogBase)
	Info(...log_client.LogBase)
	Print(...log_client.LogBase)
	Warn(...log_client.LogBase)
	Warning(...log_client.LogBase)
	Error(...log_client.LogBase)
	Fatal(...log_client.LogBase)
	Panic(...log_client.LogBase)
}
