package vnet_api

import (
	"Honeypot/apps/image_server/internal/config"
	"Honeypot/apps/image_server/internal/core"
	"Honeypot/apps/image_server/internal/global"
	"Honeypot/apps/image_server/internal/middleware"
	"Honeypot/apps/image_server/internal/models"
	"Honeypot/apps/image_server/internal/utils/cmd"
	"Honeypot/apps/image_server/internal/utils/resp"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type VNetApi struct {
}

func (VNetApi) VsNetInfoView(c *gin.Context) {
	resp.OkWithData(global.Config.VsNet, c)
}

type VsNetRequest struct {
	Name   string `json:"name" binding:"required"`
	Prefix string `json:"prefix" binding:"required"`
	Net    string `json:"net" binding:"required"`
}

func (VNetApi) VsNetUpdateView(c *gin.Context) {
	cr := middleware.GetBind[VsNetRequest](c)

	// 在没有虚拟服务的情况下才能创建
	var serviceList []models.ServiceModel
	global.DB.Find(&serviceList)
	if len(serviceList) != 0 {
		resp.FailWithMsg("存在虚拟服务，不可修改虚拟子网", c)
		return
	}

	// 把之前的删掉
	command := fmt.Sprintf("docker network rm %s", global.Config.VsNet.Name)
	err := cmd.Cmd(command)
	if err != nil {
		logrus.Errorf("删除之前的虚拟网络失败 %s", err)
		resp.FailWithMsg("删除之前的虚拟网络失败", c)
		return
	}

	// 创建新的
	command = fmt.Sprintf("docker network create --driver bridge --subnet %s %s",
		cr.Net, cr.Name)
	err = cmd.Cmd(command)
	if err != nil {
		logrus.Errorf("创建虚拟网络失败 %s", err)
		resp.FailWithMsg("创建虚拟网络失败", c)
		return
	}

	// 回写到配置文件
	global.Config.VsNet = config.VsNet{
		Name:   cr.Name,
		Prefix: cr.Prefix,
		Net:    cr.Net,
	}
	core.SetConfig()

	resp.OkWithMsg("修改虚拟网络成功", c)
}
