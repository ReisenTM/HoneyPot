package docker_service

import (
	"Honeypot/apps/image_server/internal/global"
	"context"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
)

// RunContainer 使用dockerAPI创建容器
func RunContainer(containerName, networkName, ip, image string) (containerID string, err error) {
	// 创建容器配置
	containerConfig := &container.Config{
		Image: image,
	}
	hostConfig := &container.HostConfig{
		AutoRemove:  false,
		NetworkMode: container.NetworkMode(networkName),
		RestartPolicy: container.RestartPolicy{
			Name: "always",
		},
	}
	networkingConfig := &network.NetworkingConfig{
		EndpointsConfig: map[string]*network.EndpointSettings{
			networkName: {
				IPAMConfig: &network.EndpointIPAMConfig{
					IPv4Address: ip,
				},
			},
		},
	}

	// 创建容器
	createResp, err := global.DockerClient.ContainerCreate(
		context.Background(),
		containerConfig,
		hostConfig,
		networkingConfig,
		nil,
		containerName,
	)
	if err != nil {
		return
	}

	// 启动容器
	err = global.DockerClient.ContainerStart(context.Background(), createResp.ID, container.StartOptions{})
	if err != nil {
		return
	}
	containerID = createResp.ID[:12]
	return
}
