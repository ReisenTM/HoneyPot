package api

import "Honeypot/apps/image_server/internal/api/image_cloud_api"

type Api struct {
	ImageCloudApi image_cloud_api.ImageCloudApi
}

var App = new(Api)
