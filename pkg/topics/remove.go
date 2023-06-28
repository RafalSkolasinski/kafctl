package topics

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Remove(config *kafka.ConfigMap, topics []string) error {
	// Create a new Kafka admin client
	admin, err := kafka.NewAdminClient(config)
	if err != nil {
		fmt.Printf("Failed to create Admin client: %s\n", err)
	}
	defer admin.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	results, err := admin.DeleteTopics(
		ctx,
		topics,
		kafka.SetAdminOperationTimeout(timeout*time.Millisecond))

	if err != nil {
		return errors.New(fmt.Sprintf("Failed to remove topic: %v\n", err))
	}

	// Print results
	for _, result := range results {
		fmt.Printf("Removed: %s\n", result)
	}
	return nil
}
