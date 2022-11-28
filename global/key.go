package global

const (
	UserAuth      = "Authorization"
	VerifyCodeKey = ":verify_code:%s:%s"
	TokenRedisKey = ":user:login:%s"  // 用户登录token redis key
	TokenExpire   = 30 * 86400        // token有效期, 单位秒
	UserInfo      = "user_info"       // token验证后在ctx中的字段名
	AdminTokenKey = ":admin:token:%s" // admin token
	AdminUserInfo = "admin_user_info"
)
