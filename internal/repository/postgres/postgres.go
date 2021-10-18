package postgres

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const linksTable = "links"

func NewPostgresDB(config string) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", config)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
