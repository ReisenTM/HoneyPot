package flags

import (
	"github.com/sirupsen/logrus"
	"image_server/internal/global"
	models2 "image_server/internal/models"
)

// Migrate 数据库迁移
func Migrate() {
	err := global.DB.AutoMigrate(

		&models2.HostTemplateModel{},
		&models2.ImageModel{},
		&models2.MatrixTemplateModel{},
		&models2.ServiceModel{},
	)
	if err != nil {
		logrus.Fatalf("表结构迁移失败 %s", err)
	}
	logrus.Infof("表结构迁移成功")
}
