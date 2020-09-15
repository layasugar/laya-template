package response

const (
	Err     = 0 // 失败
	Success = 1 // 成功

	// 登录验证，接口签名错误
	TokenErr            = 40001 // token校验失败
	ParamsValidateErr   = 40002 // 参数校验失败
	NotFoundUser        = 40003 // 没有此用户
	FreezeUser          = 40004 // 用户被冻结
	PasswordErr         = 40005 // 密码错误
	RequestFrequentUuid = 40006 // 请求频繁
	RequestFrequentTime = 40007 // 用户时间与服务器时间差距太大
	RequestFrequentSign = 40008 // 接口签名失败

	// 文件上传错误
	NoFile           = 50001 // 未接收到文件
	SaveUploadedFail = 50002 // 保存文件失败
	NoFileType       = 50003 // 请选择文件类型
)
