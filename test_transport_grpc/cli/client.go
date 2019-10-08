package main

import (
	"context"
	"fmt"

	proto "GoMicro/proto_gen"

	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/transport/grpc"
)

func main() {
	service := micro.NewService(
		micro.Name("user.client"),
		micro.Transport(grpc.NewTransport()),
	)
	service.Init()

	user := proto.NewUserService("user", service.Client())

	res, err := user.Hello(context.TODO(), &proto.Request{Name: "World ^_^"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.Msg)
}
