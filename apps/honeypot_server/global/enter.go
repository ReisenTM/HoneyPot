package global

import (
	"Honeypot/apps/honeypot_server/config"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	Config *config.Config
)
