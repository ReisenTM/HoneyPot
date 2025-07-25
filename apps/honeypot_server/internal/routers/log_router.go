package routers

import (
	"Honeypot/apps/honeypot_server/internal/api"
	"Honeypot/apps/honeypot_server/internal/api/log_api"
	"Honeypot/apps/honeypot_server/middleware"
	"github.com/gin-gonic/gin"
)

func LogRouter(r *gin.RouterGroup) {
	app := api.App.LogApi
	r.GET("logs", middleware.AdminMiddleware, middleware.BindQueryMiddleware[log_api.LogListRequest], app.LogListView)
	r.DELETE("logs", middleware.AdminMiddleware, middleware.BindJsonMiddleware[log_api.LogRemoveRequest], app.LogRemoveView)
}
