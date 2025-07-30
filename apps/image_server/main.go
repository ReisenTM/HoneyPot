package main

import (
	core "image_server/internal/core"
	"image_server/internal/flags"
	"image_server/internal/global"
	"image_server/internal/routers"
	"image_server/internal/service/vnet_service"
)

func main() {
	global.Config = core.ReadConfig()
	core.InitIPDB()
	core.SetLogDefault() //方便本地调试
	global.DB = core.GetDB()
	global.Redis = core.GetRedis()
	global.DockerClient = core.InitDocker()
	global.Log = core.GetLogger()
	flags.Run()
	vnet_service.Run()
	routers.Run()
}
