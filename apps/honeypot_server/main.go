package main

import (
	core2 "Honeypot/apps/honeypot_server/internal/core"
	"Honeypot/apps/honeypot_server/internal/flags"
	"Honeypot/apps/honeypot_server/internal/global"
	"Honeypot/apps/honeypot_server/internal/routers"
)

func main() {
	global.Config = core2.ReadConfig()
	core2.InitIPDB()
	core2.SetLogDefault() //方便本地调试
	global.DB = core2.GetDB()
	global.Redis = core2.GetRedis()
	global.Log = core2.GetLogger()
	flags.Run()
	routers.Run()
}
