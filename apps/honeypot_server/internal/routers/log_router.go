package routers

import (
	"github.com/gin-gonic/gin"
	"honeypot_server/internal/api"
	"honeypot_server/internal/api/log_api"
	middleware2 "honeypot_server/internal/middleware"
)

func LogRouter(r *gin.RouterGroup) {
	app := api.App.LogApi
	r.GET("logs", middleware2.AdminMiddleware, middleware2.BindQueryMiddleware[log_api.LogListRequest], app.LogListView)
	r.DELETE("logs", middleware2.AdminMiddleware, middleware2.BindJsonMiddleware[log_api.LogRemoveRequest], app.LogRemoveView)
}
