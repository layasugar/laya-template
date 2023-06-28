package rdb

import (
	"encoding/json"
	"log"
)

// User 声明模型
type User struct {
	ID        uint64 `json:"id"`
	Username  string `json:"username"`   // 用户名
	Nickname  string `json:"nickname"`   // 昵称
	Avatar    string `json:"avatar"`     // 头像
	Mobile    string `json:"mobile"`     // 手机号
	Status    uint8  `json:"status"`     // 状态
	CreatedAt uint64 `json:"created_at"` // 创建时间
}

func (u User) String() string {
	res, err := json.Marshal(&u)
	if err != nil {
		log.Print("cjson marshal err")
		return ""
	}
	return string(res)
}
