package rabbitmq

import "log"

func TriggerOn() {
	log.Println("Trigger on: send mess to rabbit MQ")
	isSend = true
}

func TriggerOff() {
	log.Println("Trigger off: send mess to rabbit MQ")
	isSend = false
}

func TriggerSendOne() {
	log.Println("Sending One mess to rabbit MQ")
	sendOneReq()
}
