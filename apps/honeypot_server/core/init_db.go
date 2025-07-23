package core

import (
	"Honeypot/apps/honeypot_server/global"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func InitDB() (db *gorm.DB) {
	cfg := global.Config.DB
	db, err := gorm.Open(mysql.Open(cfg.Dsn()), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 不生成实体外键
	})
	if err != nil {
		logrus.Fatalf("数据库连接失败 %s", err)
		return
	}
	// 配置连接池
	sqlDB, err := db.DB()
	if err != nil {
		logrus.Fatalf("获取数据库连接失败 %s", err)
		return
	}
	err = sqlDB.Ping()
	if err != nil {
		logrus.Fatalf("数据库连接失败 %s", err)
		return
	}
	// 设置连接池
	if cfg.MaxIdleConns == 0 {
		cfg.MaxIdleConns = 10
	}
	if cfg.MaxOpenConns == 0 {
		cfg.MaxOpenConns = 100
	}
	if cfg.ConnMaxLifetime == 0 {
		cfg.ConnMaxLifetime = 10000
	}
	//logrus.Infof("最大空闲数 %d", cfg.MaxIdleConns)
	//logrus.Infof("最大连接数 %d", cfg.MaxOpenConns)
	//logrus.Infof("超时时间 %s", time.Duration(cfg.ConnMaxLifetime)*time.Second)
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.ConnMaxLifetime) * time.Second)
	logrus.Infof("数据库连接成功")
	return
}
