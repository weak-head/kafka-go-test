package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	kafka "github.com/segmentio/kafka-go"
)

func newKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

func produce(w *kafka.Writer) {
	for i := 0; ; i++ {
		msg := kafka.Message{
			Key:   []byte(fmt.Sprintf("key-%d", i)),
			Value: []byte(fmt.Sprintf("%s", uuid.New())),
		}

		err := w.WriteMessages(context.Background(), msg)
		if err != nil {
			fmt.Println(err)
		}

		time.Sleep(1 * time.Second)
	}
}

func main() {
	kafkaURL := os.Getenv("kafkaURL")
	topic := os.Getenv("topic")

	writer := newKafkaWriter(kafkaURL, topic)
	defer writer.Close()

	produce(writer)
}
