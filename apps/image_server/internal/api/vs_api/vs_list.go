package vs_api

import (
	"Honeypot/apps/image_server/internal/middleware"
	"Honeypot/apps/image_server/internal/models"
	"Honeypot/apps/image_server/internal/service/common_service"
	"Honeypot/apps/image_server/internal/utils/resp"
	"github.com/gin-gonic/gin"
)

type VsListRequest struct {
	models.PageInfo
	Port  int    `form:"port"`
	IP    string `form:"ip"`
	Title string `form:"title"`
}

func (VsApi) VsListView(c *gin.Context) {
	cr := middleware.GetBind[VsListRequest](c)
	list, count, _ := common_service.ListQuery(models.ServiceModel{
		Title: cr.Title,
		IP:    cr.IP,
		Port:  cr.Port,
	}, common_service.ListQueryOption{
		Likes:    []string{"title"},
		PageInfo: cr.PageInfo,
		OrderBy:  "created_at desc",
	})
	resp.OkWithList(list, count, c)
}
