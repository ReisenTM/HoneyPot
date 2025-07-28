package image_cloud_api

import (
	"Honeypot/apps/image_server/internal/global"
	"Honeypot/apps/image_server/internal/middleware"
	"Honeypot/apps/image_server/internal/models"
	"Honeypot/apps/image_server/internal/utils/resp"
	"github.com/gin-gonic/gin"
)

type ImageUpdateRequest struct {
	ID       uint   `json:"id"`
	Title    string `json:"title" binding:"required"`
	Port     int    `json:"port" binding:"required,min=1,max=65535"`
	Protocol int8   `json:"protocol" binding:"required,oneof=1"`
	Status   int8   `json:"status" binding:"required,oneof=1 2"` // 1 成功  2 禁用
	Logo     string `json:"logo"`                                // 镜像的logo
	Desc     string `json:"desc"`                                // 镜像描述
}

func (ImageCloudApi) ImageUpdateView(c *gin.Context) {
	cr := middleware.GetBind[ImageUpdateRequest](c)

	var model models.ImageModel
	err := global.DB.Take(&model, cr.ID).Error
	if err != nil {
		resp.FailWithMsg("镜像不存在", c)
		return
	}

	// title不能和现在的重名
	var newModel models.ImageModel
	err = global.DB.Take(&newModel, "id <> ? and title = ?", cr.ID, cr.Title).Error
	if err == nil {
		resp.FailWithMsg("修改的镜像名称不能重复", c)
		return
	}

	err = global.DB.Model(&model).Updates(models.ImageModel{
		Title:    cr.Title,
		Port:     cr.Port,
		Protocol: cr.Protocol,
		Status:   cr.Status,
		Logo:     cr.Logo,
		Desc:     cr.Desc,
	}).Error
	if err != nil {
		resp.FailWithMsg("镜像更新失败", c)
		return
	}

	resp.OkWithMsg("镜像修改成功", c)
}
