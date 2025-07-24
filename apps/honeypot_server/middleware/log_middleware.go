package middleware

import (
	""
	"Honeypot/apps/honeypot_server/global"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func LogMiddleware(c *gin.Context) {
	log := global.Log
	uid := uuid.New().String()
	logger := log.WithField("logID", uid)
	c.Set("log", logger)
}

func GetLog(c *gin.Context) *logrus.Entry {
	return c.MustGet("log").(*logrus.Entry)
}
