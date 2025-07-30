package vs_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"image_server/internal/global"
	"image_server/internal/middleware"
	"image_server/internal/models"
	"image_server/internal/utils/resp"
)

func (VsApi) VsRemoveView(c *gin.Context) {
	cr := middleware.GetBind[models.IDListRequest](c)

	var serviceList []models.ServiceModel
	global.DB.Find(&serviceList, "id in ?", cr.IdList)
	if len(serviceList) == 0 {
		resp.FailWithMsg("不存在的虚拟服务", c)
		return
	}

	result := global.DB.Delete(&serviceList)
	successCount := result.RowsAffected
	err := result.Error
	if err != nil {
		resp.FailWithMsg("删除虚拟服务失败", c)
		return
	}

	msg := fmt.Sprintf("删除虚拟服务成功 共%d个", successCount)
	resp.OkWithMsg(msg, c)
}
