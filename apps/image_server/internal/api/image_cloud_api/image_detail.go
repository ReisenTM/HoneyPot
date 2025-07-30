package image_cloud_api

import (
	"github.com/gin-gonic/gin"
	"image_server/internal/global"
	"image_server/internal/middleware"
	"image_server/internal/models"
	"image_server/internal/utils/resp"
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
