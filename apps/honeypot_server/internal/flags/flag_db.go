package flags

import (
	"Honeypot/apps/honeypot_server/internal/global"
	models2 "Honeypot/apps/honeypot_server/internal/models"
	"github.com/sirupsen/logrus"
)

// Migrate 数据库迁移
func Migrate() {
	err := global.DB.AutoMigrate(
		&models2.TrapIPModel{},
		&models2.TrapPortModel{},
		&models2.HostModel{},
		&models2.HostTemplateModel{},
		&models2.ImageModel{},
		&models2.LogModel{},
		&models2.MatrixTemplateModel{},
		&models2.NetModel{},
		&models2.NodeModel{},
		&models2.NodeNetworkModel{},
		&models2.ServiceModel{},
		&models2.UserModel{},
	)
	if err != nil {
		logrus.Fatalf("表结构迁移失败 %s", err)
	}
	logrus.Infof("表结构迁移成功")
}
