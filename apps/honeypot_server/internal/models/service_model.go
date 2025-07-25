package models

//服务

type ServiceModel struct {
	Model
	Title       string     `json:"title"`
	Protocol    int8       `json:"protocol"`
	ImageID     uint       `json:"imageID"`
	ImageModel  ImageModel `gorm:"foreignKey:ImageID" json:"-"`
	IP          string     `json:"ip"`
	Port        int        `json:"port"`
	Status      int8       `json:"status"`
	TrapIPCount int        `json:"trap_ip_count"` //关联诱捕ip数量
	ContainerID string     `json:"container_id"`  // 容器id
}
