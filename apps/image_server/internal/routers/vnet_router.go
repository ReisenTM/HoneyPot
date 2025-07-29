package routers

import (
	"Honeypot/apps/image_server/internal/api"
	"Honeypot/apps/image_server/internal/api/vnet_api"
	"Honeypot/apps/image_server/internal/middleware"
	"github.com/gin-gonic/gin"
)

func VNetRouter(r *gin.RouterGroup) {
	app := api.App.VNetApi

	r.PUT("vs_net", middleware.BindJsonMiddleware[vnet_api.VsNetRequest], app.VsNetUpdateView)
	r.GET("vs_net", app.VsNetInfoView)

}
