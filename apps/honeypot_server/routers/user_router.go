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
	//列表
	r.GET("users", middleware.BindQueryMiddleware[user_api.UserListRequest], app.UserListView)
	//用户注销
	r.POST("logout", app.UserLogoutView)
	//用户批量删除
	r.DELETE("users", middleware.BindJsonMiddleware[user_api.UserRemoveRequest], app.UserRemoveView)
}
