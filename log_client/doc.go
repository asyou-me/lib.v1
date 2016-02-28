package log_client

/*
var err error
utils.Log, err = log_client.New(&[]log_client.LogConf{
  // 主log服务 192.168.1.131
  log_client.LogConf{"127.0.0.1:6379", "cf_log_queue", "", "", "redis", false, 10},
  // 备用log服务
  log_client.LogConf{"/home/xiaobai/dumps", "cf_log", "", "", "file", true, 1},
})

// 当日志服务不用取消启动服务器
if err != nil {
  fmt.Println(err)
  os.Exit(2)
}

func log_test() {
  for {
    for i := 0; i < 30000; i++ {
      go func() {
        utils.Log.Debug(&log_err{
          Err: "",
          Msg: "检查redis服务器服务无法使用()",
        })
      }()
    }
    time.Sleep(1 * time.Second)
  }
}

type log_err struct {
  Msg   string `json:"msg"`
  Err   string `json:"err"`
  Level string `json:"level"`
  Time  int64  `json:"time"`
}

func (l *log_err) GetLevel() string {
  return ""
}

func (l *log_err) SetLevel(level string) {
  l.Level = level
}

func (l *log_err) SetTime(t int64) {
  l.Time = t
}
*/
