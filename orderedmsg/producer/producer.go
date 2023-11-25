package main

import (
	"context"
	"fmt"
	"kafka-go-examples/orderedmsg/config"

	"github.com/segmentio/kafka-go"
)

func NewProducer() *kafka.Writer {
	return &kafka.Writer{
		Addr:     config.Addr,
		Topic:    config.Topic,
		Balancer: &kafka.Hash{}, // 哈希分区
	}
}

func NewMessages(count int) []kafka.Message {
	res := make([]kafka.Message, count)
	for i := 0; i < count; i++ {
		res[i] = kafka.Message{
			Key:   config.MessageKey,
			Value: []byte(fmt.Sprintf("msg-%d", i+1)),
		}
	}
	return res
}

func main() {
	producer := NewProducer()
	messages := NewMessages(100)
	if err := producer.WriteMessages(context.Background(), messages...); err != nil {
		panic(err)
	}
	_ = producer.Close()
}
