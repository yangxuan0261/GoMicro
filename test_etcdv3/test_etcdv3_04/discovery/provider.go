package discovery

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.etcd.io/etcd/clientv3"
)

// StartProvider start provider
func StartProvider(name string, host string, endpoints []string) {
	p := &Provider{}
	p.name = name
	p.endpoints = endpoints
	p.stop = make(chan error)
	p.info.Name = name
	p.info.Host = host
	p.Start()
}

// Provider 服务提供者
type Provider struct {
	client    *clientv3.Client
	leaseid   clientv3.LeaseID
	stop      chan error
	endpoints []string
	info      ServiceInfo
	name      string
}

// Start start
func (p *Provider) Start() {
	// start service
	go p.runService()

	// dial etcd
	if p.dial() != nil {
		return
	}

	// regist service and keep alive
	ch, err := p.putAndKeepAlive()
	if err != nil {
		return
	}

	// wait for exist
	for {
		select {
		case <-p.stop:
			p.revoke()
		case <-p.client.Ctx().Done():
			log.Printf("server closed\n")
		case _, ok := <-ch:
			if !ok {
				p.revoke()
			} else {
				// log.Printf("recv reply from service: %s, ttl:%d\n", p.name, ka.TTL)
			}
		}
	}
}

// Stop 结束服务
func (p *Provider) Stop() {
	p.stop <- nil
}

func (p *Provider) dial() error {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   p.endpoints,
		DialTimeout: time.Second,
	})

	if err != nil {
		fmt.Printf("dial etcd fail:%+v\n", err)
		return err
	}

	p.client = cli
	return nil
}

// putAndKeepAlive 注册服务,并使用Lease保持连接
func (p *Provider) putAndKeepAlive() (<-chan *clientv3.LeaseKeepAliveResponse, error) {
	key := ServiceKey + p.name
	val, _ := json.Marshal(p.info)
	// minimum lease TTL is 5-second
	rsp, err := p.client.Grant(context.TODO(), 5)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	_, err = p.client.Put(context.TODO(), key, string(val), clientv3.WithLease(rsp.ID))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	log.Printf("register service:%s,%s\n", key, val)
	return p.client.KeepAlive(context.TODO(), rsp.ID)
}

// revoke 撤销
func (p *Provider) revoke() error {
	_, err := p.client.Revoke(context.TODO(), p.leaseid)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("service stop:%+v\n", p.name)
	return err
}

// 提供简单的http服务
func (p *Provider) runService() {
	log.Printf("run service\n")
	http.HandleFunc(ServiceAPI, func(w http.ResponseWriter, r *http.Request) {
		log.Printf("recv req\n")
		w.Write([]byte("ok"))
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
