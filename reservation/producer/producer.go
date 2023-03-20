package producer

import (
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
)

// ProducerConfig holds the configuration for creating a new Producer instance.
type ProducerConfig struct {
	Hosts         string
	OutgoingTopic string
}

// Producer represents a Kafka reservation instance.
type Producer struct {
	p     *kafka.Producer
	topic string
}

// NewProducer creates a new Kafka reservation instance.
func NewProducer(config ProducerConfig) *Producer {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": config.Hosts})
	if err != nil {
		log.Fatalf("Failed to create reservation: %s", err)
	}

	return &Producer{p: producer, topic: config.OutgoingTopic}
}

// Emit sends a message to the configured Kafka topic.
func (p *Producer) Emit(obj interface{}) error {
	messageBytes, err := json.Marshal(obj)
	if err != nil {
		log.Fatalf("Failed to marshal message: %s", err)
	}

	deliveryChan := make(chan kafka.Event)
	err = p.p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &p.topic, Partition: kafka.PartitionAny},
		Value:          messageBytes,
	}, deliveryChan)

	e := <-deliveryChan
	m := e.(*kafka.Message)
	if m.TopicPartition.Error != nil {
		log.Fatalf("Delivery failed: %v\n", m.TopicPartition.Error)
	} else {
		log.Printf("Delivered message to topic %s [%d] at offset %v\n", *m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	}

	if err != nil {
		log.Fatalf("Failed to produce message: %s", err)
		return err
	}

	return nil
}

// Close flushes and closes the Kafka reservation instance.
func (p *Producer) Close() {
	// Wait for messages to be delivered before shutting down the reservation.
	p.p.Flush(5 * 1000)
	// Close the reservation instance.
	p.p.Close()
}
