package service

import (
	"encoding/json"
	"log"

	"github.com/danielblagy/l0-blagy/entity"
	"github.com/nats-io/stan.go"
)

// NATS Streaming client that subscribes to a channel and stores incoming data in a database
type StreamManager struct {
	ClusterID    string
	ClientID     string
	connection   stan.Conn
	subscription stan.Subscription
}

func NewStreamManager(clusterID, clientID string) *StreamManager {
	return &StreamManager{
		ClusterID: clusterID,
		ClientID:  clientID,
	}
}

func (sm *StreamManager) ConnectAndSubscribe(channelName string) error {
	sc, err := stan.Connect(sm.ClusterID, sm.ClientID)
	if err != nil {
		return err
	}

	sm.connection = sc

	// Simple Async Subscriber
	// TODO make a durable subscription instead of a reguler one (to store missed messages when was offline)
	sub, err := sc.Subscribe(channelName, sm.handleMessage)
	if err != nil {
		return err
	}

	sm.subscription = sub

	log.Print("Connected to the NATS server and subscribed to ", channelName, "channel")

	return nil
}

func (sm *StreamManager) Close() {
	log.Print("Closing a subscriber")
	// TODO don't unsubscribe a durable subscription, only close the connection
	sm.subscription.Unsubscribe()
	sm.connection.Close()
}

func (sm *StreamManager) handleMessage(m *stan.Msg) {
	var order entity.Order

	json.Unmarshal([]byte(m.Data), &order)
	log.Println("converted incoming json to entity", order)

	orderJson, err := json.MarshalIndent(order, "", "\t")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("converted entity to json (just for testing)", string(orderJson))
}
