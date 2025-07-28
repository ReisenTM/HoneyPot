package image_cloud_api

import (
	"Honeypot/apps/image_server/internal/global"
	"Honeypot/apps/image_server/internal/middleware"
	"Honeypot/apps/image_server/internal/models"
	"Honeypot/apps/image_server/internal/utils/path"
	"Honeypot/apps/image_server/internal/utils/resp"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"path/filepath"
)

type ImageCreateRequest struct {
	ImageID   string `json:"image_id" binding:"required"`             // 镜像id
	ImageName string `json:"image_name" binding:"required"`           // 镜像的名称
	ImageTag  string `json:"image_tag" binding:"required"`            // 镜像的tag
	ImagePath string `json:"image_path" binding:"required"`           // 镜像上传的路径
	Title     string `json:"title" binding:"required"`                // 镜像别名
	Port      int    `json:"port" binding:"required,min=1,max=65535"` // 镜像端口
	Protocol  int8   `json:"protocol" binding:"required,oneof=1"`     // 镜像的协议
}

func (ImageCloudApi) ImageCreateView(c *gin.Context) {
	cr := middleware.GetBind[ImageCreateRequest](c)

	// 1. 检查镜像文件是否存在
	if _, err := os.Stat(cr.ImagePath); errors.Is(err, os.ErrNotExist) {
		resp.FailWithMsg("镜像文件不存在", c)
		return
	}

	// 2. 检查镜像title是否重名
	var titleExists models.ImageModel
	if err := global.DB.Take(&titleExists, "title = ?", cr.Title).Error; err == nil {
		resp.FailWithMsg("镜像别名不能重复", c)
		return
	}

	// 3. 检查镜像名+tag是否重复
	var nameTagExists models.ImageModel
	if err := global.DB.Take(&nameTagExists, "image_name = ? AND tag = ?", cr.ImageName, cr.ImageTag).Error; err == nil {
		resp.FailWithMsg("镜像名称和标签组合不能重复", c)
		return
	}

	// 4. 使用docker load命令导入镜像
	cmd := exec.Command("docker", "load", "-i", cr.ImagePath)
	// 移动到我们的项目路径下
	cmd.Dir = path.GetRootPath()
	output, err := cmd.CombinedOutput()
	if err != nil {
		resp.FailWithMsg(fmt.Sprintf("镜像导入失败: %s, 输出: %s", err.Error(), string(output)), c)
		return
	}
	fmt.Println(string(output))

	// 5. 移动镜像文件到正式目录
	finalDir := "uploads/images/"

	// 确保目标目录存在
	if err := os.MkdirAll(finalDir, 0755); err != nil {
		resp.FailWithMsg(fmt.Sprintf("创建目标目录失败: %s", err.Error()), c)
		return
	}

	// 获取文件名
	_, fileName := filepath.Split(cr.ImagePath)
	finalPath := filepath.Join(finalDir, fileName)

	// 移动文件
	if err := os.Rename(cr.ImagePath, finalPath); err != nil {
		// 如果移动失败，尝试复制后删除
		logrus.Errorf("文件移动失败 %s", err)
		resp.FailWithMsg("文件移动失败", c)
		return
	}

	// 6. 数据入库
	imageModel := models.ImageModel{
		DockerImageID: cr.ImageID,
		ImageName:     cr.ImageName,
		Tag:           cr.ImageTag,
		ImagePath:     finalPath,
		Title:         cr.Title,
		Port:          cr.Port,
		Protocol:      cr.Protocol,
		Status:        1,
	}

	if err := global.DB.Create(&imageModel).Error; err != nil {
		resp.FailWithMsg(fmt.Sprintf("数据库插入失败: %s", err.Error()), c)
		return
	}

	resp.Ok(imageModel.ID, "镜像创建成功", c)
}
