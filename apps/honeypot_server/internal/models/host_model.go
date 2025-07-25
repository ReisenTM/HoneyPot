package models

//主机

type HostModel struct {
	Model
	NodeID      uint      `json:"node_id"`
	NodeModel   NodeModel `gorm:"foreignKey:NodeID" json:"-"`
	NetID       uint      `json:"net_id"` //归属的网络id
	NetModel    NetModel  `gorm:"foreignKey:NetID" json:"-"`
	IP          string    `gorm:"32" json:"ip"`
	Mac         string    `gorm:"64" json:"mac"`
	Manufacture string    `gorm:"64" json:"manuf"` // 网卡生产厂商信息
}
