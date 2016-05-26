package http

import (
	"net"
	"net/http"
	"time"
)

var (
	DefaultClient *http.Client = NewTimeoutClient(10*time.Second, 20*time.Second)
)

type TimeoutConn struct {
	net.Conn
	timeout time.Duration
}

func NewTimeoutConn(conn net.Conn, timeout time.Duration) *TimeoutConn {
	return &TimeoutConn{conn, timeout}
}

func (c *TimeoutConn) Read(b []byte) (n int, err error) {
	c.SetReadDeadline(time.Now().Add(c.timeout))
	return c.Conn.Read(b)
}

func (c *TimeoutConn) Write(b []byte) (n int, err error) {
	c.SetWriteDeadline(time.Now().Add(c.timeout))
	return c.Conn.Write(b)
}

func TimeoutDialer(cTimeout time.Duration, rwTimeout time.Duration) func(net, addr string) (c net.Conn, err error) {
	return func(netw, addr string) (net.Conn, error) {
		if DnsCacheDuration > 0 {
			addr = dnsCache.Get(addr)
		}

		conn, err := net.DialTimeout(netw, addr, cTimeout)
		if err != nil {
			return nil, err
		}
		return NewTimeoutConn(conn, rwTimeout), nil
	}
}

func NewTimeoutClient(connectTimeout time.Duration, readWriteTimeout time.Duration) *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			Dial:  TimeoutDialer(connectTimeout, readWriteTimeout),
			Proxy: http.ProxyFromEnvironment,
		},
	}
}
