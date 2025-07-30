package matrix_template_api

import (
	"Honeypot/apps/image_server/internal/global"
	"Honeypot/apps/image_server/internal/models"
	"Honeypot/apps/image_server/internal/utils/resp"
	"github.com/gin-gonic/gin"
)

type OptionsListResponse struct {
	Label string `json:"label"`
	Value uint   `json:"value"`
}

func (MatrixTemplateApi) OptionsView(c *gin.Context) {
	var list = make([]OptionsListResponse, 0)
	global.DB.Model(models.MatrixTemplateModel{}).Select("id as value", "title as label").Scan(&list)
	resp.OkWithData(list, c)
}
