package main

import (
	"kafka-test/client"
	"kafka-test/server"
)


/*
docker run -d \
  --name zookeeper \
  -p 2181:2181 \
  -e ZOOKEEPER_CLIENT_PORT=2181 \
  -e ZOOKEEPER_TICK_TIME=2000 \
  confluentinc/cp-zookeeper:7.5.0

docker run -d \
  --name kafka \
  -p 9092:9092 \
  --link zookeeper:zookeeper \
  -e KAFKA_BROKER_ID=1 \
  -e KAFKA_ZOOKEEPER_CONNECT=zookeeper:2181 \
  -e KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://localhost:9092 \
  -e KAFKA_LISTENERS=PLAINTEXT://0.0.0.0:9092 \
  -e KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR=1 \
  -e KAFKA_AUTO_CREATE_TOPICS_ENABLE="true" \
  confluentinc/cp-kafka:7.5.0

kafka-topics --create --topic <topic_name> --bootstrap-server <server_address> --partition 1 --replication-factor 1

kafka-topics --create --topic example-topic --bootstrap-server localhost:9092 --partitions 3 --replication-factor 1

kafka-topics --bootstrap-server localhost:9092 --topic example-topic --delete
*/
func main() {
	go client.NewClient(1)
	// go client.NewClient(2)
	go server.NewServer(1)
	server.NewServer(2)
}
