// Package service 提供了一些列资源管理的方法
package service

import (
	"errors"
	"strings"
	"sync"
	"time"
)

var servicePool = map[string]Service{}
var serviceLock sync.RWMutex

// GetService 得到服务
func GetService(name string) (Service, bool) {
	serviceLock.RLock()
	s, ok := servicePool[name]
	if !ok {
		return defaultService(), true
	}
	serviceLock.RUnlock()
	return s, ok
}

// CleanService 清除服务
func CleanService() {
	serviceLock.Lock()
	servicePool = map[string]Service{}
	serviceLock.Unlock()
}

// RemoveService 清除指定的Service
func RemoveService(name string) {
	s, _ := GetService(name)
	if s == nil {
		return
	}
	serviceLock.Lock()
	delete(servicePool, name)
	serviceLock.Unlock()
}

func NewService(cfg *Config) Service {
	cfg.Format()
	return cfg
}

// LoadService 添加服务
func LoadService(cfg []map[string]interface{}) error {
	for _, item := range cfg {
		var addr string
		var protocol string
		as := strings.Split(item["addr"].(string), "://")
		if len(as) == 2 {
			protocol = as[0]
			addr = as[1]
		} else {
			panic(errors.New("protocol is nil"))
		}

		var retry int64 = 0
		if _, ok := item["retry"]; ok {
			if _, ok = item["retry"].(int64); ok {
				retry = item["retry"].(int64)
			}
		}

		var connTimeOut time.Duration
		if _, ok := item["conn_time_out"]; ok {
			connTimeOut = time.Duration(item["conn_time_out"].(int64))
		}

		var writeTimeOut time.Duration
		if _, ok := item["write_time_out"]; ok {
			writeTimeOut = time.Duration(item["write_time_out"].(int64))
		}

		var readTimeOut time.Duration
		if _, ok := item["read_time_out"]; ok {
			readTimeOut = time.Duration(item["read_time_out"].(int64))
		}

		var conf = Config{
			Name:         item["name"].(string),
			Addr:         addr,
			Retry:        int(retry),
			NSProvider:   "1",
			Protocol:     protocol,
			Reuse:        true,
			Converter:    "json",
			ConnTimeOut:  connTimeOut,
			WriteTimeOut: writeTimeOut,
			ReadTimeOut:  readTimeOut,
			Headers:      make(map[string]string),
		}
		service := NewService(&conf)
		serviceLock.Lock()
		servicePool[service.GetName()] = service
		serviceLock.Unlock()
	}

	return nil
}

func defaultService() Service {
	var conf = Config{
		Name:         "default",
		Addr:         "",
		Retry:        0,
		NSProvider:   "1",
		Protocol:     "",
		Reuse:        true,
		Converter:    "json",
		ConnTimeOut:  1500,
		WriteTimeOut: 1500,
		ReadTimeOut:  1500,
		Headers:      make(map[string]string),
	}

	return NewService(&conf)
}
