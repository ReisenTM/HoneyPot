package models

//诱捕端口

type TrapPortModel struct {
	Model
	NodeID       uint         `json:"nodeID"`
	NodeModel    NodeModel    `gorm:"foreignKey:NodeID" json:"-"`
	NetID        uint         `json:"net_id"`
	NetModel     NetModel     `gorm:"foreignKey:NetID" json:"-"`
	TrapIPID     uint         `json:"trap_ip_id"`
	HoneyIpModel TrapIPModel  `gorm:"foreignKey:HoneyIpID" json:"-"`
	ServiceID    uint         `json:"service_id"` // 服务id
	ServiceModel ServiceModel `gorm:"foreignKey:ServiceID" json:"-"`
	Port         int          `json:"port"`                  // 服务的端口
	DstIP        string       `gorm:"size:32" json:"dst_ip"` // 目标ip
	DstPort      int          `json:"dst_port"`              // 目标端口
	Status       int8         `json:"status"`
}
