package config

import (
	"bufio"
	"io"
	"log"
	"strings"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type ConfigMapReader struct {
	input  io.Reader
	output kafka.ConfigMap
}

func NewConfigMapReader(input io.Reader) *ConfigMapReader {
	return &ConfigMapReader{input: input, output: kafka.ConfigMap{}}
}

func (cmr *ConfigMapReader) ReadConfigMap() (*kafka.ConfigMap, error) {

	scanner := bufio.NewScanner(cmr.input)
	for scanner.Scan() {
		line := scanner.Text()
		err := cmr.addFromLine(line)
		if err != nil {
			return &kafka.ConfigMap{}, err
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return &cmr.output, nil
}

func (cmr *ConfigMapReader) addFromLine(line string) error {
	// Reject empty lines
	if line == "" {
		return nil
	}

	// Reject comment lines
	if strings.HasPrefix(line, "#") {
		return nil
	}

	// Set key value from string (there will be validation inside Set)
	return cmr.output.Set(line)
}
