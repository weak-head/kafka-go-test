package main

import (
	"context"
	"log"
	"os"
	"strings"

	kafka "github.com/segmentio/kafka-go"
)

func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		GroupID:  groupID,
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
}

func consume(r *kafka.Reader) {
	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		log.Printf(
			"message at topic:%v partition:%v offset:%v	%s = %s\n",
			m.Topic,
			m.Partition,
			m.Offset,
			string(m.Key),
			string(m.Value),
		)
	}
}

func main() {
	kafkaURL := os.Getenv("kafkaURL")
	groupID := os.Getenv("groupID")
	topic := os.Getenv("topic")

	reader := getKafkaReader(kafkaURL, topic, groupID)
	defer reader.Close()

	consume(reader)
}
