package pub_sub

import "os"

type KafkaServerConfig struct {
	BootstrapServers string
	Topic            string
	GroupId          string
}

var kafkaConfig *KafkaServerConfig

func SetConfig(config *KafkaServerConfig) {
	kafkaConfig = config
}

func GetConfig() *KafkaServerConfig {
	if kafkaConfig == nil {
		bootstrapServers := os.Getenv("KAFKA_BOOTSTRAP_SERVERS")

		if bootstrapServers == "" {
			panic("KAFKA_BOOTSTRAP_SERVERS is not set")
		}

		kafkaConfig = &KafkaServerConfig{
			BootstrapServers: bootstrapServers,
			Topic:            "SOFTWARE-CONCEPTS-GUIDE-PUB-SUB",
			GroupId:          "software-concepts-guide/publisher-sub",
		}
	}
	return kafkaConfig
}
