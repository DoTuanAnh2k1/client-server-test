package server

import (
	"client/problem"
	"client/producer/grpc"
	"client/producer/kafka"
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
	case Kafka:
		kafka.TriggerOn()
	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
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
	case Kafka:
		kafka.TriggerOff()
	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
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
	w.Write([]byte("ok"))
}

func triggerSendOneHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	switch name {
	case RabbitMQ:
		rabbitmq.TriggerSendOne()
	case GRPC:
		grpc.TriggerSendOne()
	case Kafka:
		kafka.TriggerSendOne()
	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func newRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/trigger/on", triggerOnHandler)
	mux.HandleFunc("/trigger/off", triggerOffHandler)
	mux.HandleFunc("/trigger/one", triggerSendOneHandler)
	mux.HandleFunc("/prob", problemHandler)
	return mux
}
