package server

import (
	"client/problem"
	"client/producer/grpc"
	rabbitmq "client/producer/rabbitMQ"
	"client/producer/service"
	serviceheadless "client/producer/service_headless"
	"net/http"
)

func triggerOnHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	switch name {
	case ServiceHeadless:
		serviceheadless.TriggerOn()
	case Service:
		service.TriggerOn()
	case RabbitMQ:
		rabbitmq.TriggerOn()
	case GRPC:
		grpc.TriggerOn()
	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func triggerOffHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	switch name {
	case ServiceHeadless:
		serviceheadless.TriggerOff()
	case Service:
		service.TriggerOff()
	case RabbitMQ:
		rabbitmq.TriggerOff()
	case GRPC:
		grpc.TriggerOff()
	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func problemHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	switch name {
	case ProblemInit:
		problem.InitConnectionForProblem()
	case ProblemReconnect:
		problem.ReconnectClient()
	case ProblemDo:
		problem.ProblemDo()
	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func triggerSendOneHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	switch name {
	case RabbitMQ:
		rabbitmq.TriggerSendOne()
	case GRPC:
		grpc.TriggerSendOne()
	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func newRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/trigger/on", triggerOnHandler)
	mux.HandleFunc("/trigger/off", triggerOffHandler)
	mux.HandleFunc("/trigger/one", triggerSendOneHandler)
	mux.HandleFunc("/prob", problemHandler)
	return mux
}
