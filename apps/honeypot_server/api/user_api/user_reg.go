package user_api

import (
	"Honeypot/apps/honeypot_server/enum"
	"Honeypot/apps/honeypot_server/middleware"
	"Honeypot/apps/honeypot_server/service/user_service"
	"Honeypot/apps/honeypot_server/utils/resp"
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserRegisterRequest struct {
	Username string    `json:"username" binding:"required"`
	Password string    `json:"password" binding:"required" label:"密码"`
	Role     enum.Role `json:"role" binding:"required,ne=1"` //ne = not equal
}

func (UserApi) UserRegisterView(c *gin.Context) {
	cr := middleware.GetBind[UserRegisterRequest](c)

	log := middleware.GetLog(c)
	us := user_service.NewUserService(log)
	user, err := us.Create(user_service.UserCreateRequest{
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
