package log

import (
	"errors"
	"fmt"
	"time"
)

// 日志对象
type Logger struct {
	// 配置文件
	Conf     []LogConf
	CurrConf LogConf
	// 运行状态
	Run int64
	// 后端对象
	BackEnd      LogInterface
	BackEndTrash LogInterface
	// 重启动相关
	ChoiceChannel chan bool
	lastChoice    int64

	//日志记录等级
	Level Level

	//日志打印等级
	PrintKey bool

	// 消息管道
	NewsChannel     chan LogBase
	RecoveryChannel chan string

	Err         chan error
	ErrNum      int64
	StopChannel chan bool
}

// 日志对象初始化流程
func (l *Logger) Init(conf *[]LogConf) error {

	// 据权重调整配置文件顺序
	LogConfSort(*conf)

	l.Conf = *conf
	l.BackEndTrash = nil

	// 日志列表
	l.NewsChannel = make(chan LogBase, 1000)
	l.StopChannel = make(chan bool)
	l.Err = make(chan error, 100)
	l.ChoiceChannel = make(chan bool, 1)

	l.Level = DebugLevel
	// 开启日志消费模式
	l.Choice(false)
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
	var news = l.NewsChannel
	var err = l.Err
	var stop = l.StopChannel
	var choise = l.ChoiceChannel
	for {
		select {
		// log消息列表
		case log := <-news:
			l.BackEnd.Do(log)

		// 错误消息列表
		case e := <-err:
			file_local_log.Do(&log_err{
				Level: "WARN",
				Msg:   e.Error(),
				Time:  time.Now().Unix(),
			})

			l.ErrNum = l.ErrNum + 1
			var now_unix_time = time.Now().Unix()
			if l.ErrNum > 100 && now_unix_time-l.lastChoice > 5 {
				l.lastChoice = now_unix_time
				go func() {
					if len(choise) == 0 {
						choise <- true
					}
				}()
			}

		// 服务变更通道
		case _ = <-choise:
			if len(choise) == 0 {
				ok := l.BackEnd.Check()
				if ok {
					l.Run = 1
					l.ErrNum = 0
				} else {
					l.Run = 0
				}
				// 低优先和当前服务不可用时
				if !ok || l.CurrConf.Spare == true {
					// 检查高优先服务的可用性
					l.Choice(false)
					if l.Run == 0 && !ok {
						//暂时低优先服务
						l.Choice(true)
					}
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
func (l *Logger) Choice(spare bool) {
	var err error
	for i := 0; i < len(l.Conf); i++ {

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
				file_local_log.Do(&log_err{
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
				fmt.Println("\033[31;1mPanic：文件日志无法使用")
			}
			break
		default:
			continue
		}

		if l.Run == 1 {
			if l.CurrConf.Spare == true {
				fmt.Println("\033[31;1m日志进入低优先模式,系统会以", ReconnectInterval,
					"秒为周期检查高权重服务的可用性")
				file_local_log.Do(&log_err{
					Level: "INFO",
					Msg:   "高权重服务不可用,日志进入备用模式",
					Time:  time.Now().Unix(),
				})
				go func() {
					time.Sleep(time.Duration(ReconnectInterval) * time.Second)
					l.ChoiceChannel <- true
				}()
			} else {
				go func() {
					time.Sleep(60 * time.Second)
					l.ChoiceChannel <- true
				}()
				fmt.Println("\033[32;1m日志服务正常启动")
			}
			break
		}
	}
}

//传入debug日志
func (l *Logger) Debug(obj ...LogBase) {
	if l.Level >= DebugLevel {
		for _, v := range obj {
			v.SetLevel("DEBUG")
			go func() {
				l.NewsChannel <- v
			}()
		}
	}
}

//传入info日志
func (l *Logger) Info(obj ...LogBase) {
	if l.Level >= InfoLevel {
		for _, v := range obj {
			v.SetLevel("INFO")
			go func() {
				l.BackEnd.Do(v)
			}()
		}
	}
}

//传入Print日志
func (l *Logger) Print(obj ...LogBase) {
	if l.Level >= InfoLevel {
		for _, v := range obj {
			v.SetLevel("PRINT")
			go func() {
				l.BackEnd.Do(v)
			}()
		}
	}
}

//传入Warn日志
func (l *Logger) Warn(obj ...LogBase) {
	if l.Level >= WarnLevel {
		for _, v := range obj {
			v.SetLevel("WARN")
			go func() {
				l.BackEnd.Do(v)
			}()
		}
	}
}

//传入Error日志
func (l *Logger) Error(obj ...LogBase) {
	if l.Level >= ErrorLevel {
		for _, v := range obj {
			v.SetLevel("ERROR")
			go func() {
				l.BackEnd.Do(v)
			}()
		}
	}
}

//传入Fatal日志
func (l *Logger) Fatal(obj ...LogBase) {
	if l.Level >= FatalLevel {
		for _, v := range obj {
			v.SetLevel("FATAL")
			go func() {
				l.BackEnd.Do(v)
			}()
		}
	}
}

/*func (l *Logger) Get() {

}

func (l *Logger) List() {

}

func (l *Logger) Put(i int64) {

}*/
