package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/mvcc/mvccpb"
)

// 参考: https://www.jianshu.com/p/600f14e9e443
func main() {
	// 配置客户端连接
	client, err := clientv3.New(clientv3.Config{
		// Endpoints:   []string{"127.0.0.1:2379"},
		Endpoints:   []string{"127.0.0.1:2379", "127.0.0.1:22379", "127.0.0.1:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// 启动watch监听
	watch := client.Watch(context.TODO(), "aaa")
	go func() {
		for {
			watchResponse := <-watch
			for _, ev := range watchResponse.Events {
				switch ev.Type {
				case mvccpb.DELETE:
					fmt.Printf("监听到del：%s\n", ev.Kv.Key)
				case mvccpb.PUT:
					fmt.Printf("监听到put：%s, %s\n", ev.Kv.Key, ev.Kv.Value)
				}
			}
		}
	}()

	// 新增
	putResponse, err := client.Put(context.TODO(), "aaa", "xxx")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(putResponse.Header.String())

	// 查询
	getResponse, err := client.Get(context.TODO(), "aaa")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(getResponse.Kvs)

	// 删除
	deleteResponse, err := client.Delete(context.TODO(), "aaa")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(deleteResponse.Header.String())

	// 申请租约
	grantResponse, err := client.Grant(context.TODO(), 10)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 使用租约
	response, err := client.Put(context.TODO(), "aaa", "xxx", clientv3.WithLease(grantResponse.ID))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(response.Header.String())

	// 等待租约自动过期
	time.Sleep(time.Second * 20)
}
