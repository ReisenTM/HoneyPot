package log_service

import (
	"github.com/gin-gonic/gin"
	"honeypot_server/internal/core"
	"honeypot_server/internal/global"
	"honeypot_server/internal/models"
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
