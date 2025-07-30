package matrix_template_api

import (
	"Honeypot/apps/image_server/internal/global"
	"Honeypot/apps/image_server/internal/middleware"
	"Honeypot/apps/image_server/internal/models"
	"Honeypot/apps/image_server/internal/utils/resp"
	"fmt"
	"github.com/gin-gonic/gin"
)

type CreateRequest struct {
	Title            string                    `json:"title" binding:"required"`
	HostTemplateList []models.HostTemplateInfo `json:"host_template_list" binding:"required,dive,required"` // 主机模板列表
}

func (MatrixTemplateApi) CreateView(c *gin.Context) {
	cr := middleware.GetBind[CreateRequest](c)

	if len(cr.HostTemplateList) == 0 {
		resp.FailWithMsg("矩阵模板需要关联至少一个主机模板", c)
		return
	}

	var model models.MatrixTemplateModel
	err := global.DB.Take(&model, "title = ? ", cr.Title).Error
	if err == nil {
		resp.FailWithMsg("矩阵模板名称不能重复", c)
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

	// 消息入库
	model = models.MatrixTemplateModel{
		Title:            cr.Title,
		HostTemplateList: cr.HostTemplateList,
	}
	err = global.DB.Create(&model).Error
	if err != nil {
		resp.FailWithMsg("矩阵模板创建失败", c)
		return
	}
	resp.OkWithData(model.ID, c)
}
