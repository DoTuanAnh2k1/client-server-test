package main

import (
	"mgt/common"
	"mgt/server"
	"mgt/utils"
)

func main() {
	port := utils.GetEnv("Port", "1234")
	common.InitVar()
	mgtServer := server.NewGin()
	mgtServer.Run(":" + port)
}
