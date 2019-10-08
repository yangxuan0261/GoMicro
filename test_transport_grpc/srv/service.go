package main

import (
	"context"
	"fmt"

	proto "GoMicro/proto_gen"

	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/transport/grpc"
)

type User struct{}

func (u *User) Hello(ctx context.Context, req *proto.Request, res *proto.Response) error {
	fmt.Printf("--- req, name:%s\n", req.Name)
	res.Msg = "Hello " + req.Name
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("user"),
		micro.Transport(grpc.NewTransport()),
	)

	service.Init()

	proto.RegisterUserHandler(service.Server(), new(User))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
