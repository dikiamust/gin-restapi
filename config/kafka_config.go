package config

import (
	"strings"
	"time"

	"github.com/IBM/sarama"
)

type KafkaConfig struct {
	Brokers []string
	Timeout time.Duration
}

func LoadKafkaConfig(cfg Config) *KafkaConfig {
	kafkaBrokers := cfg.KafkaBrokers

	// Parsing Kafka brokers, assuming broker addresses are separated by commas
	brokerList := strings.Split(kafkaBrokers, ",")

	return &KafkaConfig{
		Brokers: brokerList,
		Timeout: 10 * time.Second, // Example default timeout
	}
}

func NewKafkaProducer(config *KafkaConfig) (sarama.SyncProducer, error) {
	saramaConfig := sarama.NewConfig()
	saramaConfig.Producer.Return.Successes = true
	saramaConfig.Net.DialTimeout = config.Timeout

	producer, err := sarama.NewSyncProducer(config.Brokers, saramaConfig)
	if err != nil {
		return nil, err
	}

	return producer, nil
}

func NewKafkaConsumer(config *KafkaConfig) (sarama.Consumer, error) {
	saramaConfig := sarama.NewConfig()
	saramaConfig.Consumer.Offsets.Initial = sarama.OffsetNewest

	consumer, err := sarama.NewConsumer(config.Brokers, saramaConfig)
	if err != nil {
		return nil, err
	}

	return consumer, nil
}
