package main

import (
	"fmt"
	"log"

	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/config/cmd"

	"github.com/micro/go-plugins/broker/grpc"

	micro "github.com/micro/go-micro"
)

// 官方demo示例: src/github.com/micro/examples/broker/producer/producer.go
// 默认使用的是 http

var (
	topic = "topic.aaa.bbb.TpTest"
)

var grpcBroker broker.Broker

// Example of a subscription which receives all the messages
func sub() {
	_, err := grpcBroker.Subscribe(topic, func(p broker.Event) error {
		fmt.Println("[sub] received message:", string(p.Message().Body), "header", p.Message().Header)
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	cmd.Init()

	grpcBroker = grpc.NewBroker()

	if err := grpcBroker.Init(); err != nil {
		log.Fatalf("Broker Init error: %v", err)
	}
	if err := grpcBroker.Connect(); err != nil {
		log.Fatalf("Broker Connect error: %v", err)
	}

	go sub()

	service := micro.NewService(
		micro.Name("com.aaa.bbb.SrvGame"),
		micro.Broker(grpcBroker),
	)

	service.Init()

	// grpcBroker == service.Server().Options().Broker
	// log.Printf("--- addr1:%p\n", grpcBroker)
	// log.Printf("--- addr2:%p\n", service.Server().Options().Broker)

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
