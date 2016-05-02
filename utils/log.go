package utils

import (
	"github.com/asyoume/lib/log_client"
)

func GetDefaultLog() *log_client.Logger {
	Log, _ := log_client.New(&[]log_client.LogConf{
		log_client.LogConf{
			Type:   "console",
			Spare:  false,
			Weight: 1,
		},
	})
	return Log
}
