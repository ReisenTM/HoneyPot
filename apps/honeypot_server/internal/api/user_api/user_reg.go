package user_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"honeypot_server/internal/enum"
	middleware2 "honeypot_server/internal/middleware"
	user_service2 "honeypot_server/internal/service/user_service"
	"honeypot_server/internal/utils/resp"
)

type UserRegisterRequest struct {
	Username string    `json:"username" binding:"required"`
	Password string    `json:"password" binding:"required" label:"密码"`
	Role     enum.Role `json:"role" binding:"required,ne=2"` //ne = not equal
}

func (UserApi) UserRegisterView(c *gin.Context) {
	cr := middleware2.GetBind[UserRegisterRequest](c)

	log := middleware2.GetLog(c)
	us := user_service2.NewUserService(log)
	user, err := us.Create(user_service2.UserCreateRequest{
		Username: cr.Username,
		Password: cr.Password,
		Role:     cr.Role,
	})
	if err != nil {
		msg := fmt.Sprintf("创建用户失败 %s", err)
		log.Errorf(msg)
		resp.FailWithMsg(msg, c)
		return
	}
	resp.OkWithData(user.ID, c)
}
