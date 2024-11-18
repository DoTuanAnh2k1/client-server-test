package serviceheadless

import (
	"client/utils"
	"fmt"
	"time"
)

func ScanIp() {
	for {
		time.Sleep(5 * time.Second)
		go scanIp()
	}
}

func scanIp() {
	ipList := utils.LookupListIp(serverHeadlessSvc)
	// check remove
	for i := 0; i < len(connectionList); i++ {
		connection := connectionList[i]
		isExist := false
		for _, ip := range ipList {
			if connection.Ip == ip {
				isExist = true
				break
			}
		}
		if !isExist {
			fmt.Println("Remove Connection: ", connection.Ip)
			connectionList = append(connectionList[:i], connectionList[i+1:]...)
			fmt.Println("Connection List: ", connectionList)
			i--
		}
	}

	// check add new
	for _, ip := range ipList {
		isActive := false
		for _, connection := range connectionList {
			if connection.Ip == ip {
				isActive = true
				break
			}
		}
		if !isActive {
			fmt.Println("Add Connection: ", ip)
			connection := InitConnection(ip, serverPort)
			connectionList = append(connectionList, connection)
			fmt.Println("Connection List: ", connectionList)
		}
	}
}
