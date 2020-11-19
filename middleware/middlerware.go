package middleware

import "github.com/LaYa-op/laya/middleware"

var Base = &Middleware{}

type Middleware struct {
	middleware.Middleware
}
