package grpc

import (
	"context"
	"errors"
	rkt "github.com/instinctG/protos/rocket/v1"
	"github.com/instinctG/simple-grpc-service/internal/rocket"
	"google.golang.org/grpc"
	"log"
	"net"
)

type RocketService interface {
	GetRocketByID(ctx context.Context, id string) (rocket.Rocket, error)
	InsertRocket(ctx context.Context, rkt rocket.Rocket) (rocket.Rocket, error)
	DeleteRocket(ctx context.Context, id string) error
}

type Handler struct {
	rkt.UnimplementedRocketServiceServer
	RocketService RocketService
}

func New(rkt RocketService) Handler {
	return Handler{RocketService: rkt}
}

func (h Handler) Serve() error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Println("couldn't listen on port 50051")
		return err
	}

	grpcServer := grpc.NewServer()
	rkt.RegisterRocketServiceServer(grpcServer, &h)

	if err := grpcServer.Serve(lis); err != nil {
		log.Printf("failed to server: %s\n", err)
		return err
	}

	return nil
}

// GetRocket - gets a rocket from the database
func (h Handler) GetRocket(ctx context.Context, request *rkt.GetRocketRequest) (*rkt.GetRocketResponse, error) {
	log.Print("Get Rocket gRPC Endpoint Hit")

	rocket, err := h.RocketService.GetRocketByID(ctx, request.Id)
	if err != nil {
		log.Print("Failed to retrieve rocket by ID")
		return &rkt.GetRocketResponse{}, err
	}

	return &rkt.GetRocketResponse{Rocket: &rkt.Rocket{
		Id:   rocket.ID,
		Name: rocket.Name,
		Type: rocket.Type,
	}}, nil
}

// AddRocket - adds a rocket to the database
func (h Handler) AddRocket(ctx context.Context, request *rkt.AddRocketRequest) (*rkt.AddRocketResponse, error) {
	log.Print("Add Rocket gRPC Endpoint Hit")

	newRkt, err := h.RocketService.InsertRocket(ctx, rocket.Rocket{
		ID:   request.Rocket.Id,
		Name: request.Rocket.Name,
		Type: request.Rocket.Type,
	})
	if err != nil {
		log.Print("failed to insert rocket into database")
		return &rkt.AddRocketResponse{}, err
	}

	return &rkt.AddRocketResponse{Rocket: &rkt.Rocket{
		Id:   newRkt.ID,
		Name: newRkt.Name,
		Type: newRkt.Type,
	}}, nil
}

// DeleteRocket - handler for deleting a rocket
func (h Handler) DeleteRocket(ctx context.Context, request *rkt.DeleteRocketRequest) (*rkt.DeleteRocketResponse, error) {
	log.Print("Delete Rocket gRPC Endpoint Hit")

	err := h.RocketService.DeleteRocket(ctx, request.Rocket.Id)
	if err != nil {
		return &rkt.DeleteRocketResponse{}, errors.New("failed to delete rocket from database")
	}

	return &rkt.DeleteRocketResponse{Status: "successfully deleted rocket"}, nil
}
