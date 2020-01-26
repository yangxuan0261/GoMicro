package main

import (
	"fmt"
	"strings"

	"context"

	"github.com/micro/go-micro"

	pb "github.com/micro/go-micro/agent/proto"
)

type Command struct{}

// Help returns the command usage
func (c *Command) Help(ctx context.Context, req *pb.HelpRequest, rsp *pb.HelpResponse) error {
	rsp.Usage = "wilker"
	rsp.Description = "This is an example bot command as a micro service wilker"
	return nil
}

// Exec executes the command
func (c *Command) Exec(ctx context.Context, req *pb.ExecRequest, rsp *pb.ExecResponse) error {
	for i, val := range req.Args {
		fmt.Printf("--- i:%d, val:%s\n", i, val)
	}
	rsp.Result = []byte(strings.Join(req.Args, "---"))
	// rsp.Error could be set to return an error instead
	// the function error would only be used for service level issues
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.bot.wilker"),
	)

	service.Init()

	pb.RegisterCommandHandler(service.Server(), new(Command))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
