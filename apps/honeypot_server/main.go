package main

import (
	core2 "honeypot_server/internal/core"
	"honeypot_server/internal/flags"
	"honeypot_server/internal/global"
	"honeypot_server/internal/routers"
	"honeypot_server/internal/service/grpc_service"
)

var (
	Version   = "v1.0.1"
	Commit    = " 7805a04452"
	BuildTime = "Wed Jul 30 14:48:19 2025 "
)

func main() {
	global.Config = core2.ReadConfig()
	core2.InitIPDB()
	core2.SetLogDefault() //方便本地调试
	global.DB = core2.GetDB()
	global.Redis = core2.GetRedis()
	global.Log = core2.GetLogger()
	go grpc_service.Run()
	flags.Run()
	routers.Run()
}
