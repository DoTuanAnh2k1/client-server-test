package metric

import (
	"log"
	"os/exec"
	"runtime"
	"server/utils"
	"strings"
	"strconv"
)

type MetricPod struct {
	Name string `json:"name"`
	Ip string `json:"ip"`
	NumberOfGoroutine int `json:"number_of_goroutine"`
	NumberOfThread int `json:"number_of_thread"`
	NumberOfProcess int `json:"number_of_process"`
	NumberOfCpu int `json:"number_of_cpu"`
	MemUsage int `json:"mem_usage"`
	CpuUsage int `json:"cpu_usage"`
}

func GetMetric() *MetricPod {
	numGoroutines := runtime.NumGoroutine()
	numCpu := runtime.NumCPU()

	
	return &MetricPod{
		Name:              utils.GetHostName(),
		Ip:                utils.GetLocalIP(),
		NumberOfGoroutine: numGoroutines,
		NumberOfCpu:       numCpu,
		MemUsage:          int(memUsage),
	}
}

func getMemCpuUsage() (float64, float64, error) {
	cmd, err := exec.Command("top", "-b", "-n", "1", "|", "grep", "main").Output()
	if err != nil {
		log.Println(err)
		return 0, 0, err
	}

	var cpuPercentTotal, memPercentTotal float64
	lines := strings.Split(string(cmd), "\n")
	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) > 10 {
			cpuPercent, err := strconv.ParseFloat(parts[8], 64)
			if err != nil {
				return 0, 0, err
			}

			memPercent, err := strconv.ParseFloat(parts[9], 64)
			if err != nil {
				return 0, 0, err
			}

			cpuPercentTotal += cpuPercent
			memPercentTotal += memPercent
		}
	}

	memLimits := utils.GetEnv("MY_MEM")

	return cpuPercentTotal, memPercentTotal, nil
}

/*
func GetMetric() *MetricPod {
	

	return &MetricPod {
		Name: utils.GetHostName(),
		Ip: utils.GetLocalIP(),
		NumberOfGoroutine: runtime.NumGoroutine(),
		NumberOfCpu: runtime.NumCPU(),
	}
}
*/