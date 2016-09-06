package log_client

import (
	"errors"
	"fmt"
	"time"

	"github.com/asyou-me/lib.v1/pulic_type"
)

// 日志对象
type Logger struct {
	// 配置文件
	Conf     []pulic_type.LogConf
	CurrConf pulic_type.LogConf
	// 运行状态
	Run int64
	// 后端对象
	BackEnd      LogInterface
	BackEndTrash LogInterface
	// 重启动相关
	TransitionChannel chan bool
	lastTransition    int64

	//日志记录等级
	Level Level

	// 消息管道
	MsgChannel      chan pulic_type.LogBase
	RecoveryChannel chan string

	Err         chan error
	ErrNum      int64
	StopChannel chan bool
}

// 日志对象初始化流程
func (l *Logger) Init(conf *[]pulic_type.LogConf) error {

	// 据权重调整配置文件顺序
	LogConfSort(*conf)

	l.Conf = *conf
	l.BackEndTrash = nil

	// 日志列表
	l.MsgChannel = make(chan pulic_type.LogBase, 1000)
	l.StopChannel = make(chan bool)
	l.Err = make(chan error, 100)
	l.TransitionChannel = make(chan bool, 1)

	l.Level = DebugLevel
	// 选择一个可用的日志后端
	l.Transition(false)
	// 开启日志处理携程
	go func() {
		l.Consumer()
	}()
	if l.Run == 0 {
		return errors.New("初始化日志模块出错,没有可用的日志服务")
	}
	return nil
}

// 日志对象消费者
func (l *Logger) Consumer() {
	var news = l.MsgChannel
	var err = l.Err
	var stop = l.StopChannel
	var transition = l.TransitionChannel
	for {
		select {
		// log消息列表
		case log := <-news:
			//go func() {
			l.BackEnd.WriteTo(log)
			//}()

		// 错误消息列表
		case e := <-err:
			_baseLog.WriteTo(&Loggerr{
				Level: "WARN",
				Msg:   e.Error(),
				Time:  time.Now().Unix(),
			})

			l.ErrNum = l.ErrNum + 1
			var now_unix_time = time.Now().Unix()
			if l.ErrNum > 100 && now_unix_time-l.lastTransition > 5 {
				l.lastTransition = now_unix_time
				go func() {
					if len(transition) == 0 {
						transition <- true
					}
				}()
			}

		// 服务变更通道
		case _ = <-transition:
			ok := l.BackEnd.CheckHealth()
			if ok {
				l.Run = 1
				l.ErrNum = 0
			} else {
				l.Run = 0
			}
			// 低优先和当前服务不可用时
			if !ok || l.CurrConf.Spare == true {
				// 检查高优先服务的可用性
				l.Transition(false)
				if l.Run == 0 && !ok {
					//暂时低优先服务
					l.Transition(true)
				}
			}

		case s := <-stop:
			if s == true {
				break
			}
		}
	}
}

//选择最优的日志远端
func (l *Logger) Transition(spare bool) {
	var err error
	for i := 0; i < len(l.Conf); i++ {

		if l.Conf[i].Addr == l.CurrConf.Addr &&
			l.Conf[i].Area == l.CurrConf.Area &&
			l.Conf[i].Spare == l.CurrConf.Spare &&
			l.Conf[i].Auth_id == l.CurrConf.Auth_id &&
			l.Conf[i].Auth_Secret == l.CurrConf.Auth_Secret {
			continue
		}

		// 主日志服务
		if !spare && l.Conf[i].Spare {
			continue
		}

		// 备日志服务
		if spare && !l.Conf[i].Spare {
			continue
		}

		switch l.Conf[i].Type {
		case "redis":
			var backend LogInterface
			backend, err = NewRedisHandle(l.Conf[i], l)
			if err == nil {
				l.CurrConf = l.Conf[i]
				l.ErrNum = 0
				l.BackEndTrash = l.BackEnd
				l.BackEnd = backend
				l.Run = 1
			} else {
				_baseLog.WriteTo(&Loggerr{
					Level: "ERROR",
					Err:   err.Error(),
					Msg:   "初始化redis服务器" + l.Conf[i].Addr + "失败",
					Time:  time.Now().Unix(),
				})
			}
			break
		case "file":
			var backend LogInterface
			backend, err = NewFileHandle(l.Conf[i], l)
			if err == nil {
				l.CurrConf = l.Conf[i]
				l.ErrNum = 0
				l.BackEndTrash = l.BackEnd
				l.BackEnd = backend
				l.Run = 1
			} else {
				_baseLog.WriteTo(&Loggerr{
					Level: "ERROR",
					Err:   err.Error(),
					Msg:   "文件日志" + l.Conf[i].Addr + "无法使用",
					Time:  time.Now().Unix(),
				})
			}
			break
		case "kafka":
			var backend LogInterface
			backend, err = NewKafkaHandle(l.Conf[i], l)
			if err == nil {
				l.CurrConf = l.Conf[i]
				l.ErrNum = 0
				l.BackEndTrash = l.BackEnd
				l.BackEnd = backend
				l.Run = 1
			} else {
				_baseLog.WriteTo(&Loggerr{
					Level: "ERROR",
					Err:   err.Error(),
					Msg:   "文件日志" + l.Conf[i].Addr + "无法使用",
					Time:  time.Now().Unix(),
				})
			}
			break
		case "console":
			var backend LogInterface
			backend, err = NewConsoleHandle(l.Conf[i], l)
			if err == nil {
				l.CurrConf = l.Conf[i]
				l.ErrNum = 0
				l.BackEndTrash = l.BackEnd
				l.BackEnd = backend
				l.Run = 1
			} else {
				_baseLog.WriteTo(&Loggerr{
					Level: "ERROR",
					Err:   err.Error(),
					Msg:   "终端日志" + l.Conf[i].Addr + "无法使用",
					Time:  time.Now().Unix(),
				})
			}
			break
		default:
			continue
		}

		if l.Run == 1 {
			if l.CurrConf.Spare == true {
				fmt.Println("{time:\"" + fmt.Sprint(time.Now().Unix()) + "\",msg:\"" + "日志进入低优先模式,系统会以" + fmt.Sprint(ReconnectInterval) +
					"秒为周期检查高权重服务的可用性\"}")
				_baseLog.WriteTo(&Loggerr{
					Level: "INFO",
					Msg:   "高权重服务不可用,日志进入备用模式",
					Time:  time.Now().Unix(),
				})
				go func() {
					time.Sleep(time.Duration(ReconnectInterval) * time.Second)
					l.TransitionChannel <- true
				}()
			} else {
				fmt.Println("{time:\"" + fmt.Sprint(time.Now().Unix()) + "\",msg:\"日志服务正常启动\"}")
			}
			break
		}
	}
}

//传入debug日志
func (l *Logger) Debug(obj ...pulic_type.LogBase) {
	if l.Level >= DebugLevel {
		for _, v := range obj {
			v.SetLevel("DEBUG")
			l.MsgChannel <- v
		}
	}
}

//传入info日志
func (l *Logger) Info(obj ...pulic_type.LogBase) {
	if l.Level >= InfoLevel {
		for _, v := range obj {
			v.SetLevel("INFO")
			l.MsgChannel <- v
		}
	}
}

//传入Print日志
func (l *Logger) Print(obj ...pulic_type.LogBase) {
	if l.Level >= InfoLevel {
		for _, v := range obj {
			v.SetLevel("PRINT")
			l.MsgChannel <- v
		}
	}
}

//传入Warn日志
func (l *Logger) Warn(obj ...pulic_type.LogBase) {
	if l.Level >= WarnLevel {
		for _, v := range obj {
			v.SetLevel("WARN")
			l.MsgChannel <- v
		}
	}
}

//传入Warn日志
func (l *Logger) Warning(obj ...pulic_type.LogBase) {
	if l.Level >= WarnLevel {
		for _, v := range obj {
			v.SetLevel("WARNING")
			l.MsgChannel <- v
		}
	}
}

//传入Error日志
func (l *Logger) Error(obj ...pulic_type.LogBase) {
	if l.Level >= ErrorLevel {
		for _, v := range obj {
			v.SetLevel("ERROR")
			l.MsgChannel <- v
		}
	}
}

//传入Fatal日志
func (l *Logger) Fatal(obj ...pulic_type.LogBase) {
	if l.Level >= FatalLevel {
		for _, v := range obj {
			v.SetLevel("FATAL")
			l.MsgChannel <- v
		}
	}
}

//传入Fatal日志
func (l *Logger) Panic(obj ...pulic_type.LogBase) {
	if l.Level >= FatalLevel {
		for _, v := range obj {
			v.SetLevel("PANIC")
			l.MsgChannel <- v
		}
	}
}
