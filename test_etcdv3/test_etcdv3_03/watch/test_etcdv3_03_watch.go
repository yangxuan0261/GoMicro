package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/golang/glog"
	"go.etcd.io/etcd/clientv3"
)

func WATCH(cli *clientv3.Client, key string, opts ...clientv3.OpOption) (clientv3.WatchChan, error) {
	watchChan := cli.Watch(context.Background(), key, opts...)
	return watchChan, nil
}

func toString(obj interface{}) (string, error) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
func main() {

	config := clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	}

	cli, err := clientv3.New(config)
	if err != nil {
		glog.Fatal(err.Error())
	}

	defer func() {
		if err := cli.Close(); err != nil {
			glog.Error(err.Error())
		}
	}()

	if watchChan, err := WATCH(cli, "sample_key"); err != nil { // 监听这个 sample_key 这个 key
		glog.Errorf(err.Error())
	} else {
		for {
			wr := <-watchChan
			if str, err := toString(wr); err == nil {
				fmt.Printf("--- watch: %s\n", str)
			} else {
				glog.Error(err.Error())
			}
		}
	}
}
