package cache

import (
	"sync"
	"time"
)

type LocalCacheM struct {
	lock   *sync.RWMutex
	caches map[string]*Cache
}

type Cache struct {
	Value interface{}
	Time  int64
}

func NewLocalCacheM(size int) *LocalCacheM {
	return &LocalCacheM{new(sync.RWMutex), make(map[string]*Cache, size)}
}

func (this *LocalCacheM) Set(key string, v *Cache) {
	this.lock.Lock()
	this.caches[key] = v
	this.lock.Unlock()
}

func (this *LocalCacheM) Get(key string) *Cache {
	this.lock.Lock()
	v := this.caches[key]
	this.lock.Unlock()
	return v
}

func (this *LocalCacheM) Delete(key string) (v *Cache) {
	this.lock.Lock()
	v = this.caches[key]
	delete(this.caches, key)
	this.lock.Unlock()
	return v
}

func (this *LocalCacheM) IsExpired(key string, ttl int) bool {
	if v := this.Get(key); v != nil {
		return (time.Now().Unix() - v.Time) >= int64(ttl)
	} else {
		return true
	}
}
