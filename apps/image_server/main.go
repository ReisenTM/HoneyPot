package main

import (
	core "Honeypot/apps/image_server/internal/core"
	"Honeypot/apps/image_server/internal/flags"
	"Honeypot/apps/image_server/internal/global"
	"Honeypot/apps/image_server/internal/routers"
)

func main() {
	global.Config = core.ReadConfig()
	core.InitIPDB()
	core.SetLogDefault() //方便本地调试
	global.DB = core.GetDB()
	global.Redis = core.GetRedis()
	global.Log = core.GetLogger()
	flags.Run()
	routers.Run()
}
