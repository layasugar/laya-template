package errno

import "github.com/layatips/laya/gresp"

var Errno = map[uint32]string{
	400: "系统错误",
}

var (
	SystemErr = gresp.Err(400)
)
