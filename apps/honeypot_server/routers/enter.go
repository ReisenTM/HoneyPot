package routers

import (
	"Honeypot/apps/honeypot_server/global"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Run() {
	sysConf := global.Config.System
	gin.SetMode(sysConf.Mode)

	r := gin.Default()
	//静态路由,资源映射
	r.Static("uploads", "./uploads")
	g := r.Group("honeypot_server")
	g.Use()
	logrus.Infof("服务器监听于 %s\n", sysConf.WebAddr)
	_ = r.Run(sysConf.WebAddr)
}
