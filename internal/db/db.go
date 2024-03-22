package db

import (
	"context"
	"fmt"
	"github.com/instinctG/simple-grpc-service/internal/rocket"
	"github.com/jmoiron/sqlx"
	"os"
)

// DB
type Store struct {
	db *sqlx.DB
}

// New - returns a new Store
func New() (Store, error) {
	dbUsername := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbTable := os.Getenv("DB_TABLE")
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")
	SSLMode := os.Getenv("DB_SSL_MODE")

	connString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		dbHost,
		dbPort,
		dbUsername,
		dbTable,
		dbPassword,
		SSLMode,
	)

	db, err := sqlx.Connect("postgres", connString)
	if err != nil {
		return Store{}, err
	}

	return Store{db: db}, nil
}

// GetRocketByID - returns a rocket from the database by id
func (s Store) GetRocketByID(ctx context.Context, ID string) (rocket.Rocket, error) {
	return rocket.Rocket{}, nil
}

// InsertRocket - inserts a new rocket into the database
func (s Store) InsertRocket(ctx context.Context, rocket2 rocket.Rocket) (rocket.Rocket, error) {
	return rocket.Rocket{}, nil
}

// DeleteRocket - deletes a rocket from database by id
func (s Store) DeleteRocket(ctx context.Context, ID string) error {
	return nil
}
