package node

import (
	"encoding/json"

	psutil "github.com/shirou/gopsutil/process"
)



type Process struct {
	PID int32 `json:"pid"`
	PPID int32 `json:"ppid"`
	Cmd string `json:"cmd"`
	CreateTime int64 `json:"create_time"`
	Status string `json:"status"`
	Username string `json:"username"`
	CPU float64 `json:"cpu"`
	Mem float32 `json:"mem"`
}

func GetProcessList() (string, error) {

	processes, err := psutil.Processes()
	if err != nil {
		return "", err
	}

	list := []Process{}

	for _, process := range processes {
		ppid, _ := process.Ppid()
		cmd, _ := process.Cmdline()
		createTime, _ := process.CreateTime()
		status, _ := process.Status()
		user, _ := process.Username()
		cpu, _ := process.CPUPercent()
		mem, _ := process.MemoryPercent()
		// TODO: add connections
		// connections, _ := process.Connections()

		list = append(list, Process{
			PID: int32(process.Pid),
			PPID: ppid,
			Cmd: cmd,
			CreateTime: createTime,
			Status: status,
			Username: user,
			CPU: cpu,
			Mem: mem,
		})
	}

	output, _ := json.Marshal(list)
	return string(output), nil
}