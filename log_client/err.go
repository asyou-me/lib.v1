package log_client

type logErr struct {
	Msg   string `json:"msg"`
	Err   string `json:"err"`
	Level string `json:"level"`
	Time  int64  `json:"time"`
}

func (l *logErr) GetLevel() string {
	return ""
}

func (l *logErr) SetLevel(level string) {
	l.Level = level
}

func (l *logErr) SetTime(t int64) {
	l.Time = t
}
