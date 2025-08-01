package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"honeypot_node/internal/core"
	"honeypot_node/internal/global"
	"honeypot_node/internal/rpc"
	"honeypot_node/internal/rpc/node_rpc"
	ip2 "honeypot_node/internal/utils/ip"
	"os"
)

// 全局节点客户端实例
func main() {
	global.Config = core.ReadConfig()
	global.Log = core.GetLogger()
	core.SetLogDefault()
	addr := global.Config.System.GrpcManagerAddr
	// 使用 grpc.Dial 创建一个到指定地址的 gRPC 连接。
	// 此处使用不安全的证书来实现 SSL/TLS 连接

	conn := rpc.GetConnWithTLS(addr)
	defer conn.Close()
	//初始化客户端
	client := node_rpc.NewNodeServiceClient(conn)
	nic := global.Config.System.Network
	ip, mac, err := ip2.GetNetworkInfo(nic)
	if err != nil {
		logrus.Fatalf(fmt.Sprintf("get network info [%s] error [%s]", nic, err))
	}
	if global.Config.System.Uid == "" {
		//如果没uid就创建并回写config
		uid := uuid.New().String()
		global.Config.System.Uid = uid
		core.SetConfig()
	}
	hostname, _ := os.Hostname()
	res, err := client.Register(context.Background(), &node_rpc.RegisterRequest{
		Ip:      ip,
		Mac:     mac,
		NodeUid: global.Config.System.Uid,
		Version: global.Version,
		Commit:  global.Commit,
		SystemInfo: &node_rpc.SystemInfoMessage{
			HostName: hostname,
		},
		//NetworkList: &node_rpc.NetworkInfoMessage{
		//	Network: "",
		//	Ip:      "",
		//	Net:     "",
		//	Mask:    0,
		//},
	})
	fmt.Println(res, err)

}
