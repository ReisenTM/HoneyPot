package main

import (
	"fmt"
	"honeypot_node/internal/utils/ip"
)

func main() {
	nic := "en0"
	fmt.Println(ip.GetNetworkInfo(nic))
}
