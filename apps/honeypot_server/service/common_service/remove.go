package common_service

import (
	"Honeypot/apps/honeypot_server/global"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type RemoveOption struct {
	Debug    bool
	Where    *gorm.DB
	IDList   []uint
	Log      *logrus.Entry
	Msg      string
	Unscoped bool
}

func Remove[T any](model T, req RemoveOption) (successCount int64, err error) {
	db := global.DB
	deleteDB := global.DB
	if req.Debug {
		db = db.Debug()
		deleteDB = deleteDB.Debug()
	}
	
	if req.Unscoped {
		req.Log.Infof("启用真删除")
		deleteDB = deleteDB.Unscoped()
	}

	if req.Where != nil {
		db = db.Where(req.Where)
	}

	db = db.Where(model)

	if len(req.IDList) > 0 {
		req.Log.Infof("删除 %s idList %v", req.Msg, req.IDList)
		db = db.Where("id in ?", req.IDList)
	}

	var list []T
	db.Find(&list)

	if len(list) <= 0 {
		req.Log.Infof("没查到")
		return
	}

	result := deleteDB.Delete(&list)
	if result.Error != nil {
		req.Log.Errorf("删除失败 %s", result.Error)
		return
	}
	successCount = result.RowsAffected
	req.Log.Infof("删除 %s 成功, 成功%d个", req.Msg, successCount)
	return
}
