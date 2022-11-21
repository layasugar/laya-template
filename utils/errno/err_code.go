package errno

import "github.com/layasugar/laya-template/global"

var (
	ComUnauthorized = global.Err(401)
	UserNotFound    = global.Err(4001)
)
