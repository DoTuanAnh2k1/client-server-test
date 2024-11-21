package kafka

import "log"

func TriggerOn() {
	log.Println("Trigger on: send mess to kafka")
	isSend = true
}

func TriggerOff() {
	log.Println("Trigger off: send mess to kafka")
	isSend = false
}

func TriggerSendOne() {
	log.Println("Sending One mess to kafka")
	sendOne()
}
