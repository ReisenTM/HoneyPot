package user_api

import (
	"Honeypot/apps/honeypot_server/enum"
	"Honeypot/apps/honeypot_server/global"
	"Honeypot/apps/honeypot_server/middleware"
	"Honeypot/apps/honeypot_server/models"
	"Honeypot/apps/honeypot_server/utils/resp"
	"github.com/gin-gonic/gin"
)

type UserInfoResponse struct {
	UserID        uint      `json:"user_id"`
	Username      string    `json:"username"`
	Role          enum.Role `json:"role"` // 1 用户 2 管理员
	LastLoginDate string    `json:"last_login_date"`
}

// UserInfoView 用户信息
func (UserApi) UserInfoView(c *gin.Context) {
	auth := middleware.GetAuth(c)

	var user models.UserModel
	err := global.DB.Take(&user, auth.UserID).Error
	if err != nil {
		resp.FailWithMsg("用户不存在", c)
		return
	}

	data := UserInfoResponse{
		UserID:        user.ID,
		Username:      user.Username,
		Role:          user.Role,
		LastLoginDate: user.LastLoginDate,
	}
	resp.OkWithData(data, c)
}
