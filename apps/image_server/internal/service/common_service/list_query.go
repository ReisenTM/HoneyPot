package common_service

import (
	"Honeypot/apps/image_server/internal/core"
	"Honeypot/apps/image_server/internal/models"
	"fmt"
	"gorm.io/gorm"
)

type ListQueryOption struct {
	Preload  []string        `json:"preload"`
	Likes    []string        `json:"likes"`
	PageInfo models.PageInfo `json:"page_info"`
	OrderBy  string          `json:"order_by"`
	Where    *gorm.DB        `json:"where"`
	Debug    bool            `json:"debug"`
}

func ListQuery[T any](model T, opt ListQueryOption) (list []T, count int64, err error) {
	db := core.GetDB()
	if opt.Debug {
		db = db.Debug()
	}

	// 针对字段的精确匹配
	db = db.Where(model)
	// 高级查询
	if opt.Where != nil {
		db = db.Where(opt.Where)
	}

	//预加载
	for _, s := range opt.Preload {
		db = db.Preload(s)
	}

	//模糊匹配字段
	if opt.PageInfo.Key != "" {
		like := core.GetDB().Where("")
		for _, c := range opt.Likes {
			//关键字模糊匹配为或关系
			like.Or(fmt.Sprintf("%s like ?", c), fmt.Sprintf("%%%s%%", opt.PageInfo.Key))
		}
		db = db.Where(like)
	}
	// 分页
	if opt.PageInfo.Limit <= 0 {
		opt.PageInfo.Limit = 10
	}
	if opt.PageInfo.Page <= 0 {
		opt.PageInfo.Page = 1
	}
	offset := (opt.PageInfo.Page - 1) * opt.PageInfo.Limit
	err = db.Offset(offset).Limit(opt.PageInfo.Limit).Order(opt.OrderBy).Find(&list).Error
	err = db.Count(&count).Error
	return
}
