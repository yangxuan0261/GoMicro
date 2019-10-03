package discovery

import (
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

	endpoints := []string{":2379"}

	if *role == kProvider {
		StartProvider("test", ":8080", endpoints)
	} else if *role == kConsumer {
		StartConsumer(endpoints)
	} else {
		log.Println("usage: discovery [provider|consumer]")
	}
}
