//go:generate mockgen -destination=rocket_mocks_test.go -package=rocket github.com/instinctG/simple-grpc-service/internal/rocket Store

package rocket

import (
	"context"
)

// Rocket - should contain things like the ID for the rocket,
// the name for the rocket and the type of rocket.
type Rocket struct {
	ID   string
	Name string
	Type string
}

// Store - defines the interface we need to satisfy for our
// service to work correctly
type Store interface {
	GetRocketByID(ctx context.Context, ID string) (Rocket, error)
	InsertRocket(ctx context.Context, rocket Rocket) (Rocket, error)
	DeleteRocket(ctx context.Context, ID string) error
}

// Service - our rocket service, used for updating our
// rocket inventory
type Service struct {
	Store Store
}

// New - returns a new rocket service
func New(store Store) Service {
	return Service{Store: store}
}

// GetRocketByID - retrieves a rocket from the store by ID
func (s Service) GetRocketByID(ctx context.Context, id string) (Rocket, error) {
	rkt, err := s.Store.GetRocketByID(ctx, id)
	if err != nil {
		return Rocket{}, err
	}
	return rkt, nil
}

// InsertRocket - Adds a rocket to our store
func (s Service) InsertRocket(ctx context.Context, rocket Rocket) (Rocket, error) {
	rkt, err := s.Store.InsertRocket(ctx, rocket)
	if err != nil {
		return Rocket{}, err
	}
	return rkt, nil
}

// DeleteRocket - deletes a rocket - most likely rapid
// unscheduled disassembly
func (s Service) DeleteRocket(ctx context.Context, id string) error {
	err := s.Store.DeleteRocket(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
