package models

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"image_server/internal/utils/cmd"
)

//服务

type ServiceModel struct {
	Model
	Title         string     `json:"title"`
	Protocol      int8       `json:"protocol"`
	ImageID       uint       `json:"imageID"`
	ImageModel    ImageModel `gorm:"foreignKey:ImageID" json:"-"`
	IP            string     `json:"ip"`
	Port          int        `json:"port"`
	Status        int8       `json:"status"`
	ErrorMsg      string     `json:"error_msg"`
	TrapIPCount   int        `json:"trap_ip_count"`  //关联诱捕ip数量
	ContainerID   string     `json:"container_id"`   // 容器id
	ContainerName string     `json:"container_name"` // 容器名

}

func (s *ServiceModel) State() string {
	switch s.Status {
	case 1:
		return "running"
	}
	return "error"
}

func (s *ServiceModel) BeforeDelete(tx *gorm.DB) error {
	// 判断是否有关联的端口转发
	var count int64
	tx.Model(TrapPortModel{}).Where("service_id = ?", s.ID).Count(&count)
	if count > 0 {
		return errors.New("存在端口转发，不能删除虚拟服务")
	}

	command := fmt.Sprintf("docker rm -f %s", s.ContainerName)
	err := cmd.Cmd(command)
	if err != nil {
		logrus.Errorf("删除容器失败 %s", err)
		return err
	}
	return nil
}
