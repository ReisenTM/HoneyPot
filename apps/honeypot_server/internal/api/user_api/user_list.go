package user_api

import (
	"Honeypot/apps/honeypot_server/internal/middleware"
	models2 "Honeypot/apps/honeypot_server/internal/models"
	"Honeypot/apps/honeypot_server/internal/service/common_service"
	"Honeypot/apps/honeypot_server/internal/utils/resp"
	"github.com/gin-gonic/gin"
)

type UserListRequest struct {
	Username string `form:"username"`
	models2.PageInfo
}

func (UserApi) UserListView(c *gin.Context) {
	cr := middleware.GetBind[UserListRequest](c)
	list, count, _ := common_service.ListQuery(models2.UserModel{Username: cr.Username}, common_service.ListQueryOption{
		Likes:    []string{"username"},
		PageInfo: cr.PageInfo,
		OrderBy:  "created_at desc",
	})
	resp.OkWithList(list, count, c)
}
