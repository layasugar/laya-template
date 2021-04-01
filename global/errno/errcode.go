package errno

import "github.com/layatips/laya/gresp"

var (
	UserNotFound = gresp.Err(40014, "用户不存在")
)
