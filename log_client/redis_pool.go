package log

import (
	"github.com/garyburd/redigo/redis"
	//"sync"
	"time"
)

// 新建一个新的redis链接池
func NewRedisPool(addr string, max int) (*redisPool, error) {
	pool := redisPool{MAX_POOL_SIZE: max}
	pool.Pool = make(chan redis.Conn, max)
	pool.Address = addr
	conn, err := redis.DialTimeout("tcp", addr, time.Duration(RedisConnectTimeout),
		time.Duration(RedisReadTimeout), time.Duration(RediswriteTimeout))
	pool.Put(conn)
	return &pool, err
}

// redis链接池
type redisPool struct {
	MAX_POOL_SIZE int
	Address       string
	Pool          chan redis.Conn
	Err           error
	//mu            sync.Mutex
}

// 将链接放回链接池
func (r *redisPool) Put(conn redis.Conn) {
	if len(r.Pool) >= r.MAX_POOL_SIZE {
		conn.Close()
		return
	}
	r.Pool <- conn
}

// 获取一个redis链接
func (r *redisPool) Get() (conn redis.Conn) {
	// 缓冲机制，在管道中没有连接的时候创建
	if len(r.Pool) == 0 {
		// 新携程创建连接
		go func() {
			for i := 0; i < r.MAX_POOL_SIZE/2; i++ {
				c, err := redis.DialTimeout("tcp", r.Address, time.Duration(RedisConnectTimeout),
					time.Duration(RedisReadTimeout), time.Duration(RediswriteTimeout))
				if err == nil {
					r.Put(c)
				} else {
					r.Err = err
				}
			}
		}()
	}
	conn = <-r.Pool
	return conn
}

// 处理redis请求
func (r *redisPool) Do(cmd string, args ...interface{}) (re interface{}, err error) {
	conn := r.Get()
	re, err = conn.Do(cmd, args...)

	// 通过携程将用完的链接归还
	go func() {
		r.Put(conn)
	}()

	return
}
