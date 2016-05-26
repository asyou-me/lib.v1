package http

import (
	"fmt"
	"net"
	"sync"
	"time"
)

var (
	DnsCacheDuration time.Duration = 5 * time.Second
	dnsCache                       = &DnsCache{caches: make(map[string]DnsCacheItem)}
)

type DnsCacheItem struct {
	IP        string
	CacheTime int64
}

type DnsCache struct {
	sync.RWMutex
	caches map[string]DnsCacheItem
}

func (this *DnsCache) Get(addr string) string {
	if DnsCacheDuration <= 0 {
		return addr
	}
	host, port, err := net.SplitHostPort(addr)
	if err != nil {
		return addr
	}
	this.RLock()
	item, ok := this.caches[host]
	this.RUnlock()

	if !ok || time.Now().Unix()-item.CacheTime > int64(DnsCacheDuration/time.Second) {
		go func() {
			netAddr, err := net.ResolveTCPAddr("tcp", addr)
			if err == nil {
				this.Lock()
				this.caches[host] = DnsCacheItem{IP: netAddr.IP.String(), CacheTime: time.Now().Unix()}
				this.Unlock()
			}
		}()
	}
	if ok {
		return fmt.Sprintf("%s:%s", item.IP, port)
	} else {
		return addr
	}
}
