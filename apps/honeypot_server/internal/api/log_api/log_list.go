package log_api

import (
	"Honeypot/apps/honeypot_server/middleware"
	"Honeypot/apps/honeypot_server/models"
	"Honeypot/apps/honeypot_server/service/common_service"
	"Honeypot/apps/honeypot_server/utils/resp"
	"fmt"
	"github.com/gin-gonic/gin"
)

type LogListRequest struct {
	IP       string `form:"ip"`
	Type     int8   `form:"type"`
	Location string `form:"location"`
	models.PageInfo
}

// LogListView 日志列表
func (LogApi) LogListView(c *gin.Context) {
	cr := middleware.GetBind[LogListRequest](c)
	list, count, _ := common_service.ListQuery(models.LogModel{
		Type:     cr.Type,
		IP:       cr.IP,
		Location: cr.Location,
	}, common_service.ListQueryOption{
		Likes:    []string{"username"},
		PageInfo: cr.PageInfo,
		OrderBy:  "created_at desc",
	})
	resp.OkWithList(list, count, c)
}

type LogRemoveRequest struct {
	IDList []uint `json:"id_list"`
}

// LogRemoveView 日志批量删除
func (LogApi) LogRemoveView(c *gin.Context) {
	cr := middleware.GetBind[LogRemoveRequest](c)
	log := middleware.GetLog(c)
	count, err := common_service.Remove(models.LogModel{}, common_service.RemoveOption{
		IDList:   cr.IDList,
		Log:      log,
		Unscoped: true, //开启真删除
		Msg:      "日志",
	})
	if err != nil {
		resp.FailWithMsg("删除失败", c)
		return
	}
	msg := fmt.Sprintf("日志删除成功: %d 条\n", count)
	resp.OkWithMsg(msg, c)
}
