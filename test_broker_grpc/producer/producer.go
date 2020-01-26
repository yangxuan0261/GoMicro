package main

import (
	"fmt"
	"log"
	"time"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/config/cmd"
	"github.com/micro/go-plugins/broker/grpc"
)

var (
	topic = "topic.aaa.bbb.TpTest"
)

var grpcBroker broker.Broker

func pub() {
	tick := time.NewTicker(time.Second)
	i := 0
	for _ = range tick.C {
		msg := &broker.Message{
			Header: map[string]string{
				"id": fmt.Sprintf("%d", i),
			},
			Body: []byte(fmt.Sprintf("%d: %s", i, time.Now().String())),
		}
		if err := grpcBroker.Publish(topic, msg); err != nil {
			log.Printf("[pub] failed: %v", err)
		} else {
			fmt.Println("[pub] pubbed message:", string(msg.Body))
		}
		i++
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

	go pub()
	service := micro.NewService(
		micro.Name("com.a.aaa_producer"),
		micro.Broker(grpcBroker),
	)

	service.Init()

	// grpcBroker == service.Client().Options().Broker
	// log.Printf("--- addr1:%p\n", grpcBroker)
	// log.Printf("--- addr2:%p\n", service.Client().Options().Broker)

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
