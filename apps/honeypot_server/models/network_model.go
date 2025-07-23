package models

//网络管理

type NetModel struct {
	Model
	NodeID            uint      `json:"node_id"`
	NodeModel         NodeModel `gorm:"foreignKey:NodeID" json:"-"`
	Title             string    `gorm:"32" json:"title"`
	NIC               string    `gorm:"32" json:"nic"` // 网卡
	IP                string    `gorm:"32" json:"ip"`  // 探针ip
	SubnetMask        int8      `json:"subnet_mask"`   // 子网掩码 8-32
	Gateway           string    `gorm:"32" json:"gateway"`
	HostCount         int       `json:"host_count"` //资源数(主机数)
	TrapIPCount       int       `json:"trap_ip_count"`
	ScanStatus        int8      `json:"scan_status"`                     // 扫描状态  0 待扫描  1 扫描完成  2 扫描中
	ScanProgress      float64   `json:"scan_progress"`                   // 扫描进度
	UsableTrapIPRange string    `gorm:"256" json:"usable_trap_ip_range"` // 能够使用的诱捕ip范围
}
