package main

import (
	"context"
	"fmt"
	"time"

	"kafka-go-examples/orderedmsg/config"

	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	Id string
	*kafka.Reader
}

// NewConsumer 创建一个消费者，它属于 config.GroupId 这个消费者组
func NewConsumer(id string) *Consumer {
	c := &Consumer{
		Id: id,
		Reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers: config.Brokers,
			GroupID: config.GroupId,
			Topic:   config.Topic,
			Dialer: &kafka.Dialer{
				ClientID: id,
			},
		}),
	}
	return c
}

// Read 读取消息，intervalMs 用来控制消费者的消费速度
func (c *Consumer) Read(intervalMs int) {
	fmt.Printf("%s start read\n", c.Id)
	for {
		msg, err := c.ReadMessage(context.Background())
		if err != nil {
			fmt.Printf("%s read msg err: %v\n", c.Id, err)
			return
		}
		// 模拟消费速度
		time.Sleep(time.Millisecond * time.Duration(intervalMs))
		fmt.Printf("%s read msg: %s, time: %s\n", c.Id, string(msg.Value), time.Now().Format("03:04:05"))
	}
}

func main() {
	//c1 := NewConsumer("consumer-1")
	//c1.Read(100)

	// 先启动 c1
	c1 := NewConsumer("consumer-1")
	go func() {
		c1.Read(500)
	}()

	// 5 秒后启动 c2
	time.Sleep(5 * time.Second)
	go func() {
		c2 := NewConsumer("consumer-2")
		c2.Read(300)
	}()

	// 再 10 秒后启动 c3 和 c4
	time.Sleep(10 * time.Second)
	go func() {
		c3 := NewConsumer("consumer-3")
		c3.Read(100)
	}()
	go func() {
		c4 := NewConsumer("consumer-4")
		c4.Read(100)
	}()

	select {}
}
