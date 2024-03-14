package node

import (
	"encoding/json"

	"github.com/shirou/gopsutil/mem"
)

func GetMemoryInfo() (string, error) {
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		return "", err
	}

	output, _ := json.Marshal(memInfo)
	return string(output), nil
}