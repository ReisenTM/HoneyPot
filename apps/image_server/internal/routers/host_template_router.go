package routers

import (
	"Honeypot/apps/image_server/internal/api"
	"Honeypot/apps/image_server/internal/api/host_template_api"
	"Honeypot/apps/image_server/internal/middleware"
	"Honeypot/apps/image_server/internal/models"
	"github.com/gin-gonic/gin"
)

func HostTemplateRouter(r *gin.RouterGroup) {
	app := api.App.HostTemplateApi

	r.POST("host_template", middleware.BindJsonMiddleware[host_template_api.CreateRequest], app.CreateView)
	r.PUT("host_template", middleware.BindJsonMiddleware[host_template_api.UpdateRequest], app.UpdateView)
	r.GET("host_template", middleware.BindQueryMiddleware[models.PageInfo], app.ListView)
	r.GET("host_template/options", app.OptionsView)
	r.DELETE("host_template", middleware.BindJsonMiddleware[models.IDListRequest], app.Remove)

}
