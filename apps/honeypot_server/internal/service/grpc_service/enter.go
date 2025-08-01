package grpc_service

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"honeypot_server/internal/global"
	"honeypot_server/internal/rpc/node_rpc"
	"net"
	"os"
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
	// 加载服务端证书和私钥
	cert, err := tls.LoadX509KeyPair("internal/cert/server.crt", "internal/cert/server.key")
	if err != nil {
		logrus.Fatalf("failed to load key pair: %v", err)
	}

	// 加载 CA 证书
	caCert, err := os.ReadFile("internal/cert/ca.crt")
	if err != nil {
		logrus.Fatalf("failed to read CA certificate: %v", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// 创建 TLS 配置
	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert, // 双向认证
		ClientCAs:    caCertPool,
	}

	// 创建 credentials
	creds := credentials.NewTLS(config)

	// 创建 gRPC 服务器，使用 TLS credentials
	s := grpc.NewServer(grpc.Creds(creds))

	//// 创建 gRPC 服务器
	//s := grpc.NewServer()

	// 创建一个gRPC服务器实例。
	server := NodeService{}
	// 将server结构体注册为gRPC服务。
	node_rpc.RegisterNodeServiceServer(s, &server)
	logrus.Infof("grpc 服务器监听于 %s", addr)
	// 开始处理客户端请求。
	err = s.Serve(listen)
}
