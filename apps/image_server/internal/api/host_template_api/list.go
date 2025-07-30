package host_template_api

import (
	"Honeypot/apps/image_server/internal/global"
	"Honeypot/apps/image_server/internal/middleware"
	"Honeypot/apps/image_server/internal/models"
	"Honeypot/apps/image_server/internal/service/common_service"
	"Honeypot/apps/image_server/internal/utils/resp"
	"github.com/gin-gonic/gin"
)

type ListRequest struct {
	models.PageInfo
}
type HostTemplatePortInfo struct {
	Port          int    `json:"port"`
	ServiceID     uint   `json:"service_id"`
	ServiceTitle  string `json:"service_title"`
	ServiceStatus int8   `json:"service_status"`
}
type ListResponse struct {
	models.HostTemplateModel
	PortList []HostTemplatePortInfo `json:"port_list"`
}

func (HostTemplateApi) ListView(c *gin.Context) {

	cr := middleware.GetBind[models.PageInfo](c)

	_list, count, _ := common_service.ListQuery(models.HostTemplateModel{},
		common_service.ListQueryOption{
			Likes:    []string{"title"},
			PageInfo: cr,
			OrderBy:  "created_at desc",
		})
	var list = make([]ListResponse, 0)
	var serviceList []models.ServiceModel
	var serviceIDList []uint
	for _, model := range _list {
		for _, port := range model.PortList {
			serviceIDList = append(serviceIDList, port.ServiceID)
		}
	}
	global.DB.Find(&serviceList, "id in ?", serviceIDList)
	var serviceMap = map[uint]models.ServiceModel{}
	for _, i2 := range serviceList {
		serviceMap[i2.ID] = i2
	}
	for _, model := range _list {
		portList := make([]HostTemplatePortInfo, 0)
		for _, port := range model.PortList {
			portList = append(portList, HostTemplatePortInfo{
				Port:          port.Port,
				ServiceID:     port.ServiceID,
				ServiceTitle:  serviceMap[port.ServiceID].Title,
				ServiceStatus: serviceMap[port.ServiceID].Status,
			})
		}
		list = append(list, ListResponse{
			HostTemplateModel: model,
			PortList:          portList,
		})
	}
	resp.OkWithList(list, count, c)
}
