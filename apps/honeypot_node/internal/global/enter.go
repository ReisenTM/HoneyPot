package global

import (
	"github.com/sirupsen/logrus"
	"honeypot_node/internal/config"
)

var (
	Version   = "v1.0.1"
	Commit    = "7805a04452"
	BuildTime = "Wed Jul 30 14:48:19 2025 "
)
var (
	Config *config.Config
	Log    *logrus.Entry
)
