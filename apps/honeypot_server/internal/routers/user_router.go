package routers

import (
	"Honeypot/apps/honeypot_server/internal/api"
	user_api2 "Honeypot/apps/honeypot_server/internal/api/user_api"
	"Honeypot/apps/honeypot_server/middleware"
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup) {
	app := api.App.UserApi
	//登录
	r.POST("login", middleware.BindJsonMiddleware[user_api2.UserLoginRequest], app.UserLoginView)
	//注册
	r.POST("register", middleware.BindJsonMiddleware[user_api2.UserRegisterRequest], app.UserRegisterView)
	//列表
	r.GET("users", middleware.BindQueryMiddleware[user_api2.UserListRequest], app.UserListView)
	//用户注销
	r.POST("logout", app.UserLogoutView)
	//用户批量删除
	r.DELETE("users", middleware.BindJsonMiddleware[user_api2.UserRemoveRequest], app.UserRemoveView)
	//用户信息
	r.GET("info", app.UserInfoView)
}
