package ip

import (
	"fmt"
	"net"
)

// GetNetworkInfo 获取网卡信息
func GetNetworkInfo(i string) (ip string, mac string, err error) {
	iface, err := net.InterfaceByName(i)
	if err != nil {
		err = fmt.Errorf("无法获取网卡 %s: %v", i, err)
		return
	}

	// 获取网卡的地址列表
	addrs, err := iface.Addrs()
	if err != nil {
		err = fmt.Errorf("无法获取网卡 %s 的地址: %s", iface.Name, err)
		return
	}

	mac = iface.HardwareAddr.String()

	// 打印每个IP地址
	for _, addr := range addrs {
		var _ip net.IP
		switch v := addr.(type) {
		case *net.IPNet:
			_ip = v.IP
		case *net.IPAddr:
			_ip = v.IP
		}
		// 检查是否为IPv4地址
		if _ip.To4() != nil {
			ip = _ip.String()
		}
	}
	if ip == "" {
		err = fmt.Errorf("%s 此接口无ip的地址", iface.Name)
		return
	}
	return
}
