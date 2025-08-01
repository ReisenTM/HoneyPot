package image_cloud_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"image_server/internal/global"
	"image_server/internal/models"
	"image_server/internal/utils/resp"
)

type ImageOptionsListResponse struct {
	Label   string `json:"label"`
	Value   uint   `json:"value"`
	Disable bool   `json:"disable"`
}

func (ImageCloudApi) ImageOptionsListView(c *gin.Context) {
	var list []models.ImageModel
	global.DB.Find(&list)
	var options []ImageOptionsListResponse
	for _, model := range list {
		item := ImageOptionsListResponse{
			Label: fmt.Sprintf("%s/%d", model.Title, model.Port),
			Value: model.ID,
		}
		if model.Status == 2 {
			item.Disable = true
		}
		options = append(options, item)
	}
	resp.OkWithData(options, c)
}
