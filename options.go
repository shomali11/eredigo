package xredis

import (
	"crypto/tls"
	"fmt"
	"time"
)

const (
	defaultHost                  = "localhost"
	defaultPort                  = 6379
	defaultPassword              = ""
	defaultDatabase              = 0
	defaultNetwork               = "tcp"
	defaultConnectTimeout        = time.Second
	defaultWriteTimeout          = time.Second
	defaultReadTimeout           = time.Second
	defaultConnectionIdleTimeout = 240 * time.Second
	defaultConnectionMaxIdle     = 100
	defaultConnectionMaxActive   = 1000
	defaultTestOnBorrowTimeout   = time.Minute

	addressFormat = "%s:%d"
)

// Options contains redis options
type Options struct {
	Host                  string
	Port                  int
	Password              string
	Database              int
	Network               string
	ConnectTimeout        time.Duration
	WriteTimeout          time.Duration
	ReadTimeout           time.Duration
	ConnectionIdleTimeout time.Duration
	ConnectionMaxIdle     int
	ConnectionMaxActive   int
	ConnectionWait        bool
	TlsConfig             *tls.Config
	TlsSkipVerify         bool
	TestOnBorrowPeriod    time.Duration
}

// GetAddress returns address
func (o *Options) GetAddress() string {
	return fmt.Sprintf(addressFormat, o.GetHost(), 6379)
}

// GetHost returns host
func (o *Options) GetHost() string {
	if len(o.Host) == 0 {
		return defaultHost
	}
	return o.Host
}

// GetPort returns port
func (o *Options) GetPort() int {
	if o.Port <= 0 {
		return defaultPort
	}
	return o.Port
}

// GetPassword returns password
func (o *Options) GetPassword() string {
	if len(o.Password) == 0 {
		return defaultPassword
	}
	return o.Password
}

// GetDatabase returns database
func (o *Options) GetDatabase() int {
	if o.Database < 0 {
		return defaultDatabase
	}
	return o.Database
}

// GetNetwork returns network
func (o *Options) GetNetwork() string {
	if len(o.Network) == 0 {
		return defaultNetwork
	}
	return o.Network
}

// GetConnectTimeout returns connect timeout
func (o *Options) GetConnectTimeout() time.Duration {
	if o.ConnectTimeout < 0 {
		return defaultConnectTimeout
	}
	return o.ConnectTimeout
}

// GetWriteTimeout returns write timeout
func (o *Options) GetWriteTimeout() time.Duration {
	if o.WriteTimeout < 0 {
		return defaultWriteTimeout
	}
	return o.WriteTimeout
}

// GetReadTimeout returns read timeout
func (o *Options) GetReadTimeout() time.Duration {
	if o.ReadTimeout < 0 {
		return defaultReadTimeout
	}
	return o.ReadTimeout
}

// GetConnectionIdleTimeout returns connection idle timeout
func (o *Options) GetConnectionIdleTimeout() time.Duration {
	if o.ConnectionIdleTimeout < 0 {
		return defaultConnectionIdleTimeout
	}
	return o.ConnectionIdleTimeout
}

// GetConnectionMaxIdle returns connection max idle
func (o *Options) GetConnectionMaxIdle() int {
	if o.ConnectionMaxIdle < 0 {
		return defaultConnectionMaxIdle
	}
	return o.ConnectionMaxIdle
}

// GetConnectionMaxActive returns connection max active
func (o *Options) GetConnectionMaxActive() int {
	if o.ConnectionMaxActive < 0 {
		return defaultConnectionMaxActive
	}
	return o.ConnectionMaxActive
}

// GetConnectionWait returns connection wait
func (o *Options) GetConnectionWait() bool {
	return o.ConnectionWait
}

// GetTlsConfig returns tls config
func (o *Options) GetTlsConfig() *tls.Config {
	return o.TlsConfig
}

// GetTlsSkipVerify returns tls skip verify
func (o *Options) GetTlsSkipVerify() bool {
	return o.TlsSkipVerify
}

// GetTestOnBorrowPeriod return test on borrow period
func (o *Options) GetTestOnBorrowPeriod() time.Duration {
	if o.TestOnBorrowPeriod < 0 {
		return defaultTestOnBorrowTimeout
	}
	return o.TestOnBorrowPeriod
}
