package user_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	middleware2 "honeypot_server/internal/middleware"
	"honeypot_server/internal/models"
	"honeypot_server/internal/service/common_service"
	"honeypot_server/internal/utils/resp"
)

type UserRemoveRequest struct {
	IDList []uint `json:"id_list"`
}

func (UserApi) UserRemoveView(c *gin.Context) {
	cr := middleware2.GetBind[UserRemoveRequest](c)
	log := middleware2.GetLog(c)
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
