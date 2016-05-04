package log_client

// 日志后端接口,所有日志的对象都需要实现这个接口
type LogInterface interface {
	// 将日志写入服务,写入的是一个对象
	WriteTo(LogBase)
	// 此方法用于检查日志服务的可用性
	CheckHealth() bool
	// 此方法用于将一条将被恢复的消息重新写入日志服务
	RecoveryTo(string)
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
