package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

func main() {
	var name = "127.0.0.1:9876"
	p, err := rocketmq.NewProducer(
		producer.WithNameServer(primitive.NamesrvAddr{name}),
		//producer.WithNsResolver(primitive.NewPassthroughResolver(endPoint)),
		producer.WithRetry(2),
		producer.WithGroupName("gid-xiaozou-grpc"),
	)
	if err != nil {
		fmt.Println("new product create error ", err)

	}
	err1 := p.Start()
	if err != nil {
		fmt.Println("start prodcut error", err1)
	}

	result, err := p.SendSync(context.Background(), &primitive.Message{
		Topic: "xiao-zou-topic",
		Body:  []byte("xiao-zou-topic test grpc"),
	})
	fmt.Println("result ", result)
}
