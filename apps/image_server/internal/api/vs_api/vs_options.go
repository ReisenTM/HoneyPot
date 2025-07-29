package vs_api

import (
	"Honeypot/apps/image_server/internal/global"
	"Honeypot/apps/image_server/internal/models"
	"Honeypot/apps/image_server/internal/utils/resp"
	"fmt"
	"github.com/gin-gonic/gin"
)

type VsOptionsListResponse struct {
	Label   string `json:"label"`
	Value   uint   `json:"value"`
	Disable bool   `json:"disable"`
}

func (VsApi) VsOptionsListView(c *gin.Context) {
	var list []models.ServiceModel
	global.DB.Find(&list)
	var options []VsOptionsListResponse
	for _, model := range list {
		item := VsOptionsListResponse{
			Label: fmt.Sprintf("%s/%d", model.Title, model.Port),
			Value: model.ID,
		}
		if model.Status != 1 {
			item.Disable = true
		}
		options = append(options, item)
	}
	resp.OkWithData(options, c)
}
