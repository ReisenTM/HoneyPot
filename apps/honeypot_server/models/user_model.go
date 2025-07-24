package models

import "Honeypot/apps/honeypot_server/enum"

//用户

type UserModel struct {
	Model
	Username      string    `gorm:"size:32" json:"username"`
	Role          enum.Role `json:"role"` // 1 管理员 2 用户
	Password      string    `gorm:"size:64" json:"-"`
	LastLoginDate string    `gorm:"size:32" json:"last_login_date"`
}
