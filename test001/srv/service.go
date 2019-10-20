package main

import (
	"context"
	"fmt"

	proto "GoMicro/proto_gen"

	micro "github.com/micro/go-micro"
)

type User struct{}

func (u *User) Hello(ctx context.Context, req *proto.Request, res *proto.Response) error {
	res.Msg = "Hello " + req.Name
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("user"),
	)

	service.Init()

	fmt.Println("--- service id:", service.String())

	confSrv := service.Server().Options()
	fmt.Printf("--- confSrv:%+v\n", confSrv)

	confCli := service.Client().Options()
	fmt.Printf("--- confCli:%+v\n", confCli)

	proto.RegisterUserHandler(service.Server(), new(User))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
