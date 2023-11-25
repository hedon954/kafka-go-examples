package config

import "github.com/segmentio/kafka-go"

var (
	Topic      = "ordered-msg-topic"
	Brokers    = []string{"kafka1.com:9092", "kafka2.com:9092", "kafka3.com:9092"}
	Addr       = kafka.TCP(Brokers...)
	GroupId    = "ordered-msg-group"
	MessageKey = []byte("message-key")
)
