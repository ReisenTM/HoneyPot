package image_cloud_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"image_server/internal/utils/docker"
	"image_server/internal/utils/resp"
	"os"
	"path/filepath"
	"time"
)

type ImageUploadResponse struct {
	ImageID   string `json:"image_id"`   // 镜像id
	ImageName string `json:"image_name"` // 镜像的名称
	ImageTag  string `json:"image_tag"`  // 镜像的tag
	ImagePath string `json:"image_path"` // 镜像上传的路径
}

const (
	maxFileSize  = 2 << 30 // 2GB
	tempImageDir = "uploads/images_temp/"
)

func (ImageCloudApi) ImageUploadView(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		resp.FailWithMsg("请选择镜像文件", c)
		return
	}

	// Check file size
	if file.Size > maxFileSize {
		resp.FailWithMsg("镜像文件大小不能超过2GB", c)
		return
	}

	// Check file extension
	ext := filepath.Ext(file.Filename)
	if ext != ".tar" && ext != ".gz" {
		resp.FailWithMsg("只支持.tar和.tar.gz格式的镜像文件", c)
		return
	}

	// Create temp directory if not exists
	if err := os.MkdirAll(tempImageDir, 0755); err != nil {
		resp.FailWithMsg(fmt.Sprintf("创建临时目录失败: %v", err), c)
		return
	}

	// Save uploaded file to temp location
	tempFilePath := filepath.Join(tempImageDir, file.Filename)
	if err := c.SaveUploadedFile(file, tempFilePath); err != nil {
		resp.FailWithMsg(fmt.Sprintf("保存镜像文件失败: %v", err), c)
		return
	}

	// Parse image metadata
	imageID, imageName, imageTag, err := docker.ParseImageMetadata(tempFilePath)
	if err != nil {
		os.Remove(tempFilePath)
		resp.FailWithMsg(fmt.Sprintf("解析镜像元数据失败: %v", err), c)
		return
	}

	go func() {
		time.Sleep(5 * time.Minute)
		//五分钟后迁移到对应目录
		err = os.Remove(tempFilePath)
		if os.IsNotExist(err) {
			return
		}
		if err != nil {
			logrus.Errorf("镜像删除失败 %s", err)
		} else {
			logrus.Infof("删除镜像文件 %s", tempFilePath)
		}
	}()

	// Prepare response
	data := ImageUploadResponse{
		ImageID:   imageID,
		ImageName: imageName,
		ImageTag:  imageTag,
		ImagePath: tempFilePath,
	}

	resp.OkWithData(data, c)
}
