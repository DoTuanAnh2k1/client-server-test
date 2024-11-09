package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"time"

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
		client := http.Client{
			Timeout: 2 * time.Second,
		}
		resp, err := client.Get(url)
		if err != nil {
			fmt.Println("err: ", err)
			continue
		}
		var serverResp ServerResp
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("err: ", err)
			continue
		}
		err = json.Unmarshal(body, &serverResp)
		if err != nil {
			fmt.Println("err: ", err)
			continue
		}
		resp.Body.Close()
		podInfoList = append(podInfoList, PodInfo{
			Ip:          ip,
			NumberOfReq: serverResp.Request,
			InitRequest: serverResp.InitRequest,
		})
	}
	c.JSON(200, podInfoList)
}

func triggerOnHeadless(c *gin.Context) {
	clientSvc := getEnv("ClientSvc", "localhost")
	clientPort := getEnv("ClientPort", "3317")
	url := "http://" + clientSvc + ":" + clientPort + "/send-req"
	client := http.Client{}
	client.Get(url)
}

func triggerOffHeadless(c *gin.Context) {
	clientSvc := getEnv("ClientSvc", "localhost")
	clientPort := getEnv("ClientPort", "3317")
	url := "http://" + clientSvc + ":" + clientPort + "/off-send-req"
	client := http.Client{}
	client.Get(url)
}

func triggerOnService(c *gin.Context) {
	clientSvc := getEnv("ClientSvc", "localhost")
	clientPort := getEnv("ClientPort", "3317")
	url := "http://" + clientSvc + ":" + clientPort + "/send-req-svc"
	client := http.Client{}
	client.Get(url)
}

func triggerOffService(c *gin.Context) {
	clientSvc := getEnv("ClientSvc", "localhost")
	clientPort := getEnv("ClientPort", "3317")
	url := "http://" + clientSvc + ":" + clientPort + "/off-send-req-svc"
	client := http.Client{}
	client.Get(url)
}

func triggerOnServiceInitClient(c *gin.Context) {
	clientSvc := getEnv("ClientSvc", "localhost")
	clientPort := getEnv("ClientPort", "3317")
	url := "http://" + clientSvc + ":" + clientPort + "/send-req-svc-init-client"
	client := http.Client{}
	client.Get(url)
}

func triggerOffServiceInitClient(c *gin.Context) {
	clientSvc := getEnv("ClientSvc", "localhost")
	clientPort := getEnv("ClientPort", "3317")
	url := "http://" + clientSvc + ":" + clientPort + "/off-send-req-svc-init-client"
	client := http.Client{}
	client.Get(url)
}

func newGin() *gin.Engine {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.String(200, "ok")
	})
	r.GET("/info", getInfo)
	r.GET("/trigger-svc", triggerOnService)
	r.GET("/trigger-off-svc", triggerOffService)
	r.GET("/trigger-svc-init-client", triggerOnServiceInitClient)
	r.GET("/trigger-off-svc-init-client", triggerOffServiceInitClient)
	r.GET("/trigger-headless-svc", triggerOnHeadless)
	r.GET("/trigger-off-headless-svc", triggerOffHeadless)
	return r
}

func main() {
	server := newGin()
	server.Run(":1234")
}
