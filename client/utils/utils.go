package utils

import (
	"net"
	"os"
)

func LookupListIp(headlessSvc string) []string {
	ips, err := net.LookupIP(headlessSvc)
	if err != nil {
		panic(err)
	}
	var listIp []string
	for _, v := range ips {
		listIp = append(listIp, v.String())
	}
	return listIp
}

func GetEnv(key, defaultVal string) string {
	val := os.Getenv(key)
	if val == "" {
		val = defaultVal
	}
	return val
}
