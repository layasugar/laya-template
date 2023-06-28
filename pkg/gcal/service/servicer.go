package service

import (
	"strings"
	"time"
)

// Config cal 配置
type Config struct {
	Name       string
	Addr       string
	Retry      int
	NSProvider string
	Protocol   string
	Converter  string
	Reuse      bool // 连接复用

	ConnTimeOut  time.Duration
	WriteTimeOut time.Duration
	ReadTimeOut  time.Duration

	Headers map[string]string
}

// NewDefaultConfig 创建默认的Config
func NewDefaultConfig() *Config {
	return &Config{
		ConnTimeOut:  1500,
		WriteTimeOut: 1500,
		ReadTimeOut:  1500,
		Converter:    "form",
		Protocol:     "http",
	}
}

// Format 格式化 cal信息
func (rc *Config) Format() {
	rc.Name = strings.ToLower(rc.Name)
	rc.ConnTimeOut = rc.ConnTimeOut * time.Millisecond
	rc.WriteTimeOut = rc.WriteTimeOut * time.Millisecond
	rc.ReadTimeOut = rc.ReadTimeOut * time.Millisecond
	rc.Protocol = strings.ToLower(rc.Protocol)
	rc.Converter = strings.ToLower(rc.Converter)
}

// Clone 拷贝
func (rc *Config) Clone() *Config {
	if rc == nil {
		return &Config{}
	}

	data := &Config{
		Name:         rc.Name,
		Retry:        rc.Retry,
		NSProvider:   rc.NSProvider,
		Protocol:     rc.Protocol,
		Converter:    rc.Converter,
		Headers:      map[string]string{},
		ConnTimeOut:  rc.ConnTimeOut * time.Millisecond,
		WriteTimeOut: rc.WriteTimeOut * time.Millisecond,
		ReadTimeOut:  rc.ReadTimeOut * time.Millisecond,
	}

	for k, v := range rc.Headers {
		data.Headers[k] = v
	}

	return data
}

// GetName 服务发现名称
func (rc *Config) GetName() string {
	return rc.Name
}

// GetAddr 服务发现名称
func (rc *Config) GetAddr() string {
	return rc.Addr
}

// GetTotalTimeout 总超时
func (rc *Config) GetTotalTimeout() time.Duration {
	return rc.ConnTimeOut + rc.ReadTimeOut + rc.WriteTimeOut
}

// GetConnTimeout 连接超时
func (rc *Config) GetConnTimeout() time.Duration {
	return rc.ConnTimeOut
}

// GetWriteTimeout 写超时
func (rc *Config) GetWriteTimeout() time.Duration {
	return rc.WriteTimeOut
}

// GetReadTimeout 读超时
func (rc *Config) GetReadTimeout() time.Duration {
	return rc.ReadTimeOut
}

// GetProtocol 交互协议
func (rc *Config) GetProtocol() string {
	return rc.Protocol
}

// GetRetry 重试次数
func (rc *Config) GetRetry() int {
	return rc.Retry
}

// GetReuse 是否复用连接
func (rc *Config) GetReuse() bool {
	return rc.Reuse
}

// GetConf 得到Conf 引用
func (rc *Config) GetConf() *Config {
	return rc
}

// SetTimeOut 重写timeOut
func (rc *Config) SetTimeOut(t int64) {
	rc.ConnTimeOut = time.Duration(t) * time.Millisecond
}

type Service interface {
	GetConf() *Config
	Clone() *Config

	GetName() string
	GetAddr() string
	GetTotalTimeout() time.Duration
	GetConnTimeout() time.Duration
	GetReadTimeout() time.Duration
	GetWriteTimeout() time.Duration
	GetProtocol() string
	GetRetry() int
	GetReuse() bool

	SetTimeOut(t int64)
}
