package response

const (
	Err                 = 0     // 失败
	Success             = 1     // 成功
	TokenErr            = 40001 // token校验失败
	NoToken             = 40002 // token未找到
	ParamsValidateErr   = 40003 // 参数校验失败
	NotFoundUser        = 40004 // 没有此用户
	FreezeUser          = 40005 // 用户被冻结
	PasswordErr         = 40006 // 密码错误
	RequestFrequentUuid = 40007 // 请求频繁
	RequestFrequentTime = 40008 // 用户时间与服务器时间差距太大
	RequestFrequentSign = 40009 // 接口签名失败
)
