package matrix_template_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"image_server/internal/middleware"
	"image_server/internal/models"
	"image_server/internal/service/common_service"
	"image_server/internal/utils/resp"
)

func (MatrixTemplateApi) Remove(c *gin.Context) {
	cr := middleware.GetBind[models.IDListRequest](c)
	log := middleware.GetLog(c)
	successCount, err := common_service.Remove(models.MatrixTemplateModel{}, common_service.RemoveOption{
		IDList: cr.IdList,
		Log:    log,
		Msg:    "矩阵模板",
	})
	if err != nil {
		msg := fmt.Sprintf("删除矩阵模板失败 %s", err)
		resp.FailWithMsg(msg, c)
		return
	}
	msg := fmt.Sprintf("删除成功 共%d个，成功%d个", len(cr.IdList), successCount)
	resp.OkWithMsg(msg, c)
}
