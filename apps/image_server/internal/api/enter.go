package api

import (
	"Honeypot/apps/image_server/internal/api/image_cloud_api"
	"Honeypot/apps/image_server/internal/api/vnet_api"
	"Honeypot/apps/image_server/internal/api/vs_api"
)

type Api struct {
	ImageCloudApi image_cloud_api.ImageCloudApi
	VsApi         vs_api.VsApi
	VNet          vnet_api.VNetApi
}

var App = new(Api)
