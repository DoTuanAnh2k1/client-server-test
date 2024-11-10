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

var (
	ClientSvc         string
	ClientPort        string
	ServerHeadlessSvc string
	ServerPort        string
)

const (
	Service           = "svc"
	ServiceInitClient = "svc-init"
	ServiceHeadless   = "svc-headless"
	ServiceSol        = "svc-sol"
)

const (
	ProblemInit      = "init"
	ProblemReconnect = "reconnect"
	ProblemDo        = "do"
)

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
	ipList := lookupListIp(ServerHeadlessSvc)
	var podInfoList []PodInfo
	for _, ip := range ipList {
		url := "http://" + ip + ":" + ServerPort + "/info"
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

func triggerOn(c *gin.Context) {
	name := c.Param("name")
	var url string
	switch name {
	case Service:
		url = "http://" + ClientSvc + ":" + ClientPort + "/send-req-svc"
	case ServiceInitClient:
		url = "http://" + ClientSvc + ":" + ClientPort + "/send-req-svc-init-client"
	case ServiceHeadless:
		url = "http://" + ClientSvc + ":" + ClientPort + "/send-req"
	case ServiceSol:
		url = "http://" + ClientSvc + ":" + ClientPort + "/send-req-svc-sol"
	default:
		c.JSON(http.StatusBadRequest, "name unsupported")
		return
	}
	client := http.Client{}
	client.Get(url)
	c.JSON(http.StatusOK, "ok")
}

func triggerOff(c *gin.Context) {
	name := c.Param("name")
	var url string
	switch name {
	case Service:
		url = "http://" + ClientSvc + ":" + ClientPort + "/off-send-req-svc"
	case ServiceInitClient:
		url = "http://" + ClientSvc + ":" + ClientPort + "/off-send-req-svc-init-client"
	case ServiceHeadless:
		url = "http://" + ClientSvc + ":" + ClientPort + "/off-send-req"
	case ServiceSol:
		url = "http://" + ClientSvc + ":" + ClientPort + "/off-send-req-svc-sol"
	default:
		c.JSON(http.StatusBadRequest, "name unsupported")
		return
	}
	client := http.Client{}
	client.Get(url)
	c.JSON(http.StatusOK, "ok")
}

func triggerProblem(c *gin.Context) {
	name := c.Param("name")
	var url string
	switch name {
	case ProblemInit:
		url = "http://" + ClientSvc + ":" + ClientPort + "/problem-init"
	case ProblemReconnect:
		url = "http://" + ClientSvc + ":" + ClientPort + "/problem-reconnect"
	case ProblemDo:
		url = "http://" + ClientSvc + ":" + ClientPort + "/problem-do"
	default:
		c.JSON(http.StatusBadRequest, "name unsupported")
		return
	}
	client := http.Client{}
	client.Get(url)
	c.JSON(http.StatusOK, "ok")
}

func newGin() *gin.Engine {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.String(200, "ok")
	})
	r.GET("/info", getInfo)

	r.GET("/trigger/on/:name", triggerOn)
	r.GET("/trigger/off/:name", triggerOff)
	r.GET("/trigger/problem/:name", triggerProblem)

	return r
}

func main() {
	ClientSvc = getEnv("ClientSvc", "localhost")
	ClientPort = getEnv("ClientPort", "3317")
	ServerHeadlessSvc = getEnv("ServerHeadlessSvc", "localhost")
	ServerPort = getEnv("ServerPort", "3654")
	server := newGin()
	server.Run(":1234")
}
