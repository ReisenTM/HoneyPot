package models

import (
	"Honeypot/apps/image_server/internal/utils/cmd"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"os"
)

// 镜像

type ImageModel struct {
	Model
	ImageName     string         `json:"image_name"`
	Title         string         `json:"title"`
	Port          int            `json:"port"`
	DockerImageID string         `gorm:"size:32" json:"docker_image_id"`
	ServiceList   []ServiceModel `gorm:"foreignKey:ImageID" json:"-"` // 关联的虚拟服务列表
	Tag           string         `json:"tag"`
	Protocol      int8           `json:"protocol"`   //协议：TCP/HTTP...
	ImagePath     string         `json:"image_path"` // 镜像文件
	Status        int8           `json:"status"`
	Logo          string         `json:"logo"` // 镜像的logo
	Desc          string         `json:"desc"` // 镜像描述
}

func (i *ImageModel) BeforeDelete(tx *gorm.DB) error {
	// 删除docker镜像
	command := fmt.Sprintf("docker rmi %s", i.DockerImageID)
	err := cmd.Cmd(command)
	if err != nil {
		return err
	}
	// 删除镜像文件
	logrus.Infof("删除镜像文件 %s", i.ImagePath)
	err = os.Remove(i.ImagePath)
	if err != nil {
		return err
	}
	return nil
}
