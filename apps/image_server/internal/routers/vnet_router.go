package routers

import (
	"github.com/gin-gonic/gin"
	"image_server/internal/api"
	"image_server/internal/api/vnet_api"
	"image_server/internal/middleware"
)

func VNetRouter(r *gin.RouterGroup) {
	app := api.App.VNetApi

	r.PUT("vs_net", middleware.BindJsonMiddleware[vnet_api.VsNetRequest], app.VsNetUpdateView)
	r.GET("vs_net", app.VsNetInfoView)

}
