package log_client

type Loggerr struct {
	Msg   string `json:"msg"`
	Err   string `json:"err"`
	Level string `json:"level"`
	Time  int64  `json:"time"`
}

func (l *Loggerr) GetLevel() string {
	return ""
}

func (l *Loggerr) SetLevel(level string) {
	l.Level = level
}

func (l *Loggerr) SetTime(t int64) {
	l.Time = t
}
