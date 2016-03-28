package log_client

import (
	"fmt"
	"github.com/seefan/gossdb"
	"sync"
	"time"
)

// 创建SSdb处理对象
func NewSSdbHandle(conf LogConf, log *Logger) (*SSdbHandle, error) {
	ip, port, err := IpPort(conf.Addr)
	if err != nil {
		return nil, err
	}
	// 开启链接池
	pool, err := gossdb.NewPool(&gossdb.Config{
		Host:             ip,
		Port:             port,
		MinPoolSize:      5,
		MaxPoolSize:      50,
		AcquireIncrement: 5,
	})

	if err != nil {
		return nil, err
	}
	flog := SSdbHandle{
		pool: pool,
		Area: conf.Area,
		log:  log,
	}
	return &flog, nil
}

// SSdb处理对象
type SSdbHandle struct {
	pool   *gossdb.Connectors
	Area   string
	log    *Logger
	errNum int64
	num    int64
	// 读写锁
	mu sync.RWMutex
}

// SSdb服务健康检查
func (r *SSdbHandle) CheckHealth() bool {
	c, err := r.pool.NewClient()
	defer c.Close()
	_, err = c.Do("PING")
	if err != nil {
		_base_log.WriteTo(&logErr{
			Level: "ERROR",
			Err:   err.Error(),
			Msg:   "检查SSdb服务器" + r.log.CurrConf.Addr + ",服务无法使用(" + fmt.Sprintf("%d/%d", r.errNum, r.num) + ")",
			Time:  time.Now().Unix(),
		})
		return false
	}
	_base_log.WriteTo(&logErr{
		Level: "INFO",
		Err:   "",
		Msg:   "检查SSdb服务器" + r.log.CurrConf.Addr + ",服务可以使用(" + fmt.Sprintf("%d/%d", r.errNum, r.num) + ")",
		Time:  time.Now().Unix(),
	})
	return true
}

// SSdb处理句柄
func (r *SSdbHandle) WriteTo(msg LogBase) {
	NowTime := time.Now().Unix()
	msg.SetTime(NowTime)

	if r.log.PrintKey {
		fmt.Println(msg)
	}
	reader := jsonFormat(msg)

	// 获取一个连接
	c, err := r.pool.NewClient()

	// 错误处理回调
	back := func() {
		r.log.MsgChannel <- msg
		r.log.Err <- err
	}
	if err != nil {
		r.Error(back)
	}
	defer c.Close()

	_, err = c.Do("LPUSH", r.Area, string(reader))

	r.mu.Lock()
	r.num = r.num + 1
	r.mu.Unlock()

	if err != nil {
		r.Error(back)
	}
	reader = nil
	msg = nil
}

// SSdb处理句柄
func (r *SSdbHandle) RecoveryTo(msg string) {
	c, err := r.pool.NewClient()
	// 错误处理回调
	back := func() {
		r.log.RecoveryChannel <- msg
		r.log.Err <- err
	}
	if err != nil {
		r.Error(back)
	}
	defer c.Close()
	_, err = c.Do("LPUSH", r.Area, msg)

	r.mu.Lock()
	r.num = r.num + 1
	r.mu.Unlock()

	if err != nil {
		r.Error(back)
	}
	msg = ""
}

func (r *SSdbHandle) Error(back func()) {
	r.mu.Lock()
	r.errNum = r.errNum + 1
	r.mu.Unlock()
	go back()
	return
}