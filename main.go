package main

import (
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

	// Simple Async Subscriber
	sub, _ := sc.Subscribe("foo", func(m *stan.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
		gotMessage = true
	})
	defer sub.Unsubscribe()

	for !gotMessage {

	}
}
