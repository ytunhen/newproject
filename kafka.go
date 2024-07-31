package main

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

var kafkaWriter *kafka.Writer

func initKafka() {
	kafkaWriter = &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    "message-topic",
		Balancer: &kafka.LeastBytes{},
	}
}

func sendMessageToKafka(msg Message) error {
	return kafkaWriter.WriteMessages(
		context.Background(),
		kafka.Message{
			Key:   []byte(fmt.Sprintf("Message-%d", msg.ID)),
			Value: []byte(msg.Content),
		},
	)
}
