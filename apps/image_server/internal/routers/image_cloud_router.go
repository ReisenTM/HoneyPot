package routers

import (
	"github.com/gin-gonic/gin"
	"image_server/internal/api"
	"image_server/internal/api/image_cloud_api"
	"image_server/internal/middleware"
	"image_server/internal/models"
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
