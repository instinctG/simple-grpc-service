package main

import (
	"fmt"
	"github.com/instinctG/simple-grpc-service/internal/db"
	"github.com/instinctG/simple-grpc-service/internal/rocket"
	"log"
)

func Run() error {
	fmt.Println("Rocker Service Starting...")

	rocketStore, err := db.New()
	if err != nil {
		return err
	}

	err = rocketStore.Migrate()
	if err != nil {
		log.Println("Failed to run migrations")
		return err
	}
	_ = rocket.New(rocketStore)
	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
