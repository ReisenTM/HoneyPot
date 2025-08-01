package rpc

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"os"
)

func GetConnWithTLS(addr string) (conn *grpc.ClientConn) {
	// 加载客户端证书和私钥
	cert, err := tls.LoadX509KeyPair("internal/cert/client.crt", "internal/cert/client.key")
	if err != nil {
		logrus.Fatalf("failed to load client key pair: %v", err)
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
		RootCAs:      caCertPool,
	}

	// 创建 credentials
	creds := credentials.NewTLS(config)

	conn, err = grpc.NewClient(addr, grpc.WithTransportCredentials(creds))
	if err != nil {
		logrus.Fatalf(fmt.Sprintf("grpc connect addr [%s] 连接失败 %s", addr, err))
	}
	return
}
