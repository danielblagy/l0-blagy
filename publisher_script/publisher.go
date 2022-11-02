package main

import (
	"log"
	"os"

	stan "github.com/nats-io/stan.go"
)

func main() {
	clusterID := "test-cluster"
	clientID := "test-publisher"

	newOrder, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal("Failed to read file", err)
	}

	sc, _ := stan.Connect(clusterID, clientID)
	defer sc.Close()

	// NOTE: Sycnchronous publisher
	sc.Publish("orders", []byte(newOrder))
}
