package node

import (
	"encoding/json"

	ps "github.com/shirou/gopsutil/host"
)

func GetHostInfo() string {

	hostInfo, _ := ps.Info()
	
	output, _ := json.Marshal(hostInfo)
	return string(output)
}