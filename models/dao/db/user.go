package db

import (
	"github.com/layasugar/laya/tools/timex"
	"gorm.io/gorm"
)

// User 声明模型
type User struct {
	ID        uint64         `gorm:"column:id" json:"id"`
	Username  string         `gorm:"column:username" json:"username"` // 用户名
	Nickname  string         `gorm:"column:nickname" json:"nickname"` // 昵称
	Avatar    string         `gorm:"column:avatar" json:"avatar"`     // 头像
	Password  string         `gorm:"column:password" json:"-"`        // 密码
	Salt      string         `gorm:"column:salt" json:"-"`            // 盐
	Mobile    string         `gorm:"column:mobile" json:"mobile"`     // 手机号
	Status    uint8          `gorm:"column:status" json:"status"`     // 状态
	CreatedAt timex.Time     `gorm:"column:created_at" json:"created_at"`
	UpdatedAt timex.Time     `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName 将 User 的表名设置为 user
func (User) TableName() string {
	return "user"
}
