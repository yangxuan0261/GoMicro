package main

import (
	"encoding/json"
	"fmt"
	"time"

	dis "GoMicro/test_etcdv3/test_etcdv3_05/discovery"
)

func main() {
	m := dis.EtcdDis{Cluster: "tc"}
	m.Watch([]string{
		"http://127.0.0.1:2379",
	}, "s-test")
	m.Watch([]string{
		"http://127.0.0.1:2379",
	}, "s-xxxx")
	for {
		fmt.Println("--- check s-test")
		if e, ok := m.GetServiceInfoAllNode("s-test"); ok {
			tmp, _ := json.Marshal(e)
			fmt.Println(string(tmp))
			if len(e) > 0 {
				name, key, _ := dis.SplitServiceNameKey(e[0].Key)
				fmt.Printf("name:%s, key:%s\n", name, key)
			}
		}
		fmt.Println("--- check s-xxxx")
		if e, ok := m.GetServiceInfoAllNode("s-xxxx"); ok {
			tmp, _ := json.Marshal(e)
			fmt.Println(string(tmp))
			if len(e) > 0 {
				name, key, _ := dis.SplitServiceNameKey(e[0].Key)
				fmt.Printf("name:%s, key:%s\n", name, key)
			}
		}

		//fmt.Printf("nodes num = %d\n", len(m.Nodes))
		// fmt.Printf("nodes num = %d\n", 0)
		time.Sleep(time.Second * 2)
	}
}
