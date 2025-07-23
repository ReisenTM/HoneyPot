package main

import (
	"Honeypot/apps/honeypot_server/core"
	"Honeypot/apps/honeypot_server/flags"
	"Honeypot/apps/honeypot_server/global"
)

func main() {
	global.Config = core.InitConfig()
	global.DB = core.InitDB()
	flags.Run()
}
