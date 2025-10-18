package db

import (
	"database/sql"
	"log"

	"github.com/VergilX/my-space/internal/dblayer"
	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	Conn    *sql.DB
	Queries *dblayer.Queries
}

func New(dsn string) (*DB, error) {
	database, err := sql.Open("sqlite3", dsn)
	if err != nil {
		log.Fatalf("Could not start database. Err: %v", err)
		return nil, err
	}

	// use struct
	db := DB{
		Conn:    database,
		Queries: dblayer.New(database),
	}

	return &db, nil
}

func (db *DB) CloseConn() error {
	err := db.Conn.Close()
	if err != nil {
		log.Fatalf("DB connection termination error")
		return err
	}

	return nil
}
