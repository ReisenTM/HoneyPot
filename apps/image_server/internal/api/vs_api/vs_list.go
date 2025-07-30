package vs_api

import (
	"github.com/gin-gonic/gin"
	"image_server/internal/middleware"
	"image_server/internal/models"
	"image_server/internal/service/common_service"
	"image_server/internal/utils/resp"
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
