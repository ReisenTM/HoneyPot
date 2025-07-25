package models

// 矩阵模版

type MatrixTemplateModel struct {
	Model
	Title            string           `gorm:"size:32" json:"title"`
	HostTemplateList HostTemplateList `gorm:"serializer:json" json:"host_template_list"` // 主机模板列表
}
type HostTemplateList []HostTemplateInfo
type HostTemplateInfo struct {
	HostTemplateID uint `json:"host_template_id"`
	Weight         int  `json:"weight"` //权重
}
