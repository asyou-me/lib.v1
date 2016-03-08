package redis

import (
	"github.com/garyburd/redigo/redis"
	"sync"
	"time"
)

var (
	RedisConnectTimeout = 2000
	RedisReadTimeout    = 1000
	RediswriteTimeout   = 1000
)

// 新建一个新的redis链接池
func NewPool(addr string, max int64) (*Pool, error) {
	pool := Pool{MAX_POOL_SIZE: max}
	pool.Pool = make(chan redis.Conn, max)
	pool.Address = addr
	var i int64
	var err error

	for i = 0; i < max; i++ {
		var conn redis.Conn
		conn, err = redis.DialTimeout("tcp", addr, time.Duration(RedisConnectTimeout*1000000),
			time.Duration(RedisReadTimeout*1000000), time.Duration(RediswriteTimeout*1000000))
		if err != nil {
			return nil, err
		}
		pool.Put(conn)
	}

	return &pool, err
}

// redis链接池
type Pool struct {
	MAX_POOL_SIZE int64
	Address       string
	Pool          chan redis.Conn
	Err           error
	inUse         int64
	// 读写锁
	mu sync.RWMutex
}

// 将链接放回链接池
func (r *Pool) Put(conn redis.Conn) {
	if conn.Err() != nil {
		conn, _ = redis.DialTimeout("tcp", r.Address, time.Duration(RedisConnectTimeout),
			time.Duration(RedisReadTimeout), time.Duration(RediswriteTimeout))
	}
	if int64(len(r.Pool)) >= r.MAX_POOL_SIZE {
		conn.Close()
		return
	}
	r.Pool <- conn
}

// 获取一个redis链接
func (r *Pool) Get() (conn redis.Conn) {
	conn = <-r.Pool
	return conn
}

// 处理redis请求
func (r *Pool) Do(cmd string, args ...interface{}) (re interface{}, err error) {
	conn := r.Get()
	re, err = conn.Do(cmd, args...)

	// 通过携程将用完的链接归还
	go func() {
		r.Put(conn)
		r.mu.Lock()
		r.inUse = r.inUse - 1
		r.mu.Unlock()
	}()

	return
}
