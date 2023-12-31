package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/rafalskolasinski/kafctl/pkg/config"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func readConfig(path string) *kafka.ConfigMap {
	// Open File and Defer its closing
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Cannot read config file", err)
	}
	defer file.Close()

	config, err := config.NewConfigMapReader(file).ReadConfigMap()
	if err != nil {
		log.Fatal(err)
	}
	return config
}

func logConfig(config *kafka.ConfigMap) {
	for k, v := range *config {
		if k == "sasl.password" {
			v = "<CENSORED>"
		}
		fmt.Printf("CONFIGURATION %s: ", k)
		fmt.Println(v)
	}
}
