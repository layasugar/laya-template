package gcal

import (
	"fmt"
	"time"

	"github.com/layasugar/laya/gcal/context"
	"github.com/layasugar/laya/gcal/service"
)

type client struct {
	serv   service.Service
	isCopy bool
	err    error
}

// Client new client by serviceName
func Client(serviceName string) *client {
	serv, _ := service.GetService(serviceName)
	if serv == nil {
		return &client{
			err: fmt.Errorf("fail get service: %s", serviceName),
		}
	}

	return &client{
		serv: serv,
	}
}

// Do 执行
func (c *client) Do(request interface{}, response interface{}, converterType ConverterType) error {
	if c.err != nil {
		return c.err
	}
	ctx := context.NewContext()
	ctx.Caller = "GCAL"

	c.err = calWithService(ctx, c.serv, request, response, converterType)
	return c.err
}

func (c *client) setPrepare() {
	if !c.isCopy {
		c.serv = c.serv.Clone()
		c.isCopy = true
	}
}

// GetProtocol 得到 protocol
func (c *client) GetProtocol() string {
	if c.err != nil {
		return ""
	}
	return c.serv.GetConf().Protocol
}

// SetProtocol 设置 Protocol
func (c *client) SetProtocol(p string) *client {
	if c.err != nil {
		return c
	}
	c.setPrepare()
	c.serv.GetConf().Protocol = p
	return c
}

// GetRetry 得到重试次数
func (c *client) GetRetry() int {
	if c.err != nil {
		return 0
	}
	return c.serv.GetConf().GetRetry()
}

// SetRetry 设置 retry
func (c *client) SetRetry(retry int) *client {
	if c.err != nil {
		return c
	}
	c.setPrepare()
	c.serv.GetConf().Retry = retry
	return c
}

// GetReuse 得到是否连接
func (c *client) GetReuse() bool {
	if c.err != nil {
		return false
	}
	return c.serv.GetConf().Reuse
}

// SetReuse 设置 是否复用连接
func (c *client) SetReuse(doReuse bool) *client {
	if c.err != nil {
		return c
	}
	c.setPrepare()
	c.serv.GetConf().Reuse = doReuse
	return c
}

// GetConnTimeOut 得到连接超时
func (c *client) GetConnTimeOut() time.Duration {
	if c.err != nil {
		return 0
	}
	return c.serv.GetConf().ConnTimeOut
}

// SetConnTimeOut 设置连接超时
func (c *client) SetConnTimeOut(ct time.Duration) *client {
	if c.err != nil {
		return c
	}
	c.setPrepare()
	c.serv.GetConf().ConnTimeOut = ct

	return c
}

// GetReadTimeOut 得到读超时
func (c *client) GetReadTimeOut() time.Duration {
	if c.err != nil {
		return 0
	}
	return c.serv.GetConf().ReadTimeOut
}

// SetReadTimeOut 设置读超时
func (c *client) SetReadTimeOut(ct time.Duration) *client {
	if c.err != nil {
		return c
	}
	c.setPrepare()
	c.serv.GetConf().ReadTimeOut = ct

	return c
}

// GetWriteTimeOut 得到写超时
func (c *client) GetWriteTimeOut() time.Duration {
	if c.err != nil {
		return 0
	}
	return c.serv.GetConf().WriteTimeOut
}

// SetWriteTimeOut 设置写超时
func (c *client) SetWriteTimeOut(ct time.Duration) *client {
	if c.err != nil {
		return c
	}
	c.setPrepare()
	c.serv.GetConf().WriteTimeOut = ct

	return c
}
