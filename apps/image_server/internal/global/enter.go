package global

import (
	"github.com/docker/docker/client"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"image_server/internal/config"
)

var (
	Version = "v1.0.1"
)
var (
	DB           *gorm.DB
	Config       *config.Config
	Log          *logrus.Entry
	Redis        *redis.Client
	DockerClient *client.Client
)
