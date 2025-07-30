package user_api

import (
	"github.com/gin-gonic/gin"
	middleware2 "honeypot_server/internal/middleware"
	"honeypot_server/internal/utils/resp"
	"time"
)

func (UserApi) UserLogoutView(c *gin.Context) {
	token := c.GetHeader("token")
	log := middleware2.GetLog(c)
	auth := middleware2.GetAuth(c)
	expiresAt := time.Unix(auth.ExpiresAt, 0)

	log.Infof("用户注销 %d %s %s", auth.UserID, token, expiresAt)
	resp.OkWithMsg("注销成功", c)
}
