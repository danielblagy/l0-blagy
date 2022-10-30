package service

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/danielblagy/l0-blagy/entity"
	"github.com/nats-io/stan.go"
)

type Subscriber struct {
	connection   stan.Conn
	subscription stan.Subscription
}

func InitSubscriber() *Subscriber {
	return &Subscriber{}
}

func ConnectAndSubscribe(clusterID, clientID string) (*Subscriber, error) {
	sc, err := stan.Connect(clusterID, clientID)
	if err != nil {
		return nil, err
	}

	var order entity.Order

	// Simple Async Subscriber
	sub, err := sc.Subscribe("orders", func(m *stan.Msg) {
		json.Unmarshal([]byte(m.Data), &order)
		fmt.Println(order)

		orderJson, err := json.MarshalIndent(order, "", "\t")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(orderJson))
	})
	if err != nil {
		return nil, err
	}

	return &Subscriber{sc, sub}, nil
}

func (s *Subscriber) Close() {
	log.Print("Closing a subscriber")
	// TODO don't unsubscribe a durable subscription, only close the connection
	s.subscription.Unsubscribe()
	s.connection.Close()
}
