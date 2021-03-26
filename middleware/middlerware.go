package middleware

import (
	"github.com/layatips/laya/gmiddleware"
)

var Base = &Middleware{}

type Middleware struct {
	gmiddleware.Middleware
}
