package grpc_service

import (
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"honeypot_server/internal/global"
	"honeypot_server/internal/rpc/node_rpc"
	"net"
)

type NodeService struct {
	node_rpc.UnimplementedNodeServiceServer
}

func Run() {
	// 监听端口
	addr := global.Config.System.GrpcAddr
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		logrus.Fatalf("Failed to listen: %v", err)
	}

	// 创建 gRPC 服务器
	s := grpc.NewServer()

	// 创建一个gRPC服务器实例。
	server := NodeService{}
	// 将server结构体注册为gRPC服务。
	node_rpc.RegisterNodeServiceServer(s, &server)
	logrus.Infof("grpc 服务器监听于 %s", addr)
	// 开始处理客户端请求。
	err = s.Serve(listen)
}
