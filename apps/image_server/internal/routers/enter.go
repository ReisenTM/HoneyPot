package routers

import (
	"Honeypot/apps/image_server/internal/global"
	middleware2 "Honeypot/apps/image_server/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Run() {
	sysConf := global.Config.System
	gin.SetMode(sysConf.Mode)

	r := gin.Default()
	//静态路由,资源映射
	//r.Static("uploads", "./uploads")
	g := r.Group("image_server")
	g.Use(middleware2.LogMiddleware, middleware2.AuthMiddleware) //需要放行的使用白名单机制

	ImageCloudRouter(g)

	logrus.Infof("服务器监听于 %s\n", sysConf.WebAddr)
	_ = r.Run(sysConf.WebAddr)
}
