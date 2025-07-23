package models

// 节点

type NodeModel struct {
	Model
	Title       string         `gorm:"size:32" json:"title"`
	Uid         string         `gorm:"size:64" json:"uid"`
	IP          string         `gorm:"size:32" json:"ip"`
	Mac         string         `gorm:"size:64" json:"mac"`
	Status      int8           `json:"status"`
	NetCount    int            `json:"net_count"`
	TrapIPCount int            `json:"trap_ip_count"`
	Resource    NodeResource   `gorm:"serializer:json" json:"resource"`
	SystemInfo  NodeSystemInfo `gorm:"serializer:json" json:"system_info"`
}

type NodeResource struct {
	CpuCount              int     `json:"cpu_count"`
	CpuUsageRate          float64 `json:"cpu_usage_rate"`
	MemTotal              int64   `json:"mem_total"`
	MemUsageRate          float64 `json:"mem_usage_rate"`
	DiskTotal             int64   `json:"disk_total"`
	DiskUseRate           float64 `json:"disk_use_rate"`
	NodePath              string  `json:"node_path"`               // 节点的部署目录
	NodeResourceOccupancy int64   `json:"node_resource_occupancy"` // 节点目录的资源占用
}

type NodeSystemInfo struct {
	HostName            string `json:"host_name"`
	DistributionVersion string `json:"distribution_version"` // 发行版本
	CoreVersion         string `json:"core_version"`         // 内核版本
	SystemType          string `json:"system_type"`          // 系统类型
	StartTime           string `json:"start_time"`           // 启动时间
	NodeVersion         string `json:"node_version"`
	NodeCommit          string `json:"node_commit"`
}
