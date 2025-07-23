package global

import (
	"Honeypot/apps/honeypot_server/config"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	Version = "v1.0.1"
)
var (
	DB     *gorm.DB
	Config *config.Config
	Log    *logrus.Entry
	Redis  *redis.Client
)
