package main

import (
	"context"

	demo "go-micro/test_etcdv3/test_etcdv3_06/srv/proto/demo"

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
	log.Printf("------ watch, action: [%s], service:%v\n", rr.Action, rr.Service)

	log.Printf("--- Endpoints, len:%d\n", len(rr.Service.Endpoints))
	for i, v := range rr.Service.Endpoints {
		log.Printf("Endpoints, i:%d, v:%v\n", i, v)
	}

	log.Printf("--- Nodes, len:%d\n", len(rr.Service.Nodes))
	for i, v := range rr.Service.Nodes {
		log.Printf("Nodes, i:%d, Id:%s, Address:%s\n", i, v.Id, v.Address)
		for k, v := range v.Metadata {
			log.Printf("Metadata, k:%s, v:%s\n", k, v)
		}
	}
}

func main() {
	registerDrive := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"http://127.0.0.1:2379",
		}
	})

	rw, err := registerDrive.Watch(func(wop *registry.WatchOptions) {
		wop.Service = ""
		wop.Context = context.Background()
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
		"ccc": "333",
		"ddd": "444",
	}

	service := micro.NewService(
		micro.Name("go.micro.srv.greeter"),
		micro.Version("1.0.2"),
		micro.Metadata(metaData),
		micro.RegisterTTL(time.Second*3),
		micro.RegisterInterval(time.Second*2),
		micro.Registry(registerDrive),
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
