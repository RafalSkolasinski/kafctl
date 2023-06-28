package messages

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var keepGoing = true

func Consume(config *kafka.ConfigMap, topics []string, groupID string, delay string) error {

	config.SetKey("group.id", groupID)

	// Create Consumer
	consumer, err := kafka.NewConsumer(config)
	if err != nil {
		return err
	}

	defer consumer.Close()

	// Make sure that deferred c.Close will be closed
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt)

	go func() {
		select {
		case sig := <-channel:
			fmt.Printf("Got %s signal. Aborting...\n", sig)
			keepGoing = false
		}
	}()

	// Set topics and start reading messages
	consumer.SubscribeTopics(topics, nil)
	delayDuration := Must(time.ParseDuration(delay))

	for keepGoing {
		msg, err := consumer.ReadMessage(1 * time.Second)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else if err.(kafka.Error).Code() == kafka.ErrTimedOut {
			continue
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
		time.Sleep(delayDuration)
	}

	return nil
}
