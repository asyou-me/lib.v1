package pulic_type

import (
	log_client_types "github.com/asyoume/lib/log_client/types"
)

//日志模块接口定义
type Loger interface {
	Debug(...log_client_types.LogBase)
	Info(...log_client_types.LogBase)
	Print(...log_client_types.LogBase)
	Warn(...log_client_types.LogBase)
	Warning(...log_client_types.LogBase)
	Error(...log_client_types.LogBase)
	Fatal(...log_client_types.LogBase)
	Panic(...log_client_types.LogBase)
}
