package image_cloud_api

import (
	"github.com/gin-gonic/gin"
	"image_server/internal/global"
	"image_server/internal/middleware"
	"image_server/internal/models"
	"image_server/internal/utils/resp"
)

func (ImageCloudApi) ImageRemoveView(c *gin.Context) {
	cr := middleware.GetBind[models.IDRequest](c)
	log := middleware.GetLog(c)
	var model models.ImageModel
	err := global.DB.Preload("ServiceList").Take(&model, cr.ID).Error
	if err != nil {
		resp.FailWithMsg("镜像不存在", c)
		return
	}
	if len(model.ServiceList) > 0 {
		resp.FailWithMsg("镜像存在虚拟服务，请先删除关联的虚拟服务", c)
		return
	}

	log.Infof("删除镜像 %#v", model)

	err = global.DB.Delete(&model).Error
	if err != nil {
		log.Errorf("删除镜像失败 %s", err)
		resp.FailWithMsg("镜像删除失败", c)
		return
	}
	resp.OkWithMsg("删除成功", c)
}
