package node

import (
	"encoding/json"

	"github.com/shirou/gopsutil/net"
)


func GetNetstatInfo() (string, error) {
	netstatInfo, _ := net.Connections("all")
	
	output, _ := json.Marshal(netstatInfo)
	return string(output), nil
}