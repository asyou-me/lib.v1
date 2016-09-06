package log_client

import (
	"fmt"

	"sync"
	"time"

	"github.com/asyou-me/lib.v1/pulic_type"
	"github.com/garyburd/redigo/redis"
)

// 创建redis处理对象
func NewRedisHandle(conf pulic_type.LogConf, log *Logger) (*RedisHandle, error) {
	pool := &redis.Pool{
		MaxIdle: 20,
		Dial: func() (redis.Conn, error) {
			return redis.DialTimeout("tcp", conf.Addr, time.Second, time.Second, time.Second)
		},
	}
	flog := RedisHandle{
		Client: pool,
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
	client := r.Client.Get()
	defer func() {
		client.Close()
	}()
	_, err := client.Do("PING")
	if err != nil {
		_baseLog.WriteTo(&Loggerr{
			Level: "ERROR",
			Err:   err.Error(),
			Msg:   "检查redis服务器无法使用(" + fmt.Sprintf("%d/%d", r.errNum, r.num) + ")",
			Time:  time.Now().Unix(),
		})
		return false
	}
	_baseLog.WriteTo(&Loggerr{
		Level: "INFO",
		Err:   "",
		Msg:   "检查redis服务器务可以使用(" + fmt.Sprintf("%d/%d", r.errNum, r.num) + ")",
		Time:  time.Now().Unix(),
	})
	return true
}

// redis处理句柄
func (r *RedisHandle) WriteTo(msg pulic_type.LogBase) {
	NowTime := time.Now().Unix()
	msg.SetTime(NowTime)

	reader := jsonFormat(msg)
	client := r.Client.Get()
	defer func() {
		client.Close()
	}()
	_, err := client.Do("LPUSH", r.Area, string(reader))
	r.mu.Lock()
	r.num = r.num + 1
	r.mu.Unlock()
	if err != nil {
		r.mu.Lock()
		r.errNum = r.errNum + 1
		r.mu.Unlock()
		go func() {
			r.log.MsgChannel <- msg
			r.log.Err <- err
		}()
		return
	}
	reader = nil
	msg = nil
}

// redis处理句柄
func (r *RedisHandle) RecoveryTo(msg string) {
	client := r.Client.Get()
	defer func() {
		client.Close()
	}()
	_, err := client.Do("LPUSH", r.Area, msg)
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
