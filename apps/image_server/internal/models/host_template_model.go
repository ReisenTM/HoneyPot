package models

//主机模板

type HostTemplateModel struct {
	Model
	Title    string               `gorm:"size:32" json:"title"`
	PortList HostTemplatePortList `gorm:"serializer:json" json:"port_list"`
}
type HostTemplatePortList []HostTemplatePort
type HostTemplatePort struct {
	Port      int  `json:"port"`
	ServiceID uint `json:"service_id"`
}
