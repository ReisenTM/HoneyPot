package models

type TrapPortModel struct {
	Model
	NodeID    uint   `json:"node_id"`
	NetID     uint   `json:"net_id"`
	TrapIpID  uint   `json:"trap_ip_id"`
	ServiceID uint   `json:"service_id"`            // 服务id
	Port      int    `json:"port"`                  // 服务的端口
	DstIP     string `gorm:"size:32" json:"dst_ip"` // 目标ip
	DstPort   int    `json:"dst_port"`              // 目标端口
	Status    int8   `json:"status"`
}
