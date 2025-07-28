package log_service

import (
	"Honeypot/apps/honeypot_server/internal/core"
	"Honeypot/apps/honeypot_server/internal/global"
	"Honeypot/apps/honeypot_server/internal/models"
	"github.com/gin-gonic/gin"
)

type LoginLogService struct {
	IP       string
	Location string
}

func NewLoginLog(c *gin.Context) *LoginLogService {
	return &LoginLogService{
		IP:       c.ClientIP(),
		Location: core.GetIpAddr(c.ClientIP()),
	}
}
func (log *LoginLogService) SuccessLog(userID uint, username string) {
	log.save(userID, username, "-", "登录成功", true)
}
func (log *LoginLogService) FailLog(username, password, title string) {
	log.save(0, username, password, title, false)
}

func (log *LoginLogService) save(userID uint, username string, password string, title string, loginStatus bool) {
	global.DB.Create(&models.LogModel{
		Type:        1,
		IP:          log.IP,
		Location:    log.Location,
		UserID:      userID,
		Username:    username,
		Pwd:         password,
		LoginStatus: loginStatus,
		Title:       title,
	})
}
