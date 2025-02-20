package handlers

import (
	"go-restapi-gin/internal/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type KafkaHandler struct {
	kafkaService *services.KafkaService
}

func NewKafkaHandler(kafkaService *services.KafkaService) *KafkaHandler {
	return &KafkaHandler{kafkaService: kafkaService}
}

func (h *KafkaHandler) SendKafkaMessageHandler(c *gin.Context) {
	message := c.DefaultQuery("message", "default message")

	// Sending message to Kafka
	err := h.kafkaService.SendMessage("example-topic", message)
	if err != nil {
		log.Printf("Error sending message to Kafka: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to send message to Kafka",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Message successfully sent to Kafka",
	})
}

func (h *KafkaHandler) GetKafkaMessagesHandler(c *gin.Context) {
	topic := c.DefaultQuery("topic", "example-topic")

	// Receiving messages from Kafka
	messages, err := h.kafkaService.ConsumeMessages(topic)
	if err != nil {
		log.Printf("Error receiving messages from Kafka: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to receive messages from Kafka",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"messages": messages,
	})
}
