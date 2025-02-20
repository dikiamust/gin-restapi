package services

import (
	"fmt"
	"go-restapi-gin/config"
	"log"
	"time"

	"github.com/IBM/sarama"
)

type KafkaService struct {
	producer sarama.SyncProducer
	consumer sarama.Consumer
}

func NewKafkaService(kafkaConfig *config.KafkaConfig) (*KafkaService, error) {
	// Initializing Kafka producer
	producer, err := config.NewKafkaProducer(kafkaConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create Kafka producer: %v", err)
	}

	// Inisialisasi Kafka consumer
	consumer, err := config.NewKafkaConsumer(kafkaConfig)
	if err != nil {
		producer.Close() // If the consumer creation fails, close the producer
		return nil, fmt.Errorf("failed to create Kafka consumer: %v", err)
	}

	// Return the initialized KafkaService
	return &KafkaService{
		producer: producer,
		consumer: consumer,
	}, nil
}

// SendMessage sends a message to Kafka
func (k *KafkaService) SendMessage(topic string, message string) error {

	// Create Kafka message
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	// Send message to Kafka
	_, _, err := k.producer.SendMessage(msg)
	if err != nil {
		return fmt.Errorf("failed to send message to Kafka: %v", err)
	}

	log.Printf("Message sent to Kafka topic %s: %s", topic, message)
	return nil
}

// ConsumeMessages starts receiving messages from Kafka
func (s *KafkaService) ConsumeMessages(topic string) ([]string, error) {

	// Open partition to read messages starting from the newest offset
	// partitionConsumer, err := s.consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)

	// Open partition to read messages starting from the oldest offset
	partitionConsumer, err := s.consumer.ConsumePartition(topic, 0, sarama.OffsetOldest)

	if err != nil {
		return nil, fmt.Errorf("failed to start consuming partition: %w", err)
	}
	defer partitionConsumer.Close()

	// Store received messages
	var messages []string

	timeout := time.After(10 * time.Second) // 10-second timeout
	messageCount := 0
	maxMessages := 5 // Limit to 5 messages

	for {
		select {
		case msg := <-partitionConsumer.Messages():
			messages = append(messages, string(msg.Value))
			messageCount++
			if messageCount >= maxMessages {
				return messages, nil // Limit to 5 messages
			}
		case <-timeout:
			return messages, nil // Timeout and return messages
		}
	}

}

// Close the Kafka producer and consumer
func (k *KafkaService) Close() {
	k.producer.Close()
	k.consumer.Close()
}
