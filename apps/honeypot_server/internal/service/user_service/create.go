package user_service

import (
	"fmt"
	"honeypot_server/internal/enum"
	"honeypot_server/internal/global"
	"honeypot_server/internal/models"
	"honeypot_server/internal/utils/pwd"
)

type UserCreateRequest struct {
	Role     enum.Role `json:"role"`
	Username string    `json:"username"`
	Password string    `json:"password"`
}

func (u *UserService) Create(req UserCreateRequest) (user models.UserModel, err error) {
	err = global.DB.Take(&user, "username = ?", req.Username).Error
	if err == nil {
		err = fmt.Errorf("%s 用户名已存在", req.Username)
		return
	}

	hashPwd, _ := pwd.GenerateFromPassword(req.Password)
	user = models.UserModel{
		Username: req.Username,
		Password: hashPwd,
		Role:     req.Role,
	}
	err = global.DB.Create(&user).Error
	if err != nil {
		err = fmt.Errorf("用户创建失败 %s", err)
		return
	}
	u.log.Infof("%s 用户创建成功", req.Username)
	return
}
