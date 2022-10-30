package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/danielblagy/l0-blagy/entity"
	"github.com/danielblagy/l0-blagy/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("wb l0-blagy")

	log.Print("Connecting to the database...")

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

	subscriber, err := service.ConnectAndSubscribe(clusterID, clientID)
	if err != nil {
		log.Fatal("Failed to connect and subscribe to NATS Streaming server", err)
	}
	defer subscriber.Close()

	scanner := bufio.NewScanner(os.Stdin)
	running := true
	for running {
		if scanner.Scan() {
			switch scanner.Text() {
			case "stop":
				running = false
			}
		}
	}
}
