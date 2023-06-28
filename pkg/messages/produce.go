package messages

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/google/uuid"
)

func Must[T any](whatever T, err error) T {
	if err != nil {
		panic(err)
	}
	return whatever
}

type Message struct {
	Key     string `json:"key"`
	Message string `json:"message"`
}

func (msg Message) toJson() ([]byte, error) {
	jsonData, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

func getMsgKey(key string) string {
	if key == "" {
		return strings.Split(uuid.NewString(), "-")[0]
	} else {
		return key
	}
}

func Produce(config *kafka.ConfigMap, topic string, numMessages int, key string, delay string) error {
	producer, err := kafka.NewProducer(config)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to create Kafka Producer: %s\n", err))
	}
	defer producer.Close()

	// Delivery report handler for produced messages
	go func() {
		for e := range producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	delayDuration := Must(time.ParseDuration(delay))

	for n := 0; n < numMessages; n++ {
		msg := Message{
			Key:     getMsgKey(key),
			Message: fmt.Sprintf("msg-%d", n),
		}

		producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Key:            []byte(msg.Key),
			Value:          Must(msg.toJson()),
		}, nil)

		time.Sleep(delayDuration)
	}

	// Wait for message deliveries before shutting down
	producer.Flush(15 * 1000)

	return nil
}
