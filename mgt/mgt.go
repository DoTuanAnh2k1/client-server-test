package main

import (
	"encoding/json"
	"io"
	"net"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type PodInfo struct {
	Ip          string `json:"ip"`
	NumberOfReq uint64 `json:"number-of-request"`
	InitRequest uint64 `json:"init-req"`
}

type ServerResp struct {
	Request     uint64 `json:"request"`
	InitRequest uint64 `json:"init-req"`
}

func getEnv(key, defaultVal string) string {
	val := os.Getenv(key)
	if val == "" {
		val = defaultVal
	}
	return val
}

func lookupListIp(headlessSvc string) []string {
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

func getInfo(c *gin.Context) {
	headlessSvc := getEnv("ServerHeadlessSvc", "localhost")
	serverPort := getEnv("ServerPort", "3654")
	ipList := lookupListIp(headlessSvc)
	var podInfoList []PodInfo
	for _, ip := range ipList {
		url := "http://" + ip + ":" + serverPort + "/info"
		client := http.Client{}
		resp, err := client.Get(url)
		if err != nil {
			panic(err)
		}
		var serverResp ServerResp
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		err = json.Unmarshal(body, &serverResp)
		if err != nil {
			panic(err)
		}
		podInfoList = append(podInfoList, PodInfo{
			Ip:          ip,
			NumberOfReq: serverResp.Request,
			InitRequest: serverResp.InitRequest,
		})
	}
	c.JSON(200, podInfoList)
}

func newGin() *gin.Engine {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.String(200, "ok")
	})
	r.GET("/init", func(c *gin.Context) {
		c.String(200, "ok")
	})
	r.GET("/info", getInfo)

	return r
}

func main() {
	server := newGin()
	server.Run(":1234")
}
