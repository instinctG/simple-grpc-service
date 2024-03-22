package main

import (
	"fmt"
	"log"
)

func Run() error {
	fmt.Println("Rocker Service Starting...")
	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
