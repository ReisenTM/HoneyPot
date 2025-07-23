package main

import (
	"Honeypot/apps/honeypot_server/core"
	"Honeypot/apps/honeypot_server/flags"
	"Honeypot/apps/honeypot_server/global"
)

func main() {

	global.Config = core.InitConfig()
	global.DB = core.InitDB()
	core.SetLogDefault() //方便本地调试
	global.Log = core.GetLogger()
	flags.Run()
}
