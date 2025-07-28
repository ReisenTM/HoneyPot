package image_cloud_api

import (
	"Honeypot/apps/image_server/internal/middleware"
	"Honeypot/apps/image_server/internal/models"
	"Honeypot/apps/image_server/internal/service/common_service"
	"Honeypot/apps/image_server/internal/utils/resp"
	"github.com/gin-gonic/gin"
)

type ImageListRequest struct {
	models.PageInfo
}

func (ImageCloudApi) ImageListView(c *gin.Context) {
	cr := middleware.GetBind[ImageListRequest](c)
	list, count, _ := common_service.ListQuery(models.ImageModel{},
		common_service.ListQueryOption{
			Likes:    []string{"title", "image_name"},
			PageInfo: cr.PageInfo,
			OrderBy:  "created_at desc",
		})
	resp.OkWithList(list, count, c)
}
