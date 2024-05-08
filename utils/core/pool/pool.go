package pool

import "errors"

var (
	// ErrClosed 连接池已经关闭Error
	ErrClosed = errors.New("poolx is closed")
	// ErrFactory 错误的工厂函数
	ErrFactory = errors.New("error factory")
)

// Pool 基本方法
type Pool interface {
	Get() (interface{}, error)

	Put(interface{}) error

	Close(interface{}) error

	Release()

	Len() int
}
