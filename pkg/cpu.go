package node

import (
	"encoding/json"

	"github.com/shirou/gopsutil/cpu"
)

func GetCPUInfo() string {
	cpuInfo, _ := cpu.Info()
	
	output, _ := json.Marshal(cpuInfo)
	return string(output)
}