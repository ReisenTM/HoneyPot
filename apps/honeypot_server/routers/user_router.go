package routers

import (
	"Honeypot/apps/honeypot_server/api"
	"Honeypot/apps/honeypot_server/api/user_api"
	"Honeypot/apps/honeypot_server/middleware"
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup) {
	app := api.App.UserApi
	//登录
	r.POST("login", middleware.BindJsonMiddleware[user_api.UserLoginRequest], app.UserLoginView)
	//注册
	r.POST("register", middleware.BindJsonMiddleware[user_api.UserRegisterRequest], app.UserRegisterView)
}
