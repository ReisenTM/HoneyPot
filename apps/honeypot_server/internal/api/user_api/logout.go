package user_api

import (
	middleware2 "Honeypot/apps/honeypot_server/internal/middleware"
	"Honeypot/apps/honeypot_server/internal/utils/resp"
	"github.com/gin-gonic/gin"
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
