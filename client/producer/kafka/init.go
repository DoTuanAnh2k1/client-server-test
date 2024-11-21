package kafka

import (
	"client/utils"
	"log"

	"github.com/segmentio/kafka-go"
)

func Init() {
	serverKafkaSvc = utils.GetEnv("ServerKafkaSvc", "localhost")
	serverKafkaPort = utils.GetEnv("ServerKafkaPort", "9092")
	brokerAddress := serverKafkaSvc + ":" + serverKafkaPort
	log.Println("Init writer to kafka with address: ", brokerAddress)
	writer = kafka.Writer{
		Addr:         kafka.TCP(brokerAddress),
		Topic:        Topic,
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: kafka.RequireOne,
	}

	go sendreq()
}
