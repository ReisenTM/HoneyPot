package flags

import (
	"Honeypot/apps/honeypot_server/global"
	"Honeypot/apps/honeypot_server/models"
	"github.com/sirupsen/logrus"
)

// Migrate 数据库迁移
func Migrate() {
	err := global.DB.AutoMigrate(
		&models.TrapIPModel{},
		&models.TrapPortModel{},
		&models.HostModel{},
		&models.HostTemplateModel{},
		&models.ImageModel{},
		&models.LogModel{},
		&models.MatrixTemplateModel{},
		&models.NetModel{},
		&models.NodeModel{},
		&models.NodeNetworkModel{},
		&models.ServiceModel{},
		&models.UserModel{},
	)
	if err != nil {
		logrus.Fatalf("表结构迁移失败 %s", err)
	}
	logrus.Infof("表结构迁移成功")
}
