package model

// 声明模型User
type User struct {
	Uid      int64
	UserName string
	Phone    string
	Password string `json:"-"`
}

// 将 User 的表名设置为 `user`
func (User) TableName() string {
	return "user"
}
