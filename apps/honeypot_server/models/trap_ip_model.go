package models

// 诱捕ip

type TrapIPModel struct {
	Model
	NodeID    uint      `json:"node_id"`
	NodeModel NodeModel `gorm:"foreignKey:NodeID" json:"-"`
	NetID     uint      `json:"net_id"`
	NetModel  NetModel  `gorm:"foreignKey:NetID" json:"-"`
	IP        string    `gorm:"32" json:"ip"`
	Mac       string    `gorm:"64" json:"mac"`
	NIC       string    `gorm:"32" json:"nic"` // 网卡
	Status    int8      `json:"status"`        // 1 创建中 2 运行中 3 失败  4 删除中
	ErrorMsg  string    `json:"error_msg"`
}
