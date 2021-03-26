package middleware

import "github.com/layatips/laya/middleware"

var Base = &Middleware{}

type Middleware struct {
	middleware.Middleware
}
