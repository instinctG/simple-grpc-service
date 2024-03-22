package main

import (
	"fmt"
	"github.com/instinctG/simple-grpc-service/internal/db"
	"github.com/instinctG/simple-grpc-service/internal/rocket"
	"log"
)

func Run() error {
	fmt.Println("Rocker Service Starting...")

	database, err := db.New()
	if err != nil {
		return err
	}

	_ = rocket.New(database)
	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
