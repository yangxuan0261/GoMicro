package main

import (
	"go-micro/test_etcdv3/test_etcdv3_04/discovery"
	"flag"
	"log"
)

const (
	kProvider = 1
	kConsumer = 2
)

func main() {
	role := flag.Int("role", 0, "provider|consumer")
	flag.Parse()
	log.Printf("role=%+v\n", *role)

	endpoints := []string{":2379"}

	if *role == kProvider {
		discovery.StartProvider("test", "http://localhost:8080", endpoints)
	} else if *role == kConsumer {
		discovery.StartConsumer(endpoints)
	} else {
		log.Println("usage: discovery --role=[1|2]")
	}
}
