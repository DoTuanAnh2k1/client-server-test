package kafka

import (
	"log"
	"server/utils"
	"time"

	"github.com/segmentio/kafka-go"
)

func Init() {
	serverKafkaSvc = utils.GetEnv("ServerKafkaSvc", "localhost")
	serverKafkaPort = utils.GetEnv("ServerKafkaPort", "9092")
	brokerAddress := serverKafkaSvc + ":" + serverKafkaPort

	log.Println("Init reader to kafka with addr ", brokerAddress)
	reader = kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{brokerAddress},
		GroupID:  GroupId,
		Topic:    Topic,
		MinBytes: 10e3,
		MaxBytes: 10e6,
		MaxWait:  500 * time.Millisecond,
	})

	go receive()
}
