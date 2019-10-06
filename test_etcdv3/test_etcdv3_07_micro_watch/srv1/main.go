package main

import (
	"context"

	demo "GoMicro/test_etcdv3/test_etcdv3_06/srv/proto/demo"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"

	"log"
	"math/rand"
	"time"
)

type Say struct{}

func (s *Say) Hello(ctx context.Context, req *demo.Request, rsp *demo.Response) error {
	log.Print("Received Say.Hello request")
	rsp.Msg = "Hello " + req.Name + " srv two. rand:" + string(rand.Intn(100))
	return nil
}

func DumpRegistryResult(rr *registry.Result) {
	log.Printf("------ watch, action:%s, service:%v\n", rr.Action, rr.Service)
	for k, v := range rr.Service.Metadata {
		log.Printf("Metadata, k:%s, v:%s\n", k, v)
	}

	for i, v := range rr.Service.Endpoints {
		log.Printf("Endpoints, i:%d, v:%v\n", i, v)
	}
}

func main() {
	registerDrive := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"http://127.0.0.1:2379",
		}
	})

	// 监听服务
	rw, err := registerDrive.Watch(func(wop *registry.WatchOptions) {
		wop = &registry.WatchOptions{
			Service: "aaaa", // 无效, 原因未知
			Context: context.Background(),
		}
	})
	if err == nil {
		go func() {
			for {
				rr, err2 := rw.Next()
				if err2 == nil {
					DumpRegistryResult(rr)
				}
			}
		}()
	}

	metaData := map[string]string{
		"aaa": "111",
		"bbb": "222",
	}
	_ = metaData

	service := micro.NewService(
		micro.Name("go.micro.srv.greeter"),
		micro.Version("1.0.1"),
		micro.Metadata(metaData), // 无效, 原因未知
		micro.Registry(registerDrive),
		micro.RegisterTTL(time.Second*10),
		micro.RegisterInterval(time.Second*5),
	)
	// optionally setup command line usage
	service.Init()

	// Register Handlers
	demo.RegisterSayHandler(service.Server(), new(Say))

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
