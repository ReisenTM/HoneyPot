package user_api

import (
	"github.com/gin-gonic/gin"
	"honeypot_server/internal/enum"
	"honeypot_server/internal/global"
	"honeypot_server/internal/middleware"
	"honeypot_server/internal/models"
	"honeypot_server/internal/utils/resp"
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
