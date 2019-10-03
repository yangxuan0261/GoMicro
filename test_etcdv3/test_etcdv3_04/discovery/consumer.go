package discovery

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"go.etcd.io/etcd/clientv3"
)

// StartConsumer start consumer
func StartConsumer(endpoints []string) {
	c := &Consumer{endpoints: endpoints, nodes: make(map[string]*Node)}
	c.Start()
}

// Consumer 服务的使用者,监听目录的变化,并做相应的处理
type Consumer struct {
	client    *clientv3.Client
	endpoints []string
	nodes     map[string]*Node
}

// Node service provider node
type Node struct {
	Key  string
	Info *ServiceInfo
}

func (c *Consumer) Start() {
	c.nodes = make(map[string]*Node)
	if c.dial() != nil {
		return
	}

	c.get()
	c.watch()
}

func (c *Consumer) dial() error {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   c.endpoints,
		DialTimeout: time.Second,
	})

	if err != nil {
		return err
	}

	c.client = cli
	return nil
}

func (c *Consumer) get() {
	rsp, err := c.client.Get(context.TODO(), ServiceKey, clientv3.WithPrefix())
	if err != nil {
		log.Printf("cannot get service:%+v\n", err)
		return
	}

	if rsp.Count == 0 {
		log.Printf("connot find service=%+v\n", ServiceKey)
		return
	}

	log.Printf("find service=%+v\n", ServiceKey)
	for _, kv := range rsp.Kvs {
		c.addNode(string(kv.Key), string(kv.Value))
	}
}

func (c *Consumer) watch() {
	log.Printf("start watch")
	wch := c.client.Watch(context.Background(), ServiceKey, clientv3.WithPrefix())
	for w := range wch {
		for _, ev := range w.Events {
			switch ev.Type {
			case clientv3.EventTypePut:
				c.addNode(string(ev.Kv.Key), string(ev.Kv.Value))
			case clientv3.EventTypeDelete:
				c.removeNode(ev)
			}
		}
	}
}

func (c *Consumer) addNode(key, val string) {
	info := &ServiceInfo{}
	err := json.Unmarshal([]byte(val), info)
	if err != nil {
		log.Fatalf("unmarshal info fail:%+v\n", err)
		return
	}
	node := c.nodes[key]
	if node == nil {
		// first add
		log.Printf("add node:%+v\n", info)
		node = &Node{Key: key, Info: info}
		c.nodes[key] = node
	} else {
		log.Printf("update node info:%+v\n", info)
		node.Info = info
	}

	// call api service
	c.call(node)
}

func (c *Consumer) call(node *Node) {
	url := fmt.Sprintf("%s%s", node.Info.Host, ServiceAPI)
	rsp, err := http.Get(url)
	if err != nil {
		log.Printf("call fail:%+v\n", err)
		return
	}

	defer rsp.Body.Close()

	body, _ := ioutil.ReadAll(rsp.Body)

	log.Printf("recv rsp: %s\n", body)
}

func (c *Consumer) removeNode(ev *clientv3.Event) {
	log.Printf("delete node:%+v\n", string(ev.Kv.Key))
	key := string(ev.Kv.Key)
	node := c.nodes[key]
	if node != nil {
		log.Printf("delete node info:%+v\n", node.Info)
		delete(c.nodes, key)
	}
}
