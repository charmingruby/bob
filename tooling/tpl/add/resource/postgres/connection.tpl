package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

type ConnectionInput struct {
	User         string
	Password     string
	Host         string
	DatabaseName string
	SSL          string
}

func New(in ConnectionInput) (*sqlx.DB, error) {
	connectionString := fmt.Sprintf(
		"postgresql://%s:%s@%s/%s?sslmode=%s",
		in.User,
		in.Password,
		in.Host,
		in.DatabaseName,
		in.SSL,
	)
	dbDriver := "postgres"

	db, err := sqlx.Connect(dbDriver, connectionString)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
