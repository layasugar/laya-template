package mdb

import "time"

// User 声明模型
type User struct {
	ID        uint64    `bson:"id"`
	Username  string    `bson:"username"`   // 用户名
	Nickname  string    `bson:"nickname"`   // 昵称
	Avatar    string    `bson:"avatar"`     // 头像
	Mobile    string    `bson:"mobile"`     // 手机号
	Status    uint8     `bson:"status"`     // 状态
	CreatedAt time.Time `bson:"created_at"` // 创建时间
}

func (User) Database() string {
	return "test"
}

func (User) Collection() string {
	return "user"
}
