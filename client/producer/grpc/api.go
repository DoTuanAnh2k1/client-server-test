package grpc

import "log"

func TriggerOn() {
	log.Println("Trigger on: request to server by using gRPC")
	isSend = true
}

func TriggerOff() {
	log.Println("Trigger off: request to server by using gRPC")
	isSend = false
}

func TriggerSendOne() {
	log.Println("Send one: request to server by using gRPC")
	sendOneReq()
}
