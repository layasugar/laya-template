package validate

// loginData 登录数据验证器
type LoginData struct {
	Name     string `binding:"required"`
	Zone     string `binding:"required"`
	Password string `binding:"required,len=32"`
}

// TokenLoginData token登录参数验证器
type TokenLoginData struct {
	Token string `binding:"required,len=32"`
}

// RegisterData 注册参数验证器
type RegisterData struct {
	Password      string `binding:"required"`
	RPassword     string `binding:"required"`
	Phone         string `binding:"required"`
	Zone          string `binding:"required"`
	InviteCode    string
	PhoneCode     string
	ChannelNumber string
}

// 发送短信参数验证器
type PhoneCodeData struct {
	Phone string `binding:"required"`
	Zone  string `binding:"required"`
	ID    string `binding:"required"`
	Value string `binding:"required"`
	Type  int64  `binding:"required"`
}
