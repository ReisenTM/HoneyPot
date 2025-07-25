package ip

import "net"

// HasLocalIPAddr 判断ip是否是本地ip
func HasLocalIPAddr(_ip string) bool {
	ip := net.ParseIP(_ip)
	if ip.IsPrivate() {
		return true
	}
	if ip.IsLoopback() {
		return true
	}
	return false
}
