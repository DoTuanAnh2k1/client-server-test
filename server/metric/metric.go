package metric

import (
	"log"
	"math"
	"os/exec"
	"runtime"
	"server/utils"
	"strconv"
	"strings"
)

type PodMetric struct {
	Name              string  `json:"name"`
	Ip                string  `json:"ip"`
	NumberOfGoroutine int     `json:"number_of_goroutine"`
	NumberOfThread    int     `json:"number_of_thread"`
	NumberOfProcess   int     `json:"number_of_process"`
	NumberOfCpu       int     `json:"number_of_cpu"`
	MemUsage          float64 `json:"mem_usage"`
	CpuUsage          float64 `json:"cpu_usage"`
}

func GetMetric() (*PodMetric, error) {
	var metricPod PodMetric
	metricPod.Name = utils.GetHostName()
	metricPod.Ip = utils.GetLocalIP()
	metricPod.NumberOfGoroutine = runtime.NumGoroutine()
	metricPod.NumberOfCpu = runtime.NumCPU()

	cpuUsage, memUsage, err := getMemCpuUsage()
	if err != nil {
		return metricPod, err
	}

	metricPod.CpuUsage = cpuUsage
	metricPod.MemUsage = memUsage

	return &metricPod, nil
}

func getMemCpuUsage() (float64, float64, error) {
	// cmd, err := exec.Command("top", "-b", "-n", "1", "|", "grep", "main").Output()
	cmd, err := exec.Command("bash", "-c", "top -b -n 1 | grep main").Output()
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

	memLimitsStr := utils.GetEnv("MY_MEM_LIMIT", "4294967296")
	memLimits, err := strconv.ParseFloat(memLimitsStr, 64)
	if err != nil {
		return 0, 0, err
	}

	memUsage := (memLimits * (memPercentTotal * 100)) / 10000
	memUsage = memUsage / 1024578
	memUsage = math.Round(memUsage*100) / 100

	return cpuPercentTotal, memUsage, nil
}
