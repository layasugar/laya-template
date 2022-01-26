package db

type AdminUser struct {
	ID          uint64 `json:"id" ddb:"id"`
	DefaultRole uint64 `json:"default_role" ddb:"default_role"`
	Username    string `json:"username" ddb:"username"`
	Password    string `json:"-" ddb:"password"`
	Salt        string `json:"-" ddb:"salt"`
	Nickname    string `json:"nickname" ddb:"nickname"`
	Avatar      string `json:"avatar" ddb:"avatar"`
	Desc        string `json:"desc" ddb:"desc"`
	Status      uint8  `json:"status" ddb:"status"`
	LastLoginIp string `json:"last_login_ip" ddb:"last_login_ip"`
	LastLoginAt string `json:"last_login_at" ddb:"last_login_at"`
	CreatedAt   string `json:"created_at" ddb:"created_at"`
	UpdatedAt   string `json:"updated_at" ddb:"updated_at"`
	DeletedAt   string `json:"-" ddb:"deleted_at"`
}

// TableName 将 User 的表名设置为 `user`
func (AdminUser) TableName() string {
	return "admin_user"
}
