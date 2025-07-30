package global

import (
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"honeypot_server/internal/config"
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
