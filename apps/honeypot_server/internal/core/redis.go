package core

import (
	"Honeypot/apps/honeypot_server/internal/global"
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"sync"
)

func InitRedis() (client *redis.Client) {
	conf := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.Addr,
		Password: conf.Password,
		DB:       conf.DB,
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		logrus.Fatalf("连接redis失败 %s", err)
		return
	}
	logrus.Infof("成功连接redis")
	return rdb
}

var client *redis.Client
var onceRedis sync.Once

// GetRedis 获得redis单例
func GetRedis() *redis.Client {
	onceRedis.Do(func() {
		client = InitRedis()
	})
	return client
}
