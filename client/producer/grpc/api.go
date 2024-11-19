package grpc

func TriggerOn() {
	isSend = true
}

func TriggerOff() {
	isSend = false
}

func TriggerSendOne() {
	sendOneReq()
}
