package main

import stan "github.com/nats-io/stan.go"

func main() {
	clusterID := "test-cluster"
	clientID := "test-publisher"

	sc, _ := stan.Connect(clusterID, clientID)
	defer sc.Close()

	// Simple Synchronous Publisher
	// does not return until an ack has been received from NATS Streaming
	sc.Publish("foo", []byte("Hello World"))
}
