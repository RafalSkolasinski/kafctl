package topics

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Get(config *kafka.ConfigMap) error {
	admin, err := kafka.NewAdminClient(config)
	if err != nil {
		fmt.Printf("Failed to create Admin client: %s\n", err)
	}
	defer admin.Close()

	metadata, err := admin.GetMetadata(nil, true, timeout)
	if err != nil {
		return err
	}

	for topic, topicMetadata := range metadata.Topics {
		fmt.Printf("Topic (partitions: %d): %s\n", len(topicMetadata.Partitions), topic)
	}
	return nil
}
