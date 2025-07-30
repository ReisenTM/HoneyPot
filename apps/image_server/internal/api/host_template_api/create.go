package host_template_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"image_server/internal/global"
	"image_server/internal/middleware"
	"image_server/internal/models"
	"image_server/internal/utils/resp"
)

type CreateRequest struct {
	Title    string                      `json:"title" binding:"required"`
	PortList models.HostTemplatePortList `json:"port_list" binding:"dive"`
}

func (HostTemplateApi) CreateView(c *gin.Context) {
	cr := middleware.GetBind[CreateRequest](c)

	var model models.HostTemplateModel
	err := global.DB.Take(&model, "title = ? ", cr.Title).Error
	if err == nil {
		resp.FailWithMsg("主机模板名称不能重复", c)
		return
	}
	// 校验服务id
	// 校验端口不能重复
	var serviceIDList []uint
	var portMap = map[int]bool{}
	for _, port := range cr.PortList {
		serviceIDList = append(serviceIDList, port.ServiceID)
		portMap[port.Port] = true
	}

	if len(portMap) != len(cr.PortList) {
		resp.FailWithMsg("端口存在重复", c)
		return
	}
	var serviceList []models.ServiceModel
	global.DB.Debug().Find(&serviceList, "id in ?", serviceIDList)
	var serviceMap = map[uint]models.ServiceModel{}
	for _, serviceModel := range serviceList {
		serviceMap[serviceModel.ID] = serviceModel
	}
	for _, port := range cr.PortList {
		_, ok := serviceMap[port.ServiceID]
		if !ok {
			msg := fmt.Sprintf("虚拟服务 %d 不存在", port.ServiceID)
			resp.FailWithMsg(msg, c)
			return
		}
	}

	// 消息入库
	model = models.HostTemplateModel{
		Title:    cr.Title,
		PortList: cr.PortList,
	}
	err = global.DB.Create(&model).Error
	if err != nil {
		resp.FailWithMsg("主机模板创建失败", c)
		return
	}
	resp.OkWithData(model.ID, c)
}
