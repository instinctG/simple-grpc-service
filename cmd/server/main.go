package main

import (
	"fmt"
	"github.com/instinctG/simple-grpc-service/internal/db"
	"github.com/instinctG/simple-grpc-service/internal/rocket"
	"github.com/instinctG/simple-grpc-service/internal/transport/grpc"
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

	rktService := rocket.New(rocketStore)
	rktHandler := grpc.New(rktService)

	if err := rktHandler.Serve(); err != nil {
		return err
	}
	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
