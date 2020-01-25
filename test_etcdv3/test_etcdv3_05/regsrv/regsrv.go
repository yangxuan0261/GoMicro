package main

import (
	dis "go-micro/test_etcdv3/test_etcdv3_05/discovery"
	"time"
)

func main() {
	info := dis.EtcdServiceInfo{Info: "123453434324"}
	e := dis.EtcdDis{Cluster: "tc"}
	e.Register([]string{"http://127.0.0.1:2379"}, "s-test", "192.168.21.35", info)
	e.Register([]string{"http://127.0.0.1:2379"}, "s-xxxx", "127.0.0.1", info)
	time.Sleep(time.Second * 10)
	info.Info = "xxxxxxxxxxx"
	e.UpdateInfo("s-test", "192.168.21.35", info)
	time.Sleep(time.Second * 10)
}
