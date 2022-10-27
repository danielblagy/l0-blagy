package main

import (
	"fmt"
	"log"

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

	log.Print("Connected to the database\n", db)
}
