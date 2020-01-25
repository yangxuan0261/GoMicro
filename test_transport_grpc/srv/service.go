package main

import (
	"context"
	"fmt"

	proto "go-micro/proto_gen"

	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/transport/grpc"
)

// 官方默认使用的是 http

type User struct{}

func (u *User) Hello(ctx context.Context, req *proto.Request, res *proto.Response) error {
	fmt.Printf("--- req, name:%s\n", req.Name)
	res.Msg = "Hello " + req.Name
	return nil
}

func main() {
	var service micro.Service

	afterFunc := func() error {
		fmt.Println("--- afterFunc")
		return nil
	}

	service = micro.NewService(
		micro.Name("user"),
		micro.Transport(grpc.NewTransport()),
		micro.AfterStart(afterFunc),
	)

	service.Init()

	proto.RegisterUserHandler(service.Server(), new(User))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
