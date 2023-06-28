package topics

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Create(config *kafka.ConfigMap, topics []string, numPartitions, replicationFactor int) error {
	// Create a new Kafka admin client
	admin, err := kafka.NewAdminClient(config)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to create Admin client: %s\n", err))
	}
	defer admin.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	topicSpecs := []kafka.TopicSpecification{}
	for _, topic := range topics {
		topicSpecs = append(
			topicSpecs,
			kafka.TopicSpecification{
				Topic:             topic,
				NumPartitions:     numPartitions,
				ReplicationFactor: replicationFactor,
			},
		)
	}

	results, err := admin.CreateTopics(
		ctx,
		topicSpecs,
		kafka.SetAdminOperationTimeout(timeout*time.Millisecond))

	if err != nil {
		return errors.New(fmt.Sprintf("Failed to create topic: %v\n", err))
	}

	// Print results
	for _, result := range results {
		fmt.Printf("Created: %s\n", result)
	}
	return nil
}
