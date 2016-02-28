package log_client

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
