package routers

import (
	"github.com/gin-gonic/gin"
	"honeypot_server/internal/api"
)

func CaptchaRouter(r *gin.RouterGroup) {
	app := api.App.CaptchaApi
	r.GET("captcha", app.GenerateView)
}
