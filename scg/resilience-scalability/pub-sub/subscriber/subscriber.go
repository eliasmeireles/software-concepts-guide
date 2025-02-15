package main

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/eliasmeireles/software-concepts-guide/resilience-scalability/pub-sub"
	"log"
	"time"
)

func main() {
	pub_sub.SetConfig(&pub_sub.KafkaServerConfig{
		BootstrapServers: "localhost",
		Topic:            "SOFTWARE-CONCEPTS-GUIDE-PUB-SUB",
		GroupId:          "software-concepts-guide/subscriber",
	})

	StartSubscriber()
}

func StartSubscriber() {
	config := pub_sub.GetConfig()

	// Create a Kafka consumer
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": config.BootstrapServers, // Kafka bootstrap servers
		"group.id":          config.GroupId,          // Consumer group ID
		"auto.offset.reset": "earliest",              // Read messages from the earliest offset
	})

	if err != nil {
		panic(err)
	}

	err = consumer.SubscribeTopics([]string{config.Topic}, nil)

	if err != nil {
		panic(err)
	}

	defer func(consumer *kafka.Consumer) {
		err := consumer.Close()
		if err != nil {
			panic(err)
		}
	}(consumer)

	for {
		msg, err := consumer.ReadMessage(time.Second)
		if err == nil {
			if messageParser(msg) {
				return
			}
		} else if !err.(kafka.Error).IsTimeout() {
			log.Printf("Error: %v", err)
		}
	}
}

func messageParser(msg *kafka.Message) bool {
	json, err := pub_sub.MessageFromJson(string(msg.Value))
	if err != nil {
		log.Printf("Failed to convert message to JSON: %v", err)
		return true
	}

	HandleMessage(*json)
	return false
}

func HandleMessage(message pub_sub.Message) {
	// Implement your message handling logic here
	log.Printf("Processing message: %v", message)
}
