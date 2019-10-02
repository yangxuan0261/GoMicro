package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/registry/etcdv3"
	"log"
	"time"
)

func main() {
	registre := etcdv3.NewRegistry()

	service := micro.NewService(

		micro.Registry(registre),

		micro.Name("greeter"),

		micro.RegisterTTL(time.Second*30),

		micro.RegisterInterval(time.Second*15),
	)
	service.Init()
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
