package log_client

import (
	"testing"
	"time"
)

var err error
var logger *Logger

func Test(t *testing.T) {

	logger, err = New(&[]LogConf{
		// 主log服务 192.168.1.131
		LogConf{"127.0.0.1:6379", "cf_log_queue", "", "", "redis", false, 10},
		// 备用log服务
		LogConf{"/home/xiaobai/dumps", "cf_log", "", "", "file", true, 1},
	})
	if err != nil {
		panic(err)
	}

	go log_test()
	for {
		time.Sleep(10 * time.Second)
	}
}

func log_test() {
	for {
		for i := 0; i < 30000; i++ {
			go func() {
				logger.Debug(&log_err{
					Err: "",
					Msg: "检查redis服务器服务无法使用()",
				})
			}()
		}
		time.Sleep(1 * time.Second)
	}
}
