package core

import (
	"github.com/docker/docker/client"
	"github.com/sirupsen/logrus"
)

func InitDocker() *client.Client {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		logrus.Fatalf("创建Docker客户端失败: %v", err)
	}
	return cli
}
