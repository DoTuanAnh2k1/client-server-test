package kafka

import "github.com/segmentio/kafka-go"

const Topic = "queue"
const GroupId = "group"

var (
	serverKafkaSvc  string
	serverKafkaPort string
)

var (
	reader *kafka.Reader
)

var isSend bool
