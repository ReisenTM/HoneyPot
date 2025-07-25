package user_api

import (
	"Honeypot/apps/honeypot_server/middleware"
	"Honeypot/apps/honeypot_server/models"
	"Honeypot/apps/honeypot_server/service/common_service"
	"Honeypot/apps/honeypot_server/utils/resp"
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserRemoveRequest struct {
	IDList []uint `json:"id_list"`
}

func (UserApi) UserRemoveView(c *gin.Context) {
	cr := middleware.GetBind[UserRemoveRequest](c)
	log := middleware.GetLog(c)
	successCount, err := common_service.Remove(models.UserModel{}, common_service.RemoveOption{
		IDList: cr.IDList,
		Log:    log,
		Msg:    "用户",
	})
	if err != nil {
		msg := fmt.Sprintf("删除用户失败 %s", err)
		resp.FailWithMsg(msg, c)
		return
	}
	msg := fmt.Sprintf("删除成功 共%d个，成功%d个", len(cr.IDList), successCount)
	resp.OkWithMsg(msg, c)
}
