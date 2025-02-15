package pub_sub

import (
	"context"
	"github.com/google/uuid"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/kafka"
	"log"
	"testing"
)

func StartKafka(t *testing.T) bool {
	ctx := context.Background()

	kafkaContainer, err := kafka.Run(ctx,
		"confluentinc/confluent-local:7.5.0",
		kafka.WithClusterID(uuid.NewString()),
		testcontainers.WithHostPortAccess(9092),
	)

	defer func() {
		if err := testcontainers.TerminateContainer(kafkaContainer); err != nil {
			log.Printf("failed to terminate container: %s", err)
		}
	}()

	if err != nil {
		log.Printf("failed to start container: %s", err)
		t.Error(err)
		return false
	}

	return true
}
