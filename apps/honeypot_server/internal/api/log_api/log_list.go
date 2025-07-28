package log_api

import (
	middleware2 "Honeypot/apps/honeypot_server/internal/middleware"
	models2 "Honeypot/apps/honeypot_server/internal/models"
	common_service2 "Honeypot/apps/honeypot_server/internal/service/common_service"
	"Honeypot/apps/honeypot_server/internal/utils/resp"
	"fmt"
	"github.com/gin-gonic/gin"
)

type LogListRequest struct {
	IP       string `form:"ip"`
	Type     int8   `form:"type"`
	Location string `form:"location"`
	models2.PageInfo
}

// LogListView 日志列表
func (LogApi) LogListView(c *gin.Context) {
	cr := middleware2.GetBind[LogListRequest](c)
	list, count, _ := common_service2.ListQuery(models2.LogModel{
		Type:     cr.Type,
		IP:       cr.IP,
		Location: cr.Location,
	}, common_service2.ListQueryOption{
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
	cr := middleware2.GetBind[LogRemoveRequest](c)
	log := middleware2.GetLog(c)
	count, err := common_service2.Remove(models2.LogModel{}, common_service2.RemoveOption{
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
