package routers

import (
	"Honeypot/apps/honeypot_server/api"
	"github.com/gin-gonic/gin"
)

func CaptchaRouter(r *gin.RouterGroup) {
	app := api.App.CaptchaApi
	r.GET("captcha", app.GenerateView)
}
