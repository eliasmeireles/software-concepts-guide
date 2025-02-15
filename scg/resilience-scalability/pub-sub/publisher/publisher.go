package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/eliasmeireles/software-concepts-guide/resilience-scalability/pub-sub"
	"log"
	"math/rand"
	"time"
)

func main() {
	pub_sub.SetConfig(&pub_sub.KafkaServerConfig{
		BootstrapServers: "localhost",
		Topic:            "SOFTWARE-CONCEPTS-GUIDE-PUB-SUB",
		GroupId:          "software-concepts-guide/publisher-sub",
	})

	go func() {
		for {
			// Generate a random number of messages
			numMessages := rand.Intn(10) + 1 // Random number between 1 and 10

			for i := 0; i < numMessages; i++ {
				message := pub_sub.Message{
					Id:   rand.Intn(1000),                               // Random ID
					Data: fmt.Sprintf("This is message %d", rand.Int()), // Random message content
				}

				if !PublishMessage(message) {
					log.Printf("Failed to publish message: %v", message)
				}
			}

			// Sleep for a random duration
			time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond) // Random sleep between 0 and 3 seconds
		}
	}()

	select {}
}

func getProducer() *kafka.Producer {
	config := pub_sub.GetConfig()

	// Configurações do produtor
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": config.BootstrapServers, // Endereço do servidor Kafka
		"client.id":         config.GroupId,          // Nome do produtor
	})

	if err != nil {
		log.Fatalf("Erro ao criar o produtor: %v", err)
	}
	return producer
}

func PublishMessage(message pub_sub.Message) bool {
	messageJson, err := message.ToJson()
	if err != nil {
		log.Printf("Failed to convert message to JSON: %v", err)
		return false
	}

	producer := getProducer()
	config := pub_sub.GetConfig()
	defer producer.Close()

	err = producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &config.Topic, Partition: kafka.PartitionAny},
		Value:          []byte(messageJson),
	}, nil)

	if err != nil {
		log.Printf("Failed to produce message: %v", err)
		return false
	}
	log.Printf("Message published: %s", messageJson)
	return true
}
