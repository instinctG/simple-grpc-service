// +acceptance

package test

import (
	v1 "github.com/instinctG/protos/rocket/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func GetClient() v1.RocketServiceClient {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("couldn't connect: %s", err)
	}

	rocketClient := v1.NewRocketServiceClient(conn)
	return rocketClient
}
