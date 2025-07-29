package vs_api

import (
	"Honeypot/apps/image_server/internal/global"
	"Honeypot/apps/image_server/internal/middleware"
	"Honeypot/apps/image_server/internal/models"
	"Honeypot/apps/image_server/internal/service/docker_service"
	"Honeypot/apps/image_server/internal/utils/resp"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net"
	"time"
)

type VsCreateRequest struct {
	ImageID uint `json:"image_id" binding:"required"`
}

// 基础IP地址和子网掩码
const (
	maxIP = 254 // 最大可用IP 10.2.0.254
)

// 获取下一个可用IP
func getNextAvailableIP() (string, error) {
	ip, _, err := net.ParseCIDR(global.Config.VsNet.Net)
	if err != nil {
		return "", err
	}
	ip4 := ip.To4()
	// 查询数据库中已分配的最大IP
	var service models.ServiceModel
	err = global.DB.Order("ip DESC").First(&service).Error
	if err != nil {
		if err.Error() == "record not found" {
			// 没有记录，返回起始IP
			ip4[3] = 2
			return ip4.String(), nil
		}
		return "", fmt.Errorf("查询最大IP失败: %w", err)
	}

	serviceIP := net.ParseIP(service.IP)
	if serviceIP == nil {
		return "", fmt.Errorf("服务ip解析错误")
	}
	serviceIP4 := serviceIP.To4()

	// 检查是否达到最大IP
	if serviceIP4[3] >= maxIP {
		return "", fmt.Errorf("IP地址池已满")
	}
	// 生成新IP
	newLastOctet := serviceIP4[3] + 1
	ip4[3] = newLastOctet
	return ip4.String(), nil
}
func (VsApi) VsCreateView(c *gin.Context) {
	cr := middleware.GetBind[VsCreateRequest](c)

	var image models.ImageModel
	err := global.DB.Take(&image, cr.ImageID).Error
	if err != nil {
		resp.FailWithMsg("镜像不存在", c)
		return
	}
	if image.Status == 2 {
		resp.FailWithMsg("镜像不可用", c)
		return
	}

	// 判断这个镜像有没有跑过这个服务
	var service models.ServiceModel
	err = global.DB.Take(&service, "image_id = ?", cr.ImageID).Error
	if err == nil {
		resp.FailWithMsg("此镜像已运行虚拟服务", c)
		return
	}

	// 使用docker命令运行容器
	// docker network create --driver bridge --subnet 10.2.0.0/24 honey-hy
	// docker run -d --network honey-hy --ip 10.2.0.10 --name my_container image_name:tag
	// 获取下一个可用IP
	ip, err := getNextAvailableIP()
	if err != nil {
		logrus.Errorf("获取可用IP失败: %s", err)
		resp.FailWithMsg("IP地址池已满，无法创建新服务", c)
		return
	}
	// 如果是自己的逻辑，先找出最大的ip地址，如何下一个ip就是创建的ip
	fmt.Println(ip)
	networkName := global.Config.VsNet.Name
	containerName := global.Config.VsNet.Prefix + image.ImageName
	containerID, err := docker_service.RunContainer(containerName, networkName, ip, fmt.Sprintf("%s:%s", image.ImageName, image.Tag))
	if err != nil {
		logrus.Errorf("创建虚拟服务失败 %s", err)
		resp.FailWithMsg("创建虚拟服务失败", c)
		//TODO：如果有，删除创建好的容器
		return
	}

	command := fmt.Sprintf("docker run -d --network %s --ip %s --name %s %s:%s",
		networkName, ip, containerName, image.ImageName, image.Tag)
	fmt.Println(command)
	var model = models.ServiceModel{
		Title:         image.Title,
		ContainerName: containerName,
		Protocol:      image.Protocol,
		ImageID:       image.ID,
		IP:            ip,
		Port:          image.Port,
		Status:        1,
		ContainerID:   containerID,
	}
	err = global.DB.Create(&model).Error
	if err != nil {
		logrus.Errorf("创建虚拟服务失败 %s", err)
		resp.FailWithMsg("创建虚拟服务失败", c)
		return
	}

	go func(model *models.ServiceModel) {
		//检测容器是否运行正常
		var delayList = []<-chan time.Time{
			time.After(5 * time.Second),
			time.After(20 * time.Second),
			time.After(1 * time.Minute),
			time.After(5 * time.Minute),
			time.After(1 * time.Hour),
		}
		for _, times := range delayList {
			<-times
			ContainerStatus(model)
		}
	}(&model)

	resp.Ok(model.ID, "创建虚拟服务成功", c)
	return
}

func ContainerStatus(model *models.ServiceModel) {
	logrus.Infof("检测容器状态 %s", model.ContainerName)
	var newModel models.ServiceModel
	containers, err := docker_service.PrefixContainerStatus(model.ContainerName)
	var isUpdate bool
	var state string
	if err != nil {
		newModel.Status = 2
		newModel.ErrorMsg = err.Error()
		isUpdate = true
		state = err.Error()
	}
	if len(containers) != 1 {
		newModel.Status = 2
		newModel.ErrorMsg = "容器不存在"
		isUpdate = true
		state = newModel.ErrorMsg
	} else {
		container := containers[0]
		if container.State == "running" && model.Status != 1 {
			// 我们这边是不正常的，但是实际是正常的
			newModel.Status = 1
			newModel.ErrorMsg = ""
			isUpdate = true
			state = container.State
		}
		if container.State != "running" && model.Status == 1 {
			// 我们这边是正常的，但是实际是不正常的
			newModel.Status = 2
			newModel.ErrorMsg = fmt.Sprintf("%s(%s)", container.State, container.Status)
			isUpdate = true
			state = container.State
		}
	}

	if isUpdate {
		logrus.Infof("%s 容器存在状态修改 %s => %s", model.ContainerName, model.State(), state)
		global.DB.Model(model).Updates(map[string]any{
			"status":    newModel.Status,
			"error_msg": newModel.ErrorMsg,
		})
	}

}
