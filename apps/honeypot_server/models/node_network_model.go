package models

//节点网络

type NodeNetworkModel struct {
	Model
	NodeID     uint      `json:"node_id"`
	NodeModel  NodeModel `gorm:"foreignKey:NodeID" json:"-"`
	NIC        string    `gorm:"32" json:"nic"`     //网卡 network interface card
	IP         string    `gorm:"32" json:"ip"`      // 探针ip
	SubnetMask int8      `json:"subnet_mask"`       // 子网掩码 8-32
	Gateway    string    `gorm:"32" json:"gateway"` //网关
	Status     int8      `json:"status"`            // 是否启用 1 启用 2 未启用
}
