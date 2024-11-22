package kafka

import (
	"client/common"
	"github.com/segmentio/kafka-go"
)

const Topic = "queue"

var (
	serverKafkaSvc  string
	serverKafkaPort string
)

var (
	writer  kafka.Writer
	message = kafka.Message{
		Key:   []byte("key-A"),
		Value: []byte(common.MessageBody),
	}
)

var isSend bool
