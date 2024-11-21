package rabbitmq

import "log"

func TriggerOn() {
	isSend = true
}

func TriggerOff() {
	isSend = false
}

func TriggerSendOne() {
	log.Println("Sending One mess to rabbit MQ")
	sendOneReq()
}
