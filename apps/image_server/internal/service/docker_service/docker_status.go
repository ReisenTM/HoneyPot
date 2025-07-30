package docker_service

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"image_server/internal/global"
)

// ListAllContainers 列出所有Docker容器的状态
func ListAllContainers() ([]container.Summary, error) {
	// 获取所有容器列表（包括停止的容器）
	containers, err := global.DockerClient.ContainerList(context.Background(), container.ListOptions{
		All: true,
	})
	if err != nil {
		return nil, fmt.Errorf("获取容器列表失败: %v", err)
	}

	return containers, nil
}

// PrefixContainerStatus 根据容器名前缀获取容器状态
func PrefixContainerStatus(containerName string) (summaryList []container.Summary, err error) {
	// 使用过滤器按名称查找容器
	filter := filters.NewArgs()
	filter.Add("name", containerName)

	containers, err := global.DockerClient.ContainerList(context.Background(), container.ListOptions{
		Filters: filter,
		All:     true,
	})
	if err != nil {
		return
	}
	// 返回第一个匹配的容器（容器名应该是唯一的）
	return containers, nil
}
