package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/danielblagy/l0-blagy/entity"
	stan "github.com/nats-io/stan.go"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("wb l0-blagy")

	log.Print("Connecting to the db...")

	dsn := "host=localhost user=l0user password=l0pass dbname=l0db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database\n", err)
	}

	log.Print("Connected to the database")

	db.AutoMigrate(&entity.Order{}, &entity.Item{})

	// test subscriber
	clusterID := "test-cluster"
	clientID := "test-subscriber"

	sc, _ := stan.Connect(clusterID, clientID)
	defer sc.Close()

	gotMessage := false
	var order entity.Order

	// Simple Async Subscriber
	sub, _ := sc.Subscribe("orders", func(m *stan.Msg) {
		json.Unmarshal([]byte(m.Data), &order)
		gotMessage = true
	})
	defer sub.Unsubscribe()

	for !gotMessage {

	}

	fmt.Println(order)

	orderJson, err := json.MarshalIndent(order, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(orderJson))
}
