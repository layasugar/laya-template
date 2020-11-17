package db

// 声明模型User
type User struct {
	ID       int64
	UserName string
	Zone     string // 手机区号
	Phone    string
	Email    string
	Password string `json:"-"`
	Status   int64  // 用户状态，-1冻结,1正常
}

// 将 User 的表名设置为 `user`
func (User) TableName() string {
	return "laya_user"
}
