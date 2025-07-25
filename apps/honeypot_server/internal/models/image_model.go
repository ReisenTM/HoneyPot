package models

// 镜像

type ImageModel struct {
	Model
	ImageName     string `json:"image_name"`
	Title         string `json:"title"`
	Port          int    `json:"port"`
	DockerImageID string `gorm:"size:32" json:"docker_image_id"`
	Tag           string `json:"tag"`
	Protocol      int8   `json:"protocol"`   //协议：TCP/HTTP...
	ImagePath     string `json:"image_path"` // 镜像文件
	Status        int8   `json:"status"`
	Logo          string `json:"logo"` // 镜像的logo
	Desc          string `json:"desc"` // 镜像描述
}
