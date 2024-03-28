package db

import (
	"context"
	"errors"
	"fmt"
	"github.com/instinctG/simple-grpc-service/internal/rocket"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
	"log"
	"os"
)

// Store - is a struct for db connection
type Store struct {
	db *sqlx.DB
}

// New - returns a new Store
func New() (Store, error) {
	dbUsername := os.Getenv("DB_USERNAME")
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

	var rkt rocket.Rocket
	row := s.db.QueryRowContext(ctx, `SELECT id,name,type FROM rockets WHERE id=$1`, ID)
	err := row.Scan(&rkt.ID, &rkt.Name, &rkt.Type)
	if err != nil {
		log.Print(err.Error())
		return rocket.Rocket{}, err
	}

	return rkt, nil
}

// InsertRocket - inserts a new rocket into the database
func (s Store) InsertRocket(ctx context.Context, rocket2 rocket.Rocket) (rocket.Rocket, error) {

	//_, err := uuid.FromString(rocket2.ID)
	//if err != nil {
	//	return rocket.Rocket{}, errors.New("couldn't parse string format of id to uuid")
	//}

	_, err := s.db.QueryContext(ctx,
		`INSERT INTO rockets (id, name, type) VALUES ($1, $2, $3)`, rocket2.ID, rocket2.Name, rocket2.Type)
	if err != nil {
		return rocket.Rocket{}, errors.New("failed to insert into database")
	}
	return rocket.Rocket{
		ID:   rocket2.ID,
		Name: rocket2.Name,
		Type: rocket2.Type,
	}, nil
}

// DeleteRocket - deletes a rocket from database by id
func (s Store) DeleteRocket(ctx context.Context, ID string) error {
	uid, err := uuid.FromString(ID)
	if err != nil {
		return errors.New("couldn't parse string format of id to uuid")
	}
	_, err = s.db.ExecContext(ctx, `DELETE FROM rockets WHERE id=$1`, uid)
	if err != nil {
		return fmt.Errorf("couldn't delete rocket :%w", err)
	}

	return nil
}
