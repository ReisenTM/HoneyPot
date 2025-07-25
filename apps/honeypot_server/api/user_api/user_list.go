package user_api

import (
	"Honeypot/apps/honeypot_server/middleware"
	"Honeypot/apps/honeypot_server/models"
	"Honeypot/apps/honeypot_server/service/common_service"
	"Honeypot/apps/honeypot_server/utils/resp"
	"github.com/gin-gonic/gin"
)

type UserListRequest struct {
	Username string `form:"username"`
	models.PageInfo
}

func (UserApi) UserListView(c *gin.Context) {
	cr := middleware.GetBind[UserListRequest](c)
	list, count, _ := common_service.ListQuery(models.UserModel{Username: cr.Username}, common_service.ListQueryOption{
		Likes:    []string{"username"},
		PageInfo: cr.PageInfo,
		OrderBy:  "created_at desc",
	})
	resp.OkWithList(list, count, c)
}
