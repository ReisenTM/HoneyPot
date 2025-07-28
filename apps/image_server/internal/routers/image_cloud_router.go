package routers

import (
	"Honeypot/apps/image_server/internal/api"
	"Honeypot/apps/image_server/internal/api/image_cloud_api"
	"Honeypot/apps/image_server/internal/middleware"
	"Honeypot/apps/image_server/internal/models"
	"github.com/gin-gonic/gin"
)

func ImageCloudRouter(r *gin.RouterGroup) {
	app := api.App.ImageCloudApi
	r.POST("image_cloud/upload", app.ImageUploadView)
	r.POST("image_cloud", middleware.BindJsonMiddleware[image_cloud_api.ImageCreateRequest], app.ImageCreateView)
	r.GET("image_cloud", middleware.BindQueryMiddleware[image_cloud_api.ImageListRequest], app.ImageListView)
	r.GET("image_cloud/:id", middleware.BindUriMiddleware[models.IDRequest], app.ImageDetailView)
	r.PUT("image_cloud", middleware.BindJsonMiddleware[image_cloud_api.ImageUpdateRequest], app.ImageUpdateView)
	r.DELETE("image_cloud/:id", middleware.BindUriMiddleware[models.IDRequest], app.ImageRemoveView)
	r.GET("image_cloud/options", app.ImageOptionsListView)

}
