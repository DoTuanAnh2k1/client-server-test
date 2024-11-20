package server

import (
	"encoding/json"
	"fmt"
	"io"
	"mgt/common"
	"mgt/utils"
	"net"
	"net/http"
	"sort"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

func getInfo(c *gin.Context) {
	ipList := utils.LookupListIp(common.ServerHeadlessSvc)
	var podInfoList []common.PodInfo
	for _, ip := range ipList {
		url := "http://" + ip + ":" + common.ServerPort + "/info"
		client := http.Client{
			Timeout: 2 * time.Second,
		}
		resp, err := client.Get(url)
		if err != nil {
			fmt.Println("err: ", err)
			continue
		}
		var serverResp common.ServerResp
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
		podInfoList = append(podInfoList, common.PodInfo{
			Ip:          net.ParseIP(ip),
			NumberOfReq: serverResp.Request,
			InitRequest: serverResp.InitRequest,
			SuccessRate: serverResp.SuccessRate,
		})
	}
	sort.Slice(podInfoList, func(i, j int) bool { return podInfoList[i].Ip.String() < podInfoList[j].Ip.String() })
	c.JSON(200, podInfoList)
}

func getMeasure(c *gin.Context) {
	ipList := utils.LookupListIp(common.ServerHeadlessSvc)
	var podInfoList []common.PodInfo
	var wg sync.WaitGroup

	for _, ip := range ipList {
		wg.Add(1)
		go getPodMeasure(&wg, ip, &podInfoList)
	}
	wg.Wait()

	c.JSON(200, podInfoList)
}

func getPodMeasure(wg *sync.WaitGroup, podIp string, podInfoList *[]common.PodInfo) {
	var serverResp common.ServerMeasure
	url := "http://" + podIp + ":" + common.ServerPort + "/measure"
	client := http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	err = json.Unmarshal(body, &serverResp)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	resp.Body.Close()
	*podInfoList = append(*podInfoList, common.PodInfo{
		Ip:          net.ParseIP(podIp),
		NumberOfReq: serverResp.Request,
		SuccessRate: serverResp.SuccessRate,
	})
	wg.Done()
}

func triggerOn(c *gin.Context) {
	name := c.Param("name")
	url := "http://" + common.ClientSvc + ":" + common.ClientPort + "/trigger/on?name="
	switch name {
	case common.Service:
		url += ClientPathTriggerSvc
	case common.ServiceHeadless:
		url += ClientPathTriggerHeadlessSvc
	case common.RabbitMQ:
		url += ClientPathRabbitMQ
	case common.GRPC:
		url += ClientPathGRPC
	case common.Kafka:
		url += ClientPathKafka
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
	url := "http://" + common.ClientSvc + ":" + common.ClientPort + "/trigger/off?name="
	switch name {
	case common.Service:
		url += ClientPathTriggerSvc
	case common.ServiceHeadless:
		url += ClientPathTriggerHeadlessSvc
	case common.RabbitMQ:
		url += ClientPathRabbitMQ
	case common.GRPC:
		url += ClientPathGRPC
	case common.Kafka:
		url += ClientPathKafka
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
	url := "http://" + common.ClientSvc + ":" + common.ClientPort + "/prob?name="
	switch name {
	case common.ProblemInit:
		url += ClientPathProbInit
	case common.ProblemReconnect:
		url += ClientPathProbReconnect
	case common.ProblemDo:
		url += ClientPathProbDo
	default:
		c.JSON(http.StatusBadRequest, "name unsupported")
		return
	}
	client := http.Client{}
	client.Get(url)
	c.JSON(http.StatusOK, "ok")
}

func triggerSendOne(c *gin.Context) {
	name := c.Param("name")
	url := "http://" + common.ClientSvc + ":" + common.ClientPort + "/trigger/one?name="
	switch name {
	case common.RabbitMQ:
		url += ClientPathRabbitMQ
	case common.GRPC:
		url += ClientPathGRPC
	case common.Kafka:
		url += ClientPathKafka
	default:
		c.JSON(http.StatusBadRequest, "name unsupported")
		return
	}
	client := http.Client{}
	client.Get(url)
	c.JSON(http.StatusOK, "ok")
}
