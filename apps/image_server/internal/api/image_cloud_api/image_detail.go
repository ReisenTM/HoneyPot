package image_cloud_api

import (
	"Honeypot/apps/image_server/internal/global"
	"Honeypot/apps/image_server/internal/middleware"
	"Honeypot/apps/image_server/internal/models"
	"Honeypot/apps/image_server/internal/utils/resp"
	"github.com/gin-gonic/gin"
)

func (ImageCloudApi) ImageDetailView(c *gin.Context) {
	cr := middleware.GetBind[models.IDRequest](c)
	var model models.ImageModel
	err := global.DB.Take(&model, cr.ID).Error
	if err != nil {
		resp.FailWithMsg("镜像不存在", c)
		return
	}

	resp.OkWithData(model, c)
}
