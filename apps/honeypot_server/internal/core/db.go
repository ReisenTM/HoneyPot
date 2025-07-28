package core

import (
	"Honeypot/apps/honeypot_server/internal/global"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
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

	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.ConnMaxLifetime) * time.Second)
	logrus.Infof("数据库连接成功\n")
	return
}

var db *gorm.DB
var onceDb sync.Once

// GetDB 单例模式:如果有就返回之前创建的，没有就创建
func GetDB() *gorm.DB {
	onceDb.Do(func() {
		db = InitDB()
	})
	return db
}
