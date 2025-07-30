package api

import (
	"image_server/internal/api/host_template_api"
	"image_server/internal/api/image_cloud_api"
	"image_server/internal/api/matrix_template_api"
	"image_server/internal/api/vnet_api"
	"image_server/internal/api/vs_api"
)

type Api struct {
	ImageCloudApi     image_cloud_api.ImageCloudApi
	VsApi             vs_api.VsApi
	VNetApi           vnet_api.VNetApi
	HostTemplateApi   host_template_api.HostTemplateApi
	MatrixTemplateApi matrix_template_api.MatrixTemplateApi
}

var App = new(Api)
