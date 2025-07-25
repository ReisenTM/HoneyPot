package main

import (
	"Honeypot/apps/honeypot_server/core"
	"Honeypot/apps/honeypot_server/flags"
	"Honeypot/apps/honeypot_server/global"
	"Honeypot/apps/honeypot_server/routers"
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
