package rabbitmq

func TriggerOn() {
	isSend = true
}

func TriggerOff() {
	isSend = false
}

func TriggerSendOne() {
	sendOneReq()
}
