package db

import (
	"time"
)

// User 声明模型
type User struct {
	ID            int64     `json:"id"`
	Status        int64     `json:"status"` // 用户状态，2冻结,1正常
	UserName      string    `json:"user_name"`
	Zone          string    `json:"zone"` // 手机区号
	Phone         string    `json:"phone"`
	Email         string    `json:"email"`
	Password      string    `json:"-"`
	InviteCode    string    `json:"invite_code"`
	CreatedAt     time.Time `json:"created_at"`
	LastLoginIp   string    `json:"last_login_ip"`
	LastLoginTime time.Time `json:"last_login_time"`
}

// TableName 将 User 的表名设置为 `user`
func (User) TableName() string {
	return "user"
}
