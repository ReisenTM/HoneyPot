package matrix_template_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"image_server/internal/global"
	"image_server/internal/middleware"
	"image_server/internal/models"
	"image_server/internal/utils/resp"
)

type UpdateRequest struct {
	ID               uint                      `json:"id" binding:"required"`
	Title            string                    `json:"title" binding:"required"`
	HostTemplateList []models.HostTemplateInfo `json:"host_template_list" binding:"required,dive,required"` // 主机模板列表
}

func (MatrixTemplateApi) UpdateView(c *gin.Context) {
	cr := middleware.GetBind[UpdateRequest](c)

	var model models.MatrixTemplateModel
	err := global.DB.Take(&model, cr.ID).Error
	if err != nil {
		resp.FailWithMsg("矩阵模板不存在", c)
		return
	}

	if len(cr.HostTemplateList) == 0 {
		resp.FailWithMsg("矩阵模板需要关联至少一个主机模板", c)
		return
	}

	var newModel models.MatrixTemplateModel
	err = global.DB.Take(&newModel, "title = ? and id <> ?", cr.Title, cr.ID).Error
	if err == nil {
		resp.FailWithMsg("修改的矩阵模板名称不能重复", c)
		return
	}
	var hostTemplateIDList []uint
	for _, h := range cr.HostTemplateList {
		hostTemplateIDList = append(hostTemplateIDList, h.HostTemplateID)
	}

	var hostTemps []models.HostTemplateModel
	global.DB.Find(&hostTemps, "id in ?", hostTemplateIDList)
	var hostTempMap = map[uint]models.HostTemplateModel{}
	for _, m := range hostTemps {
		hostTempMap[m.ID] = m
	}
	for _, h := range cr.HostTemplateList {
		_, ok := hostTempMap[h.HostTemplateID]
		if !ok {
			msg := fmt.Sprintf("主机模板 %d 不存在", h.HostTemplateID)
			resp.FailWithMsg(msg, c)
			return
		}
	}

	newModel = models.MatrixTemplateModel{
		Title:            cr.Title,
		HostTemplateList: cr.HostTemplateList,
	}
	err = global.DB.Model(&model).Updates(newModel).Error
	if err != nil {
		resp.FailWithMsg("矩阵模板修改失败", c)
		return
	}
	resp.OkWithMsg("矩阵模板修改成功", c)
}
