package log_client

import (
	"fmt"
	"github.com/asyoume/lib/redis"
	"sync"
	"time"
)

// 创建redis处理对象
func NewRedisHandle(conf LogConf, log *Logger) (*RedisHandle, error) {
	c, err := redis.NewPool(conf.Addr, 20)
	if err != nil {
		return nil, err
	}
	flog := RedisHandle{
		Client: c,
		Area:   conf.Area,
		log:    log,
	}
	return &flog, nil
}

// redis处理对象
type RedisHandle struct {
	Client *redis.Pool
	Area   string
	log    *Logger
	errNum int64
	num    int64
	// 读写锁
	mu sync.RWMutex
}

// redis服务健康检查
func (r *RedisHandle) CheckHealth() bool {
	_, err := r.Client.Do("PING")
	if err != nil {
		_base_log.WriteTo(&log_err{
			Level: "ERROR",
			Err:   err.Error(),
			Msg:   "检查redis服务器" + r.Client.Address + ",服务无法使用(" + fmt.Sprintf("%d/%d", r.errNum, r.num) + ")",
			Time:  time.Now().Unix(),
		})
		return false
	}
	_base_log.WriteTo(&log_err{
		Level: "INFO",
		Err:   "",
		Msg:   "检查redis服务器" + r.Client.Address + ",服务可以使用(" + fmt.Sprintf("%d/%d", r.errNum, r.num) + ")",
		Time:  time.Now().Unix(),
	})
	return true
}

// redis处理句柄
func (r *RedisHandle) WriteTo(msg LogBase) {
	NowTime := time.Now().Unix()
	msg.SetTime(NowTime)

	if r.log.PrintKey {
		fmt.Println(msg)
	}
	reader := jsonFormat(msg)
	_, err := r.Client.Do("LPUSH", r.Area, string(reader))
	r.mu.Lock()
	r.num = r.num + 1
	r.mu.Unlock()
	if err != nil {
		r.mu.Lock()
		r.errNum = r.errNum + 1
		r.mu.Unlock()
		go func() {
			r.log.NewsChannel <- msg
			r.log.Err <- err
		}()
		return
	}
	reader = nil
	msg = nil
}

// redis处理句柄
func (r *RedisHandle) RecoveryTo(msg string) {
	_, err := r.Client.Do("LPUSH", r.Area, msg)
	r.mu.Lock()
	r.num = r.num + 1
	r.mu.Unlock()
	if err != nil {
		r.mu.Lock()
		r.errNum = r.errNum + 1
		r.mu.Unlock()
		go func() {
			r.log.RecoveryChannel <- msg
			r.log.Err <- err
		}()
		return
	}
	msg = ""
}
